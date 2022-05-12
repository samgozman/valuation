package metrics

type TerminalValueParams struct {
	FCF          float64 // Free Cash Flow
	DiscountRate float64 // Discount rate (interest rate). Usually equals to WACC
	PGR          float64 // Perpetuity Growth Rate - Typically between the inflation rate of 2-3% and the GDP growth rate of 4-5%
}

type TerminalValueMultiplesParams struct {
	Multiple       float64 // Selected EBITDA Exit Multiple (EV / EBITDA) or Revenue Multiple
	TerminalValue  float64 // Assumed terminal EBITDA or Revenue at the end of the investing cycle
	DiscountFactor float64
}

// Terminal value determines a company's value into perpetuity beyond a set forecast period — usually 5 years
func TerminalValue(p TerminalValueParams) float64 {
	return (p.FCF * (1 + p.PGR)) / (p.DiscountRate - p.PGR)
}

// Terminal value determines a company's value into perpetuity beyond a set forecast period — usually 5 years.
// Calculated from EBITDA or Revenue multiples
func TerminalValueMultiples(p TerminalValueMultiplesParams) float64 {
	return p.TerminalValue * p.Multiple * p.DiscountFactor
}
