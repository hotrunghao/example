package main

import (
	"fmt"
	"service/service"
)

func main() {

	var feeType = service.FeeTypes{
		Transaction: service.Transaction{
			FromNetwork: "CARD",
			FromAmount:  "100",
			FromAsset:   "USD",
			ToAsset:     "USD",
			FeeAsset:    "USD",
			ToNetwork:   "ETH",
		},
		Customer: service.Customer{
			Tier: 1,
		},
		Provider: service.Provider{
			Name: "FOX",
		},
	}

	var fee = service.Fee(feeType)
	fmt.Println("Fee ", fee.Fee)
	fmt.Println("Asset ", fee.Asset)
	fmt.Println("Provider ", fee.Provider)

}
