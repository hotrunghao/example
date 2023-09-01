package service

import (
	"strings"
)

type Transaction struct {
	FromAmount  string
	FromNetwork string
	FromAsset   string
	ToAsset     string
	ToNetwork   string
	FeeAsset    string
}

type Customer struct {
	Tier int
}

type Provider struct {
	Name string
}

type FeeTypes struct {
	Transaction
	Customer
	Provider
}

type Fees struct {
	Fee      float64
	Asset    string
	Provider string
}

func Fee(types FeeTypes) Fees {
	//assume all fee in USD else we have to convert to the Fee Assert via an API
	var fromNetwork = strings.ToUpper(types.FromNetwork)
	var fromNetworkFee = 0.0
	var toNetworkFee = 0.0
	switch fromNetwork {
	case "CARD", "ACH", "WIRE":
		fromNetworkFee = FiatFeeService(types).Fee
	default:
		fromNetworkFee = CryptoFeeService(types).Fee
	}

	var toNetwork = strings.ToUpper(types.ToNetwork)
	switch toNetwork {
	case "CARD", "ACH", "WIRE":
		toNetworkFee = FiatFeeService(types).Fee
	default:
		toNetworkFee = CryptoFeeService(types).Fee
	}
	var fee = Fees{
		Fee:      fromNetworkFee + toNetworkFee,
		Asset:    types.FeeAsset,
		Provider: types.Provider.Name,
	}
	return fee

}
