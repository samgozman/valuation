package metrics

import (
	"testing"
)

func TestEPV(t *testing.T) {
	t.Run("Should calculate EPV", func(t *testing.T) {
		adjEarnings := int64(81917)
		wacc := float64(0.088)

		want := int64(930875)
		got := EPV(adjEarnings, wacc)

		if want != got {
			t.Errorf("Expected '%v', but got '%v'", want, got)
		}
	})
}
