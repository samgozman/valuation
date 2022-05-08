package types

import "time"

// First index is always the current state of balance sheet.
// If you want to calculate DCF for 5 years when you need to
// pass 6 elements of Period into function.
// First element will be the current state of balance sheet
// and the other 5 will represent your assumptions.
type Period struct {
	Revenue       int64
	EBITDA        int64     // Earnings before interest, taxes, depreciation and amortization
	OtherIncome   int64     // Other Income / Expenses
	CapEx         int64     // Capital Expenditures
	NWC           int64     // Net Working Capital
	DA            int64     // Depreciation & Amortization
	TaxRate       float32   // Tax rate assumption
	DiscountRate  float32   // Discount rate (interest rate). Usually equals to WACC
	BeginningDate time.Time // Balance Sheet beginning date
	EndingDate    time.Time // Balance Sheet ending date
}
