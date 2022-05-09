package metrics

// Difference between NWC from previous period between the current (or terminal) period
func NWCInvestment(NWCPrev int64, NWCCurrent int64) int64 {
	return NWCPrev - NWCCurrent
}
