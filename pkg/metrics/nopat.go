package metrics

type NOPATParams struct {
	EBITDA      int64   // Earnings before interest, taxes, depreciation, and amortization
	OtherIncome int64   // Other Income / Expenses
	DA          int64   // Depreciation & Amortization
	TaxRate     float64 // Tax rate assumption
}

// Net Operating Profit After Tax
func NOPAT(params NOPATParams) float64 {
	ebit := float64(params.EBITDA + params.OtherIncome - params.DA)
	proFormaTaxes := ebit * params.TaxRate
	return ebit - proFormaTaxes
}
