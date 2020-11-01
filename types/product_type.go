package types

// ProductGroup group
type ProductGroup string

// All production groups
var (
	ProductGroupEXP ProductGroup = "EXP"
	ProductGroupDOM ProductGroup = "DOM"
)

// ProductType type
type ProductType string

// All production types of express product group
var (
	ProductTypeExpressPDX ProductType = "PDX" // Priority Document Express -> No Urgent, time sensitive consignments containing printed matter or document material
	ProductTypeExpressPPX ProductType = "PPX" // Priority Parcel Express -> Yes Urgent, time sensitive consignments containing non-printed matter or non-document material
	ProductTypeExpressPLX ProductType = "PLX" // Priority Parcel Express -> Yes Urgent, time sensitive consignments containing non-printed matter or non-document material
	ProductTypeExpressDDX ProductType = "DDX" // Priority Parcel Express -> Yes Urgent, time sensitive consignments containing non-printed matter or non-document material
	ProductTypeExpressDPX ProductType = "DPX" // Priority Parcel Express -> Yes Urgent, time sensitive consignments containing non-printed matter or non-document material
	ProductTypeExpressGDX ProductType = "GDX" // Priority Parcel Express -> Yes Urgent, time sensitive consignments containing non-printed matter or non-document material
	ProductTypeExpressGPX ProductType = "GPX" // Priority Parcel Express -> Yes Urgent, time sensitive consignments containing non-printed matter or non-document material
	ProductTypeExpressEPX ProductType = "EPX" // Priority Parcel Express -> Yes Urgent, time sensitive consignments containing non-printed matter or non-document material
)

// All production types of domestic product group
var (
	ProductTypeDomesticOND ProductType = "OND"
	ProductTypeDomesticSMP ProductType = "SMP"
)

func (t ProductGroup) String() string {
	return string(t)
}

func (t ProductType) String() string {
	return string(t)
}
