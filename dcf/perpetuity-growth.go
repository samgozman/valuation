package dcf

// TODO: Rename, refactor
type PerpetuityGrowthPeriod struct {
	Revenue      int64
	EBITDA       int64   // Earnings before interest, taxes, depreciation, and amortization
	CapEx        int64   // Capital Expenditures
	NWC          int64   // Net Working Capital
	DA           int64   // Depreciation & Amortization
	TaxRate      float32 // Tax rate assumption
	DiscountRate float32 // Discount rate (interest rate). Usually equals to WACC
	PGR          float32 // Perpetuity Growth Rate - Typically between the inflation rate of 2-3% and the GDP growth rate of 4-5%
}

// Perpetuity Growth Model
func PerpetuityGrowthFairValue(periods *[]PerpetuityGrowthPeriod) {

}
