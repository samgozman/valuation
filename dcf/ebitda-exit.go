package dcf

import (
	"errors"
	"time"

	"github.com/samgozman/valuation/internal/utils"
	"github.com/samgozman/valuation/types"
)

func EBITDAExit(periods *[]types.Period, currentDate time.Time, EBITDAExitMultiple float32) (float32, error) {
	periodsNumber := len(*periods)

	if periodsNumber < 2 {
		return 0, errors.New("number of periods must be 2+")
	}

	// 1. Calculate discrete Cash Flows
	dFCF_sum, _ := utils.DiscreteCashFlows(periods, currentDate)

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
		Multiple:       EBITDAExitMultiple,
		TerminalValue:  float32((*periods)[periodsNumber-1].EBITDA),
		DiscountFactor: discountFactor,
	})

	// 5. Calculate enterprise value
	enterpriseValue := terminalValue + dFCF_sum

	return enterpriseValue, nil
}
