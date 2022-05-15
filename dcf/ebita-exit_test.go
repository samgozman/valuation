package dcf

import (
	"testing"
	"time"

	"github.com/samgozman/valuation/internal/utils/testsHelpers"
	"github.com/samgozman/valuation/types"
)

func TestEBITDAExit(t *testing.T) {
	t.Run("Should calculate EV by EBITDA Exit strategy", func(t *testing.T) {
		periods := testsHelpers.PeriodsDataSet
		currentDate := time.Date(2022, 5, 4, 0, 0, 0, 0, time.UTC)
		EBITDAExitMultiple := float64(15.1)

		var want float64 = 2449491.5
		got, _ := EBITDAExit(&periods, currentDate, EBITDAExitMultiple)

		if int(want/10) != int(got/10) {
			t.Errorf("Expected '%v', but got '%v'", want, got)
		}
	})

	t.Run("Should return error if periods number are less than 2", func(t *testing.T) {
		periods := []types.Period{}
		currentDate := time.Date(2022, 5, 4, 0, 0, 0, 0, time.UTC)
		EBITDAExitMultiple := float64(0)

		_, err := EBITDAExit(&periods, currentDate, EBITDAExitMultiple)

		if err == nil {
			t.Error("Expected error but didn't get it")
		}
	})
}
