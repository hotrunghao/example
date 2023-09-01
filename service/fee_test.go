package service

import (
	"testing"
)

func TestFee(t *testing.T) {
	var feeType = FeeTypes{
		Transaction: Transaction{
			FromNetwork: "CARD",
			FromAmount:  "100",
			FromAsset:   "USD",
			ToAsset:     "USD",
			FeeAsset:    "USD",
			ToNetwork:   "ETH",
		},
		Customer: Customer{
			Tier: 1,
		},
		Provider: Provider{
			Name: "FOX",
		},
	}
	var fees Fees = Fee(feeType)
	var expected = 2.05
	if fees.Fee != expected {
		t.Errorf("got %f, expected %f", fees.Fee, expected)
	}
}
