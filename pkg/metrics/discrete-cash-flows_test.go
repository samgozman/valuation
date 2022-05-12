package metrics

import (
	"testing"
	"time"

	"github.com/samgozman/valuation/types"
)

func TestDiscreteCashFlows(t *testing.T) {
	t.Run("Should calculate discrete cash flows", func(t *testing.T) {
		periods := []types.Period{
			{
				NWC: 9975,
			},
			{
				Revenue:       392625,
				EBITDA:        131478,
				OtherIncome:   0,
				CapEx:         11409,
				NWC:           2929,
				DA:            12076,
				TaxRate:       0.17,
				DiscountRate:  0.083,
				BeginningDate: time.Date(2022, 3, 26, 0, 0, 0, 0, time.UTC),
				EndingDate:    time.Date(2022, 9, 30, 0, 0, 0, 0, time.UTC),
			},
			{
				Revenue:       416632,
				EBITDA:        134219,
				OtherIncome:   0,
				CapEx:         12904,
				NWC:           2623,
				DA:            10281,
				TaxRate:       0.17,
				DiscountRate:  0.083,
				BeginningDate: time.Date(2022, 9, 30, 0, 0, 0, 0, time.UTC),
				EndingDate:    time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC),
			},
			{
				Revenue:       430880,
				EBITDA:        139974,
				OtherIncome:   0,
				CapEx:         14168,
				NWC:           1557,
				DA:            14459,
				TaxRate:       0.17,
				DiscountRate:  0.083,
				BeginningDate: time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC),
				EndingDate:    time.Date(2024, 9, 30, 0, 0, 0, 0, time.UTC),
			},
			{
				Revenue:       457900,
				EBITDA:        157853,
				OtherIncome:   0,
				CapEx:         17428,
				NWC:           2952,
				DA:            15037,
				TaxRate:       0.17,
				DiscountRate:  0.083,
				BeginningDate: time.Date(2024, 9, 30, 0, 0, 0, 0, time.UTC),
				EndingDate:    time.Date(2025, 9, 30, 0, 0, 0, 0, time.UTC),
			},
			{
				Revenue:       518374,
				EBITDA:        186210,
				OtherIncome:   0,
				CapEx:         16386,
				NWC:           6607,
				DA:            13027,
				TaxRate:       0.17,
				DiscountRate:  0.083,
				BeginningDate: time.Date(2025, 9, 30, 0, 0, 0, 0, time.UTC),
				EndingDate:    time.Date(2026, 9, 30, 0, 0, 0, 0, time.UTC),
			},
		}
		currentDate := time.Date(2022, 5, 4, 0, 0, 0, 0, time.UTC)

		want_sum := 479780.6842214495
		want_terminal := 136727.88999999998
		got_sum, got_terminal := DiscreteCashFlows(&periods, currentDate)

		if want_sum != got_sum {
			t.Errorf("Expected '%v', but got '%v' with attributes: %v", want_sum, got_sum, periods)
		}

		if want_terminal != got_terminal {
			t.Errorf("Expected '%v', but got '%v' with attributes: %v", want_terminal, got_terminal, periods)
		}
	})
}
