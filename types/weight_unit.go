package types

// WeightUnit unit
type WeightUnit string

// All units
var (
	WeightUnitKG WeightUnit = "KG"
	WeightUnitLB WeightUnit = "LB"
)

func (t WeightUnit) String() string {
	return string(t)
}
