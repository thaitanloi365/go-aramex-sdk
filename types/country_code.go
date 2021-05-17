package types

// CountryCode string
type CountryCode string

// All country code
// https://www.aramex.com/docs/default-source/resourses/resourcesdata/shipping-services-api-manual.pdf
// Page 52 to 54
var (
	CountryCodeSG CountryCode = "SG" // Singapore
	CountryCodeUS CountryCode = "US" // United States
	CountryCodeAD CountryCode = "AD" // Andorra
	CountryCodeAE CountryCode = "AE" // United Arab Emirates
	CountryCodeJO CountryCode = "JO" // Jordan
	CountryCodeAF CountryCode = "AF" // Afghanistan
	CountryCodeAG CountryCode = "AG" // Antigua and Barbuda
	CountryCodePH CountryCode = "PH" // Philipines
	CountryCodeMY CountryCode = "MY" // Malaysia
	CountryCodeVN CountryCode = "VN" // Vietnam
	CountryCodeHK CountryCode = "HK" // Hong Kong
	CountryCodeID CountryCode = "HK" // Indonesia
	CountryCodeTH CountryCode = "TH" // Thailand

)

func (cc CountryCode) String() string {
	return string(cc)
}
