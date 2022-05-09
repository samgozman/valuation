package metrics

import (
	"testing"
)

func TestUnleveredFCF(t *testing.T) {
	t.Run("Should calculate Unlevered FCF", func(t *testing.T) {
		attributes := FCFParams{
			NWCInvestment: 1699,
			CapEx:         16385,
			DA:            15566,
			NOPAT:         139928,
		}

		want := 140808
		got := UnleveredFCF(attributes)

		if want != int(got) {
			t.Errorf("Expected '%v', but got '%v' with attributes: %v", want, got, attributes)
		}
	})
}
