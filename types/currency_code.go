package types

// CurrencyCode code
type CurrencyCode string

// All currencies
var (
	CurrencyCodeUSD CurrencyCode = "USD" // United States Dollar
	CurrencyCodeSGD CurrencyCode = "SGD"
	CurrencyCodeVND CurrencyCode = "VND"
	CurrencyCodeHKD CurrencyCode = "HKD" // Hong Kong Dollar
	CurrencyCodeIDR CurrencyCode = "IDR" // Indonesian Rupiah
	CurrencyCodeMYR CurrencyCode = "MYR" // Malaysia Ringgit
	CurrencyCodePHP CurrencyCode = "PHP" // Philippine Peso
)

func (t CurrencyCode) String() string {
	return string(t)
}
