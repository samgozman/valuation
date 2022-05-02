package utils

type FCFParams struct {
	EBITDA        int64   // Earnings before interest, taxes, depreciation, and amortization
	TaxRate       float32 // Tax rate assumption
	NWCInvestment int64   // Difference between NWC from previous period between the current
	CapEx         int64   // Capital Expenditures
	DA            int64   // Depreciation & Amortization

}

// Caclulate FCF based on periodic data for DCF model
func FCF(params FCFParams) float32 {
	ebit := float32(params.EBITDA - params.DA)
	proFormaTaxes := -1 * ebit * params.TaxRate
	nopat := ebit + proFormaTaxes

	return nopat + float32(params.NWCInvestment) + float32(params.DA) - float32(params.CapEx)
}
