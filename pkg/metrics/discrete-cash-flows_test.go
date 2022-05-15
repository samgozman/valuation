package metrics

import (
	"testing"
	"time"

	"github.com/samgozman/valuation/internal/utils/testsHelpers"
)

func TestDiscreteCashFlows(t *testing.T) {
	t.Run("Should calculate discrete cash flows", func(t *testing.T) {
		periods := testsHelpers.PeriodsDataSet
		currentDate := time.Date(2022, 5, 4, 0, 0, 0, 0, time.UTC)

		want_sum := 479780
		want_terminal := 136727
		got_sum, got_terminal := DiscreteCashFlows(&periods, currentDate)

		if want_sum != int(got_sum) {
			t.Errorf("Expected '%v', but got '%v' with attributes: %v", want_sum, got_sum, periods)
		}

		if want_terminal != int(got_terminal) {
			t.Errorf("Expected '%v', but got '%v' with attributes: %v", want_terminal, got_terminal, periods)
		}
	})
}
