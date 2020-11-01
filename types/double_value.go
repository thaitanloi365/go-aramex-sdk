package types

// Double double
type Double float64

// Float64 to float64
func (v Double) Float64() float64 {
	return float64(v)
}
