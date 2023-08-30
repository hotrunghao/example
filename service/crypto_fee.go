package service

import (
	"strconv"
	"strings"
)

func CryptoFeeService(types FeeTypes) Fees {

	var paymentNetwork = strings.ToUpper(types.FromNetwork)
	var customerTier = types.Customer.Tier

	var networkFee = 0.0
	var fee = 0.0

	switch paymentNetwork {
	case `ETH`:
		networkFee = 10
	case `BTC`:
		networkFee = 25
	default:
		networkFee = 2
	}
	switch customerTier {
	case 1:
		fee = networkFee
	case 2:
		fee = networkFee - (0.25 * networkFee)
	case 3:
		if paymentNetwork == `ETH` {
			fee = networkFee - (0.25 * networkFee)
			break
		}
		fee = networkFee - (0.5 * networkFee)
	}
	var amount, _ = strconv.ParseFloat(types.FromAmount, 64)
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
