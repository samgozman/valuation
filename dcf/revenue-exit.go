package dcf

import (
	"errors"
	"time"

	"github.com/samgozman/valuation/internal/utils"
)

func RevenueExit(periods *[]Period, currentDate time.Time, revenueExitMultiple float32) (float32, error) {
	periodsNumber := len(*periods)

	if periodsNumber < 2 {
		return 0, errors.New("number of periods must be 2+")
	}

	// 1. Calculate discrete Cash Flows
	var dFCF_sum float32
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
	terminalValue := utils.TerminalValueMultiples(utils.TerminalValueMultiplesParams{
		Multiple:       revenueExitMultiple,
		TerminalValue:  float32((*periods)[periodsNumber-1].Revenue),
		DiscountFactor: discountFactor,
	})

	// 5. Calculate enterprise value
	enterpriseValue := terminalValue + dFCF_sum

	return enterpriseValue, nil
}
