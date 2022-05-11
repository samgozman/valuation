package roe

import (
	"testing"
)

// ! Todo: Add tests for dupont

func TestTrendAnalysis(t *testing.T) {
	t.Run("Should calculate Discounting Period", func(t *testing.T) {
		var tests = []struct {
			values []float32
			want   bool
		}{
			{[]float32{10, 9, 7, 5, 4, 3}, true},
			{[]float32{10, 15, 12, 10, 40, 50}, false},
			{[]float32{1, 2, 3}, false},
			{[]float32{3, 2, 3}, false},
			{[]float32{3, 2, 1}, true},
			{[]float32{10.1, 10.0, 9.9, 9.8}, true},
			{[]float32{0, 0, 0}, false},
			{[]float32{-10, -11, -12}, true},
			{[]float32{-10, -8, -5}, false},
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
