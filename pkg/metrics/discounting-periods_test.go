package metrics

import (
	"testing"
	"time"
)

func TestDiscountingPeriod(t *testing.T) {
	t.Run("Should calculate Discounting Period", func(t *testing.T) {
		attributes := BalanceSheetDates{
			Today: time.Date(2022, 5, 3, 0, 0, 0, 0, time.UTC),
			Begin: time.Date(2022, 3, 26, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, 9, 30, 0, 0, 0, 0, time.UTC),
		}

		want := float32(0.15555556)
		got := DiscountingPeriod(attributes)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with attributes: %v", want, got, attributes)
		}
	})
}

func TestDaysBetween(t *testing.T) {
	t.Run("Should calculate Discounting Period", func(t *testing.T) {
		endDate := time.Date(2022, 9, 30, 0, 0, 0, 0, time.UTC)
		startDate := time.Date(2022, 3, 26, 0, 0, 0, 0, time.UTC)

		want := int64(188)
		got := daysBetween(endDate, startDate)

		if want != got {
			t.Errorf("Expected '%v', but got '%v'", want, got)
		}
	})
}

func TestDiscountFactor(t *testing.T) {
	t.Run("Should calculate Discounting Factor", func(t *testing.T) {
		attributes := DiscountFactorParams{
			PeriodsNumber: 4.48,
			DiscountRate:  0.083,
		}

		want := float32(0.6996239)
		got := DiscountFactor(attributes)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with attributes: %v", want, got, attributes)
		}
	})
}
