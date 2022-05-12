package roe

import "errors"

type DuPontAnalysisParams struct {
	Revenue      int64
	NetIncome    int64
	TotalAssets  int64
	CommonEquity int64
}

type DuPontAnalysisTrends struct {
	IsRoeUpTrend              bool // If ROE is rising most of the time during given periods
	IsNetProfitMarginUpTrend  bool // If NetProfitMargin is rising most of the time during given periods
	IsAssetTurnoverUpTrend    bool // If AssetTurnover is rising most of the time during given periods
	IsEquityMultiplierUpTrend bool // If EquityMultiplier is rising most of the time during given periods
}

type DuPontAnalysisResult struct {
	// Array of ROE for each period
	ROE []float64

	// The net profit margin represents a company's "bottom line" profitability
	// once all expenses have been deducted, including the interest expense payments
	// on debt obligations and taxes to the government.
	NetProfitMargin []float64

	// Total asset turnover ratio is an efficiency ratio tracking the ability
	// of a company to generate more revenue per dollar of asset owned.
	AssetTurnover []float64

	// Equity multiplier is a financial leverage ratio.
	// The Equity Multiplier measures the proportion of a company's assets
	// funded by its equity shareholders as opposed to debt providers.
	//
	// Higher equity multipliers typically signify that the company is utilizing
	// a high percentage of debt in its capital structure to finance working capital needs and asset purchases.
	EquityMultiplier []float64 // LTM (Avg Assets / Avg Equity) for current year

	// Trends for each ROE component
	Trends DuPontAnalysisTrends
}

// Classic DuPont 3 steps analysis.
//
// DuPont Analysis is a framework used to break apart the underlying components
// of the return on equity (ROE) metric to determine the strengths and weaknesses of a company.
//
// Note: Strong companies should have ROE that is increasing because Net Profit and Asset Turnover.
//
// `params` - Direction: from current year to previous
func DuPontAnalysis(params *[]DuPontAnalysisParams) (DuPontAnalysisResult, error) {
	if len(*params) < 2 {
		return DuPontAnalysisResult{}, errors.New("number of periods must be 2+")
	}

	var roeData []float64
	var netProfitMarginData []float64
	var assetTurnoverData []float64
	var equityMultiplierData []float64

	for _, c := range *params {
		netProfitMargin := float64(c.NetIncome) / float64(c.Revenue)
		assetTurnover := float64(c.Revenue) / float64(c.TotalAssets)
		equityMultiplier := float64(c.TotalAssets) / float64(c.CommonEquity)

		roe := netProfitMargin * assetTurnover * equityMultiplier

		roeData = append(roeData, roe)
		netProfitMarginData = append(netProfitMarginData, netProfitMargin)
		assetTurnoverData = append(assetTurnoverData, assetTurnover)
		equityMultiplierData = append(equityMultiplierData, equityMultiplier)
	}

	trends := DuPontAnalysisTrends{
		IsRoeUpTrend:              trendAnalysis(&roeData),
		IsNetProfitMarginUpTrend:  trendAnalysis(&netProfitMarginData),
		IsAssetTurnoverUpTrend:    trendAnalysis(&assetTurnoverData),
		IsEquityMultiplierUpTrend: trendAnalysis(&equityMultiplierData),
	}

	result := DuPontAnalysisResult{
		ROE:              roeData,
		NetProfitMargin:  netProfitMarginData,
		AssetTurnover:    assetTurnoverData,
		EquityMultiplier: equityMultiplierData,
		Trends:           trends,
	}

	return result, nil
}

// Returns true if the trend was up most of the time during the given periods.
//
// Array order: from current year to previous
func trendAnalysis(a *[]float64) bool {
	count := len(*a) - 1
	numberOfTrueValues := 0
	for i := count; i > 0; i-- {
		if (*a)[i-1] > (*a)[i] {
			numberOfTrueValues++
		}
	}

	if float64(numberOfTrueValues)/float64(count) > 0.5 {
		return true
	} else {
		return false
	}
}
