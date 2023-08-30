package service

import (
	"testing"
)

func TestCryptoFeeService(t *testing.T) {
	var feetypes_eth_customer_tier_1 = FeeTypes{
		Transaction: Transaction{
			FromAmount:  "10",
			FromNetwork: "ETH",
			FromAsset:   "ETH",
			ToAsset:     "ETH",
		},
		Customer: Customer{
			Tier: 1,
		},
		Provider: Provider{name: "Goose"},
	}
	var result = CryptoFeeService(feetypes_eth_customer_tier_1)
	var expected = 10.03
	if result.Fee != expected {
		t.Errorf("got %f, expected %f", result.Fee, expected)
	}

	var feetypes_btc_customer_tier_1 = FeeTypes{
		Transaction: Transaction{
			FromAmount:  "10",
			FromNetwork: "BTC",
			FromAsset:   "BTC",
			ToAsset:     "BTC",
		},
		Customer: Customer{
			Tier: 2,
		},
		Provider: Provider{name: "Duck"},
	}
	result = CryptoFeeService(feetypes_btc_customer_tier_1)
	expected = 23.76
	if result.Fee != expected {
		t.Errorf("got %f, expected %f", result.Fee, expected)
	}

	var feetypes_sol_customer_tier_1 = FeeTypes{
		Transaction: Transaction{
			FromAmount:  "10",
			FromNetwork: "SOLANA",
			FromAsset:   "SOL",
			ToAsset:     "SOL",
		},
		Customer: Customer{
			Tier: 3,
		},
		Provider: Provider{name: "Fox"},
	}
	result = CryptoFeeService(feetypes_sol_customer_tier_1)
	expected = 1.05
	if result.Fee != expected {
		t.Errorf("got %f, expected %f", result.Fee, expected)
	}

}
