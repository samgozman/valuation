package dcf

import (
	"errors"
	"time"

	"github.com/samgozman/valuation/internal/utils"
	"github.com/samgozman/valuation/pkg/metrics"
	"github.com/samgozman/valuation/types"
)

// Calculate enterprise value by EBITDA	Exit Model
func EBITDAExit(periods *[]types.Period, currentDate time.Time, EBITDAExitMultiple float64) (float64, error) {
	periodsNumber := len(*periods)

	if periodsNumber < 2 {
		return 0, errors.New("number of periods must be 2+")
	}

	// 1. Calculate discrete Cash Flows
	dFCF_sum, _ := metrics.DiscreteCashFlows(periods, currentDate)

	// 2. Calculate terminal discount factor
	discountFactor := utils.DiscountFactorFromIntervals(utils.DiscountFactorFromIntervalsParams{
		Today:        currentDate,
		Begin:        (*periods)[periodsNumber-1].BeginningDate,
		End:          (*periods)[periodsNumber-1].EndingDate,
		DiscountRate: (*periods)[periodsNumber-1].DiscountRate,
	})

	// 3. Calculate terminal value
	terminalValue := metrics.TerminalValueMultiples(metrics.TerminalValueMultiplesParams{
		Multiple:       EBITDAExitMultiple,
		TerminalValue:  float64((*periods)[periodsNumber-1].EBITDA),
		DiscountFactor: discountFactor,
	})

	// 4. Calculate enterprise value
	enterpriseValue := terminalValue + dFCF_sum

	return enterpriseValue, nil
}
