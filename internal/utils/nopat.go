package utils

type NOPATParams struct {
	EBITDA      int64   // Earnings before interest, taxes, depreciation, and amortization
	OtherIncome int64   // Other Income / Expenses
	DA          int64   // Depreciation & Amortization
	TaxRate     float32 // Tax rate assumption
}

// Net Operating Profit After Tax
func NOPAT(params NOPATParams) float32 {
	ebit := float32(params.EBITDA + params.OtherIncome - params.DA)
	proFormaTaxes := ebit * params.TaxRate
	return ebit - proFormaTaxes
}
