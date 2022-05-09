package dcf

import (
	"errors"
	"time"

	"github.com/samgozman/valuation/pkg/metrics"
	"github.com/samgozman/valuation/types"
)

// Calculate enterprise value by Revenue Exit model
func RevenueExit(periods *[]types.Period, currentDate time.Time, revenueExitMultiple float32) (float32, error) {
	periodsNumber := len(*periods)

	if periodsNumber < 2 {
		return 0, errors.New("number of periods must be 2+")
	}

	// 1. Calculate discrete Cash Flows
	dFCF_sum, _ := metrics.DiscreteCashFlows(periods, currentDate)

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
	terminalValue := metrics.TerminalValueMultiples(metrics.TerminalValueMultiplesParams{
		Multiple:       revenueExitMultiple,
		TerminalValue:  float32((*periods)[periodsNumber-1].Revenue),
		DiscountFactor: discountFactor,
	})

	// 5. Calculate enterprise value
	enterpriseValue := terminalValue + dFCF_sum

	return enterpriseValue, nil
}
