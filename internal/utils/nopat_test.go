package utils

import (
	"testing"
)

func TestNOPAT(t *testing.T) {
	t.Run("Should calculate NOPAT", func(t *testing.T) {
		attributes := NOPATParams{
			EBITDA:  186210,
			DA:      15566,
			TaxRate: 0.18,
		}

		want := 139928
		got := NOPAT(attributes)

		if want != int(got) {
			t.Errorf("Expected '%v', but got '%v' with attributes: %v", want, got, attributes)
		}
	})
}
