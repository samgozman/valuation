package metrics

// Earnings power value (EPV) is a technique for valuing stocks
// by making assumptions about the sustainability of current earnings and the cost of capital.
//
// `adjEarnings` - Adjusted earnings
//
// `wacc` - Weighted average cost of capital
//
// returns - enterprise value
func EPV(adjEarnings int64, wacc float64) (ev int64) {
	return int64(float64(adjEarnings) / wacc)
}
