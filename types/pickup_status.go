package types

// PickupStatus status
type PickupStatus string

// All status
var (
	PickupStatusPending PickupStatus = "Pending"
	PickupStatusReady   PickupStatus = "Ready"
)

func (t PickupStatus) String() string {
	return string(t)
}
