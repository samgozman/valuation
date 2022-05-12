package metrics

import (
	"time"

	"github.com/samgozman/valuation/types"
)

// Calculate sum of discrete cash flows and terminal FCF value
func DiscreteCashFlows(periods *[]types.Period, currentDate time.Time) (dfcf_sum float64, terminal_fcf float64) {
	periodsNumber := len(*periods)

	var sum float64
	var terminal float64
	for i := 1; i < periodsNumber; i++ {
		nopat := NOPAT(NOPATParams{
			EBITDA:      (*periods)[i].EBITDA,
			OtherIncome: (*periods)[i].OtherIncome,
			DA:          (*periods)[i].DA,
			TaxRate:     (*periods)[i].TaxRate,
		})

		nwcInvestment := NWCInvestment((*periods)[i-1].NWC, (*periods)[i].NWC)

		fcf := UnleveredFCF(FCFParams{
			NWCInvestment: nwcInvestment,
			CapEx:         (*periods)[i].CapEx,
			DA:            (*periods)[i].DA,
			NOPAT:         nopat,
		})

		dp := DiscountingPeriod(BalanceSheetDates{
			Today: currentDate,
			Begin: (*periods)[i].BeginningDate,
			End:   (*periods)[i].EndingDate,
		})

		df := DiscountFactor(DiscountFactorParams{
			PeriodsNumber: dp,
			DiscountRate:  (*periods)[i].DiscountRate,
		})

		sum += fcf * df
		terminal = fcf
	}

	return sum, terminal
}
