package utils

import (
	"time"

	"github.com/samgozman/valuation/pkg/metrics"
)

type DiscountFactorFromIntervalsParams struct {
	Today        time.Time // Current or target date
	Begin        time.Time // Balance Sheet Beginning Date
	End          time.Time // Balance Sheet End Date
	DiscountRate float64   // Discount rate (interest rate). Usually equals to WACC
}

// A helper function to calculate the discount factor and discount periods for it
func DiscountFactorFromIntervals(p DiscountFactorFromIntervalsParams) float64 {
	// Calculate terminal discounting period
	discountingPeriod := metrics.DiscountingPeriod(metrics.BalanceSheetDates{
		Today: p.Today,
		Begin: p.Begin,
		End:   p.End,
	}) + 0.5

	discountFactor := metrics.DiscountFactor(metrics.DiscountFactorParams{
		PeriodsNumber: discountingPeriod,
		DiscountRate:  p.DiscountRate,
	})

	return discountFactor
}
