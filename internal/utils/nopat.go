package utils

type NOPATParams struct {
	EBITDA  int64   // Earnings before interest, taxes, depreciation, and amortization
	DA      int64   // Depreciation & Amortization
	TaxRate float32 // Tax rate assumption
}

// Net Operating Profit After Tax
func NOPAT(params NOPATParams) float32 {
	ebit := float32(params.EBITDA - params.DA)
	proFormaTaxes := -1 * ebit * params.TaxRate
	return ebit + proFormaTaxes
}
