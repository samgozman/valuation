package roe

import (
	"testing"
)

func TestDuPontAnalysis(t *testing.T) {
	t.Run("Should return full DuPont analysis", func(t *testing.T) {
		attributes := []DuPontAnalysisParams{
			{365817, 94680, 351002, 63090}, // current
			{274515, 57411, 323888, 65339},
			{260174, 55256, 338516, 90488},
			{265595, 59531, 365725, 107147},
			{229234, 48351, 375319, 134047},
			{215639, 45687, 321686, 128249},
		}

		expectedTrends := DuPontAnalysisTrends{
			IsRoeUpTrend:              true,
			IsNetProfitMarginUpTrend:  false,
			IsAssetTurnoverUpTrend:    true,
			IsEquityMultiplierUpTrend: true,
		}
		want := DuPontAnalysisResult{
			ROE:              []float64{1.5, 0.87, 0.61, 0.55, 0.36, 0.35},
			NetProfitMargin:  []float64{0.25, 0.20, 0.21, 0.22, 0.21, 0.21},
			AssetTurnover:    []float64{1.04, 0.84, 0.76, 0.72, 0.61, 0.67},
			EquityMultiplier: []float64{5.56, 4.95, 3.74, 3.41, 2.79, 2.50},
			Trends:           expectedTrends,
		}
		got, _ := DuPontAnalysis(&attributes)

		if expectedTrends != got.Trends {
			t.Errorf("Expected '%v', but got '%v' with attributes: %v", expectedTrends, got.Trends, attributes)
		}

		compareArrays(&want.ROE, &got.ROE, t)
		compareArrays(&want.NetProfitMargin, &got.NetProfitMargin, t)
		compareArrays(&want.AssetTurnover, &got.AssetTurnover, t)
		compareArrays(&want.EquityMultiplier, &got.EquityMultiplier, t)
	})

	t.Run("Should return error if elements are not enough", func(t *testing.T) {
		attributes := []DuPontAnalysisParams{
			{365817, 94680, 351002, 63090},
		}
		_, error := DuPontAnalysis(&attributes)

		if error == nil {
			t.Error("Expected error but didn't get it")
		}
	})
}

func TestTrendAnalysis(t *testing.T) {
	t.Run("Should return expected trend analysis result", func(t *testing.T) {
		var tests = []struct {
			values []float64
			want   bool
		}{
			{[]float64{10, 9, 7, 5, 4, 3}, true},
			{[]float64{10, 15, 12, 10, 40, 50}, false},
			{[]float64{1, 2, 3}, false},
			{[]float64{3, 2, 3}, false},
			{[]float64{3, 2, 1}, true},
			{[]float64{10.1, 10.0, 9.9, 9.8}, true},
			{[]float64{0, 0, 0}, false},
			{[]float64{-10, -11, -12}, true},
			{[]float64{-10, -8, -5}, false},
		}

		for _, tt := range tests {
			want := tt.want
			got := trendAnalysis(&tt.values)

			if want != got {
				t.Errorf("Expected '%v', but got '%v' with values: %v", want, got, &tt.values)
			}
		}
	})
}

func compareArrays(a *[]float64, b *[]float64, t *testing.T) {
	if len(*a) != len(*b) {
		t.Error("Arrays are not the same length")
	}

	for i := 0; i < len(*a); i++ {
		t.Run("Values should be equal", func(t *testing.T) {
			want := int((*a)[i] * 100)
			got := int((*b)[i] * 100)
			if want != got {
				t.Errorf("Expected '%v', but got '%v'", want, got)
			}
		})
	}
}
