package utils

type FCFParams struct {
	NWCInvestment int64   // Difference between NWC from previous period between the current
	CapEx         int64   // Capital Expenditures
	DA            int64   // Depreciation & Amortization
	NOPAT         float32 // Net Operating Profit After Tax
}

// Caclulate Unlevered FCF
func UnleveredFCF(params FCFParams) float32 {
	return params.NOPAT + float32(params.NWCInvestment) + float32(params.DA) - float32(params.CapEx)
}
