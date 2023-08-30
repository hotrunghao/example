package service

import (
	"strconv"
	"strings"
)

func FiatFeeService(types FeeTypes) Fees {

	var paymentNetwork = strings.ToUpper(types.FromNetwork)
	var customerTier = types.Customer.Tier

	var networkFee float64 = 0
	var fee float64 = 0
	var amount, _ = strconv.ParseFloat(types.FromAmount, 64)

	switch paymentNetwork {
	case `ACH`:
		var fee = amount * 0.01
		if fee > 2 {
			networkFee = fee
			break
		}
		networkFee = 2
	case `WIRE`:
		networkFee = 25
	case `CARD`:
		networkFee = amount * 0.03

	}
	switch customerTier {
	case 1:
		fee = networkFee
	case 2:
		fee = networkFee - (0.25 * networkFee)
	case 3:
		if paymentNetwork == `WIRE` {
			fee = networkFee - (0.25 * networkFee)
			break
		}
		fee = 0
	}

	var provider = strings.ToUpper(types.Provider.name)
	switch provider {
	case `DUCK`:
		fee = fee + 5 + 0.001*amount
	case `GOOSE`:
		fee = fee + 0.003*amount
	case `FOX`:
		fee = fee + 0.005*amount

	}

	var fees = Fees{}
	fees.Fee = fee
	fees.Provider = types.Provider.name
	fees.Asset = types.FeeAsset
	return fees
}
