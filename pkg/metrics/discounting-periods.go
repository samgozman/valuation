package metrics

import (
	"math"
	"time"
)

type BalanceSheetDates struct {
	Today time.Time // Current or target date
	Begin time.Time // Balance Sheet Beginning Date
	End   time.Time // Balance Sheet End Date
}

type DiscountFactorParams struct {
	PeriodsNumber float64
	DiscountRate  float64 // Discount rate (interest rate). Usually equals to WACC
}

func DiscountingPeriod(d BalanceSheetDates) float64 {
	daysBetweenBeginAndEnd := daysBetween(d.End, d.Begin)
	midYearConvention := d.Begin.Add(time.Duration(daysBetweenBeginAndEnd/2) * 24 * time.Hour)
	daysBetweenMidYearAndToday := daysBetween(midYearConvention, d.Today)

	return float64(daysBetweenMidYearAndToday) / float64(360)
}

func DiscountFactor(params DiscountFactorParams) float64 {
	return 1 / math.Pow(1+params.DiscountRate, params.PeriodsNumber)
}

func daysBetween(d1 time.Time, d2 time.Time) int64 {
	return int64(d1.Sub(d2).Hours() / 24)
}
