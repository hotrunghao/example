package service

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
	name string
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
