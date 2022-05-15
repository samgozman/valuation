package dcf

import (
	"testing"
	"time"

	"github.com/samgozman/valuation/internal/utils/testsHelpers"
	"github.com/samgozman/valuation/types"
)

func TestRevenueExit(t *testing.T) {
	t.Run("Should calculate EV by Revenue Exit strategy", func(t *testing.T) {
		periods := testsHelpers.PeriodsDataSet
		currentDate := time.Date(2022, 5, 4, 0, 0, 0, 0, time.UTC)
		revenueExitMultiple := float64(4.7)

		var want float64 = 2186505.8
		got, _ := RevenueExit(&periods, currentDate, revenueExitMultiple)

		if int(want/10) != int(got/10) {
			t.Errorf("Expected '%v', but got '%v'", want, got)
		}
	})

	t.Run("Should return error if periods number are less than 2", func(t *testing.T) {
		periods := []types.Period{}
		currentDate := time.Date(2022, 5, 4, 0, 0, 0, 0, time.UTC)
		revenueExitMultiple := float64(0)

		_, err := RevenueExit(&periods, currentDate, revenueExitMultiple)

		if err == nil {
			t.Error("Expected error but didn't get it")
		}
	})
}
