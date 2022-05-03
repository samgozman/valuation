package utils

import (
	"time"
)

type BalanceSheetDates struct {
	Today time.Time // Current or target date
	Begin time.Time // Balance Sheet Beginning Date
	End   time.Time // Balance Sheet End Date
}

func DiscountingPeriod(d BalanceSheetDates) float32 {
	daysBetweenBeginAndEnd := daysBetween(d.End, d.Begin)
	midYearConvention := d.Begin.Add(time.Duration(daysBetweenBeginAndEnd/2) * 24 * time.Hour)
	daysBetweenMidYearAndToday := daysBetween(midYearConvention, d.Today)

	return float32(daysBetweenMidYearAndToday) / 360
}

func daysBetween(d1 time.Time, d2 time.Time) int64 {
	return int64(d1.Sub(d2).Hours() / 24)
}
