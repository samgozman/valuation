package utils

type TerminalValueParams struct {
	FCF          float32 // Free Cash Flow
	DiscountRate float32 // Discount rate (interest rate). Usually equals to WACC
	PGR          float32 // Perpetuity Growth Rate - Typically between the inflation rate of 2-3% and the GDP growth rate of 4-5%
}

// Terminal value determines a company's value into perpetuity beyond a set forecast period — usually 5 years
func TerminalValue(p TerminalValueParams) float32 {
	return (p.FCF * (1 + p.PGR)) / (p.DiscountRate - p.PGR)
}
