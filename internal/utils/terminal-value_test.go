package utils

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

		want := float32(2488417.2)
		got := TerminalValue(attributes)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with attributes: %v", want, got, attributes)
		}
	})
}

func TestTerminalValueEBITDAExit(t *testing.T) {
	t.Run("Should calculate Terminal Value by EBITDA", func(t *testing.T) {
		attributes := TerminalValueEBITDAParams{
			EBITDAExitMultiple: 15,
			TerminalEBITDA:     186210,
			DiscountFactor:     0.7,
		}

		want := float32(1955205)
		got := TerminalValueEBITDAExit(attributes)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with attributes: %v", want, got, attributes)
		}
	})
}
