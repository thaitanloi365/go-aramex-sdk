package types

// VehicleType status
type VehicleType string

// All status
var (
	VehicleTypeSmall  VehicleType = "small"
	VehicleTypeMedium VehicleType = "medium"
	VehicleTypeLarge  VehicleType = "large"
)

func (t VehicleType) String() string {
	return string(t)
}
