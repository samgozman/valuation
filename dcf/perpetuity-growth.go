package dcf

import (
	"errors"
	"time"

	"github.com/samgozman/valuation/internal/utils"
)

// First index is always the current state of balance sheet.
// If you want to calculate DCF for 5 years when you need to
// pass 6 elements of PerpetuityGrowthPeriod into function.
// First element will be the current state state of balance sheet
// and the other 5 will represent your assumptions.
type PerpetuityGrowthPeriod struct {
	Revenue       int64
	EBITDA        int64     // Earnings before interest, taxes, depreciation and amortization
	OtherIncome   int64     // Other Income / Expenses
	CapEx         int64     // Capital Expenditures
	NWC           int64     // Net Working Capital
	DA            int64     // Depreciation & Amortization
	TaxRate       float32   // Tax rate assumption
	DiscountRate  float32   // Discount rate (interest rate). Usually equals to WACC
	PGR           float32   // Perpetuity Growth Rate - Typically between the inflation rate of 2-3% and the GDP growth rate of 4-5%
	BeginningDate time.Time // Balance Sheet beginning date
	EndingDate    time.Time // Balance Sheet ending date
}

// Calculate enterprise value by Perpetuity Growth Model
func PerpetuityGrowth(periods *[]PerpetuityGrowthPeriod, currentDate time.Time) (float32, error) {
	periodsNumber := len(*periods)

	if periodsNumber < 2 {
		return 0, errors.New("number of periods must be 2+")
	}

	// 1. Calculate discrete Cash Flows
	var dFCF_sum float32
	var terminal_fcf float32
	for i := 1; i < periodsNumber; i++ {
		nopat := utils.NOPAT(utils.NOPATParams{
			EBITDA:      (*periods)[i].EBITDA,
			OtherIncome: (*periods)[i].OtherIncome,
			DA:          (*periods)[i].DA,
			TaxRate:     (*periods)[i].TaxRate,
		})

		nwcInvestment := utils.NWCInvestment((*periods)[i-1].NWC, (*periods)[i].NWC)

		fcf := utils.UnleveredFCF(utils.FCFParams{
			NWCInvestment: nwcInvestment,
			CapEx:         (*periods)[i].CapEx,
			DA:            (*periods)[i].DA,
			NOPAT:         nopat,
		})

		dp := utils.DiscountingPeriod(utils.BalanceSheetDates{
			Today: currentDate,
			Begin: (*periods)[i].BeginningDate,
			End:   (*periods)[i].EndingDate,
		})

		df := utils.DiscountFactor(utils.DiscountFactorParams{
			PeriodsNumber: dp,
			DiscountRate:  (*periods)[i].DiscountRate,
		})

		dFCF_sum += fcf * df
		terminal_fcf = fcf
	}

	// 2. Calculate terminal discounting periods
	discountingPeriod := utils.DiscountingPeriod(utils.BalanceSheetDates{
		Today: currentDate,
		Begin: (*periods)[periodsNumber-1].BeginningDate,
		End:   (*periods)[periodsNumber-1].EndingDate,
	}) + 0.5

	// 3. Calculate terminal discount factor
	discountFactor := utils.DiscountFactor(utils.DiscountFactorParams{
		PeriodsNumber: discountingPeriod,
		DiscountRate:  (*periods)[periodsNumber-1].DiscountRate,
	})

	// 4. Calculate terminal value
	terminalValue := utils.TerminalValue(utils.TerminalValueParams{
		FCF:          terminal_fcf,
		DiscountRate: (*periods)[periodsNumber-1].DiscountRate,
		PGR:          (*periods)[periodsNumber-1].PGR,
	})

	// 5. Calculate enterprise value
	enterpriseValue := utils.EnterpriseValueFromFCF(utils.EVFromFCFParams{
		TV:                terminalValue,
		DiscreteCashFlows: dFCF_sum,
		DiscountFactor:    discountFactor,
	})

	return enterpriseValue, nil
}
