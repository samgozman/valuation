package dcf

import (
	"testing"
	"time"
)

func TestPerpetuityGrowth(t *testing.T) {
	t.Run("Should calculate EV by PerpetuityGrowth", func(t *testing.T) {
		periods := []PerpetuityGrowthPeriod{
			{
				NWC: 9975,
			},
			{
				Revenue:       392625,
				EBITDA:        131478,
				CapEx:         11409,
				NWC:           2929,
				DA:            12076,
				TaxRate:       0.17,
				DiscountRate:  0.083,
				PGR:           0.025,
				BeginningDate: time.Date(2022, 3, 26, 0, 0, 0, 0, time.UTC),
				EndingDate:    time.Date(2022, 9, 30, 0, 0, 0, 0, time.UTC),
			},
			{
				Revenue:       416632,
				EBITDA:        134219,
				CapEx:         12904,
				NWC:           2623,
				DA:            10281,
				TaxRate:       0.17,
				DiscountRate:  0.083,
				PGR:           0.025,
				BeginningDate: time.Date(2022, 9, 30, 0, 0, 0, 0, time.UTC),
				EndingDate:    time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC),
			},
			{
				Revenue:       430880,
				EBITDA:        139974,
				CapEx:         14168,
				NWC:           1557,
				DA:            14459,
				TaxRate:       0.17,
				DiscountRate:  0.083,
				PGR:           0.025,
				BeginningDate: time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC),
				EndingDate:    time.Date(2024, 9, 30, 0, 0, 0, 0, time.UTC),
			},
			{
				Revenue:       457900,
				EBITDA:        157853,
				CapEx:         17428,
				NWC:           2952,
				DA:            15037,
				TaxRate:       0.17,
				DiscountRate:  0.083,
				PGR:           0.025,
				BeginningDate: time.Date(2024, 9, 30, 0, 0, 0, 0, time.UTC),
				EndingDate:    time.Date(2025, 9, 30, 0, 0, 0, 0, time.UTC),
			},
			{
				Revenue:       518374,
				EBITDA:        186210,
				CapEx:         16386,
				NWC:           6607,
				DA:            13027,
				TaxRate:       0.17,
				DiscountRate:  0.083,
				PGR:           0.025,
				BeginningDate: time.Date(2025, 9, 30, 0, 0, 0, 0, time.UTC),
				EndingDate:    time.Date(2026, 9, 30, 0, 0, 0, 0, time.UTC),
			},
		}

		want := 2206649
		got, _ := PerpetuityGrowth(&periods)

		if want != int(got) {
			t.Errorf("Expected '%v', but got '%v'", want, got)
		}
	})
}