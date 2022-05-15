package dcf

import (
	"testing"
	"time"

	"github.com/samgozman/valuation/internal/utils/testsHelpers"
	"github.com/samgozman/valuation/types"
)

func TestPerpetuityGrowth(t *testing.T) {
	t.Run("Should calculate EV by PerpetuityGrowth", func(t *testing.T) {
		periods := testsHelpers.PeriodsDataSet
		currentDate := time.Date(2022, 5, 4, 0, 0, 0, 0, time.UTC)
		pgr := float64(0.025)

		var want float64 = 2172463.2
		got, _ := PerpetuityGrowth(&periods, currentDate, pgr)

		if int(want/10) != int(got/10) {
			t.Errorf("Expected '%v', but got '%v'", want, got)
		}
	})

	t.Run("Should return error if periods number are less than 2", func(t *testing.T) {
		periods := []types.Period{}
		currentDate := time.Date(2022, 5, 4, 0, 0, 0, 0, time.UTC)
		pgr := float64(0)

		_, err := PerpetuityGrowth(&periods, currentDate, pgr)

		if err == nil {
			t.Error("Expected error but didn't get it")
		}
	})
}
