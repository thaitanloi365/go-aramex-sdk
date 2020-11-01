package types

// PaymentType shipment payment method type
type PaymentType string

// All payment types
var (
	PaymentTypePrepaid    PaymentType = "P"
	PaymentTypeCollect    PaymentType = "C"
	PaymentTypeThirdParty PaymentType = "3"
)

// PaymentTypeOption option
type PaymentTypeOption string

// All payment options for PaymenType = C
var (
	PaymentTypeOptionASCC PaymentTypeOption = "ASCC" // Needs Shipper Account Number to be filled.
	PaymentTypeOptionARCC PaymentTypeOption = "ARCC" // Needs Consignee Account Number to be filled.
)

// All payment options for PaymenType = P
var (
	PaymentTypeOptionCASH PaymentTypeOption = "CASH" // Cash
	PaymentTypeOptionACCT PaymentTypeOption = "ACCT" // Account
	PaymentTypeOptionPPST PaymentTypeOption = "PPST" // Prepaid Stock
	PaymentTypeOptionCRDT PaymentTypeOption = "CRDT" // Credit

)

func (t PaymentType) String() string {
	return string(t)
}

func (t PaymentTypeOption) String() string {
	return string(t)
}
