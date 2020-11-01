package types

// ServiceCode service code
type ServiceCode string

// ListServiceCode list services
type ListServiceCode []ServiceCode

// All service codes
var (
	ServiceCodeCODS  ServiceCode = "CODS"  // Cash on Delivery -> Receiver pays the cost of the goods
	ServiceCodeFIRST ServiceCode = "FIRST" // First Delivery -> Committed delivery time at destination country.
	ServiceCodeFRDM  ServiceCode = "FRDM"  // Free Domicile -> Sender pays part/or all of the customs charges and the receiver consignee doesn’t pay anything of the shipping or handling charges.
	ServiceCodeHFPU  ServiceCode = "HFPU"  // Hold for pick up -> Receiver picks up the consignment from an Aramex/Partner facility
	ServiceCodeNOON  ServiceCode = "NOON"  // Noon Delivery -> Committed delivery time BEFORE Noon at destination country
	ServiceCodeSIG   ServiceCode = "SIG"   // Signature Required -> Physical receiver signature required upon delivery
)

func (value ServiceCode) String() string {
	return string(value)
}

// ListString list code
func (value ListServiceCode) ListString() []string {
	var result = []string{}
	for _, v := range value {
		result = append(result, v.String())
	}
	return result
}
