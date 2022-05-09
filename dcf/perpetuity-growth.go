package dcf

import (
	"errors"
	"time"

	"github.com/samgozman/valuation/pkg/metrics"
	"github.com/samgozman/valuation/types"
)

// Calculate enterprise value by Perpetuity Growth Model
func PerpetuityGrowth(periods *[]types.Period, currentDate time.Time, PGR float32) (float32, error) {
	periodsNumber := len(*periods)

	if periodsNumber < 2 {
		return 0, errors.New("number of periods must be 2+")
	}

	// 1. Calculate discrete Cash Flows
	dFCF_sum, terminal_fcf := metrics.DiscreteCashFlows(periods, currentDate)

	// 2. Calculate terminal discounting periods
	discountingPeriod := metrics.DiscountingPeriod(metrics.BalanceSheetDates{
		Today: currentDate,
		Begin: (*periods)[periodsNumber-1].BeginningDate,
		End:   (*periods)[periodsNumber-1].EndingDate,
	}) + 0.5

	// 3. Calculate terminal discount factor
	discountFactor := metrics.DiscountFactor(metrics.DiscountFactorParams{
		PeriodsNumber: discountingPeriod,
		DiscountRate:  (*periods)[periodsNumber-1].DiscountRate,
	})

	// 4. Calculate terminal value
	terminalValue := metrics.TerminalValue(metrics.TerminalValueParams{
		FCF:          terminal_fcf,
		DiscountRate: (*periods)[periodsNumber-1].DiscountRate,
		PGR:          PGR,
	})

	// 5. Calculate enterprise value
	enterpriseValue := metrics.EnterpriseValueFromFCF(metrics.EVFromFCFParams{
		TV:                terminalValue,
		DiscreteCashFlows: dFCF_sum,
		DiscountFactor:    discountFactor,
	})

	return enterpriseValue, nil
}
