package types

// VolumeUnit unit
type VolumeUnit string

// All units
var (
	VolumeUnitCm3   VolumeUnit = "Cm3"
	VolumeUnitInch3 VolumeUnit = "Inch3"
)

func (t VolumeUnit) String() string {
	return string(t)
}
