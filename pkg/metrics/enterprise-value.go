package metrics

type EVFromFCFParams struct {
	TV                float32 // Terminal value
	DiscreteCashFlows float32 // Sum of discrete cashflows
	DiscountFactor    float32
}

func EnterpriseValueFromFCF(params EVFromFCFParams) float32 {
	discountedTV := params.TV * params.DiscountFactor
	return discountedTV + params.DiscreteCashFlows
}
