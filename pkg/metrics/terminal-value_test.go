package metrics

import (
	"testing"
)

func TestTerminalValue(t *testing.T) {
	t.Run("Should calculate Terminal Value", func(t *testing.T) {
		attributes := TerminalValueParams{
			FCF:          140808,
			DiscountRate: 0.083,
			PGR:          0.025,
		}

		want := float64(2488417.2)
		got := TerminalValue(attributes)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with attributes: %v", want, got, attributes)
		}
	})
}

func TestTerminalValueMultiples(t *testing.T) {
	t.Run("Should calculate Terminal Value by EBITDA", func(t *testing.T) {
		attributes := TerminalValueMultiplesParams{
			Multiple:       15,
			TerminalValue:  186210,
			DiscountFactor: 0.7,
		}

		want := float64(1955205)
		got := TerminalValueMultiples(attributes)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with attributes: %v", want, got, attributes)
		}
	})
}
