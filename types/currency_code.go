package types

// CurrencyCode code
type CurrencyCode string

// All currencies
var (
	CurrencyCodeUSD CurrencyCode = "USD"
	CurrencyCodeSGD CurrencyCode = "SGD"
)

func (t CurrencyCode) String() string {
	return string(t)
}
