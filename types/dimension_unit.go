package types

// DimensionUnit unit
type DimensionUnit string

// All units
var (
	DimensionUnitCM DimensionUnit = "CM"
	DimensionUnitM  DimensionUnit = "M"
)

func (t DimensionUnit) String() string {
	return string(t)
}
