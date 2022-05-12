package metrics

type EVFromFCFParams struct {
	TV                float64 // Terminal value
	DiscreteCashFlows float64 // Sum of discrete cashflows
	DiscountFactor    float64
}

func EnterpriseValueFromFCF(params EVFromFCFParams) float64 {
	discountedTV := params.TV * params.DiscountFactor
	return discountedTV + params.DiscreteCashFlows
}
