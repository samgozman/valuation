package metrics

type FCFParams struct {
	NWCInvestment int64   // Difference between NWC from previous period between the current
	CapEx         int64   // Capital Expenditures
	DA            int64   // Depreciation & Amortization
	NOPAT         float64 // Net Operating Profit After Tax
}

// Calculate Unlevered FCF
func UnleveredFCF(params FCFParams) float64 {
	return params.NOPAT + float64(params.NWCInvestment+params.DA-params.CapEx)
}
