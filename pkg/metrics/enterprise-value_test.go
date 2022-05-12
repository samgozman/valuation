package metrics

import (
	"testing"
)

func TestEnterpriseValueFromFCF(t *testing.T) {
	t.Run("Should calculate Enterprise Value from FCF", func(t *testing.T) {
		attributes := EVFromFCFParams{
			TV:                2488417.2,
			DiscreteCashFlows: 424397,
			DiscountFactor:    0.7,
		}

		want := float64(2166289)
		got := EnterpriseValueFromFCF(attributes)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with attributes: %v", want, got, attributes)
		}
	})
}
