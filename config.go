package aramex

import (
	"github.com/hooklift/gowsdl/soap"
	"github.com/thaitanloi365/go-aramex-sdk/types"
)

type service string

var (
	shippingService service = "shipping"
	rateService     service = "rate"
	trackingService service = "tracking"
	locationService service = "location"
)

// DefaultClientInfo default client info
var DefaultClientInfo = &ClientInfo{
	UserName:           "testingapi@aramex.com",
	Password:           "R123456789$r",
	Version:            "v1",
	AccountNumber:      "20016",
	AccountPin:         "331421",
	AccountEntity:      "AMM",
	AccountCountryCode: "JO",
}

// DefaultSetting default settings
var DefaultSetting = &Settings{
	CompanyName:    "Loi",
	Currency:       types.CurrencyCodeSGD,
	PaymentType:    types.PaymentTypePrepaid,
	ProductGroup:   types.ProductGroupDOM,
	ProductType:    types.ProductTypeDomesticOND,
	PaymentOptions: types.PaymentTypeOptionCASH,
	Services: types.ListServiceCode{
		types.ServiceCodeCODS,
	},
	LabelInfo: &LabelInfo{
		ReportID:   9201,
		ReportType: types.ReportTypeURL,
	},
}

// Settings settings
type Settings struct {
	CompanyName    string
	Currency       types.CurrencyCode
	ProductGroup   types.ProductGroup
	ProductType    types.ProductType
	PaymentType    types.PaymentType
	PaymentOptions types.PaymentTypeOption
	Services       types.ListServiceCode
	LabelInfo      *LabelInfo
}

// Logger logger
type Logger interface {
	Printf(format string, args ...interface{})
}

// Config config
type Config struct {
	IsLive         bool
	ClientInfo     *ClientInfo
	Logger         Logger
	DefaultSetting *Settings
}

func (config *Config) getEndpoint(service service) (endpoint string) {
	switch service {
	case shippingService:
		endpoint = "https://ws.dev.aramex.net/shippingapi.v2/shipping/service_1_0.svc?wsdl"
		if config.IsLive {
			endpoint = "https://ws.aramex.net/shippingapi.v2/shipping/service_1_0.svc?wsdl"
		}
	case rateService:
		endpoint = "https://ws.dev.aramex.net/shippingapi.v2/ratecalculator/service_1_0.svc?wsdl"
		if config.IsLive {
			endpoint = "https://ws.aramex.net/shippingapi.v2/ratecalculator/service_1_0.svc?wsdl"
		}
	case trackingService:
		endpoint = "https://ws.dev.aramex.net/shippingapi.v2/tracking/service_1_0.svc?wsdl"
		if config.IsLive {
			endpoint = "https://ws.aramex.net/shippingapi.v2/tracking/service_1_0.svc?wsdl"
		}
	case locationService:
		endpoint = "https://ws.dev.aramex.net/shippingapi.v2/location/service_1_0.svc?wsdl"
		if config.IsLive {
			endpoint = "https://ws.aramex.net/shippingapi.v2/location/service_1_0.svc?wsdl"
		}
	}

	return
}

func (config *Config) getSoapClient(service service) *soap.Client {
	var endpoint = config.getEndpoint(service)
	var client = soap.NewClient(endpoint)

	return client
}
