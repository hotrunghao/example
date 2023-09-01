package service

import (
	"testing"
)

func TestFiatFeeService(t *testing.T) {
	var feetypes_ach_customer_tier_1 = FeeTypes{
		Transaction: Transaction{
			FromAmount:  "10",
			FromNetwork: "ach",
			FromAsset:   "USD",
			ToAsset:     "USD",
		},
		Customer: Customer{
			Tier: 1,
		},
	}
	var result = FiatFeeService(feetypes_ach_customer_tier_1)
	var expected float64 = 2
	if result.Fee != expected {
		t.Errorf("got %f, expected %f", result.Fee, expected)
	}

	var feetypes_ach_customer_tier_2 = FeeTypes{
		Transaction: Transaction{
			FromAmount:  "500",
			FromNetwork: "ACH",
			FromAsset:   "USD",
			ToAsset:     "USD",
		},
		Customer: Customer{
			Tier: 2,
		},
	}
	result = FiatFeeService(feetypes_ach_customer_tier_2)
	expected = 3.75
	if result.Fee != expected {
		t.Errorf("got %f, expected %f", result.Fee, expected)
	}

	var feetypes_ach_customer_tier_3 = FeeTypes{
		Transaction: Transaction{
			FromAmount:  "500",
			FromNetwork: "ACH",
			FromAsset:   "USD",
			ToAsset:     "USD",
		},
		Customer: Customer{
			Tier: 3,
		},
	}
	result = FiatFeeService(feetypes_ach_customer_tier_3)
	expected = 0.0
	if result.Fee != expected {
		t.Errorf("got %f, expected %f", result.Fee, expected)
	}

	var feetypes_wire_customer_tier_3 = FeeTypes{
		Transaction: Transaction{
			FromAmount:  "500",
			FromNetwork: "WIRE",
			FromAsset:   "USD",
			ToAsset:     "USD",
		},
		Customer: Customer{
			Tier: 3,
		},
	}
	result = FiatFeeService(feetypes_wire_customer_tier_3)
	expected = 18.75
	if result.Fee != expected {
		t.Errorf("got %f, expected %f", result.Fee, expected)
	}

	var feetypes_card_customer_tier_3 = FeeTypes{
		Transaction: Transaction{
			FromAmount:  "500",
			FromNetwork: "CARD",
			FromAsset:   "USD",
			ToAsset:     "USD",
		},
		Customer: Customer{
			Tier: 3,
		},
	}
	result = FiatFeeService(feetypes_card_customer_tier_3)
	expected = 0.0
	if result.Fee != expected {
		t.Errorf("got %f, expected %f", result.Fee, expected)
	}

	var feetypes_ach_customer_tier_1_provider_fox = FeeTypes{
		Transaction: Transaction{
			FromAmount:  "10",
			FromNetwork: "ach",
			FromAsset:   "USD",
			ToAsset:     "USD",
		},
		Customer: Customer{
			Tier: 1,
		},
		Provider: Provider{Name: "FOX"},
	}

	result = FiatFeeService(feetypes_ach_customer_tier_1_provider_fox)
	expected = 2.05
	if result.Fee != expected {
		t.Errorf("got %f, expected %f", result.Fee, expected)
	}

	var feetypes_ach_customer_tier_1_provider_goose = FeeTypes{
		Transaction: Transaction{
			FromAmount:  "10",
			FromNetwork: "ach",
			FromAsset:   "USD",
			ToAsset:     "USD",
		},
		Customer: Customer{
			Tier: 1,
		},
		Provider: Provider{Name: "Goose"},
	}

	result = FiatFeeService(feetypes_ach_customer_tier_1_provider_goose)
	expected = 2.03
	if result.Fee != expected {
		t.Errorf("got %f, expected %f", result.Fee, expected)
	}

	var feetypes_ach_customer_tier_1_provider_duck = FeeTypes{
		Transaction: Transaction{
			FromAmount:  "10",
			FromNetwork: "ach",
			FromAsset:   "USD",
			ToAsset:     "USD",
		},
		Customer: Customer{
			Tier: 1,
		},
		Provider: Provider{Name: "duck"},
	}

	result = FiatFeeService(feetypes_ach_customer_tier_1_provider_duck)
	expected = 7.01
	if result.Fee != expected {
		t.Errorf("got %f, expected %f", result.Fee, expected)
	}
}
