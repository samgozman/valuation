package utils

import (
	"testing"
	"time"
)

func TestDiscountFactorFromIntervals(t *testing.T) {
	t.Run("Should calculate Discounting Factor from intervals", func(t *testing.T) {
		attributes := DiscountFactorFromIntervalsParams{
			Today:        time.Date(2022, 5, 3, 0, 0, 0, 0, time.UTC),
			Begin:        time.Date(2025, 9, 30, 0, 0, 0, 0, time.UTC),
			End:          time.Date(2026, 9, 30, 0, 0, 0, 0, time.UTC),
			DiscountRate: 0.083,
		}

		want := float64(0.70036805)
		got := DiscountFactorFromIntervals(attributes)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with attributes: %v", want, got, attributes)
		}
	})
}
