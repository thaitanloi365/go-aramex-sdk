package aramex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrackShipments(t *testing.T) {
	var aramex = New(&Config{
		IsLive:     false,
		ClientInfo: DefaultClientInfo,
	})
	result, err := aramex.TrackShipments(context.Background(), &ShipmentTrackingRequest{
		Shipments: []string{
			"42489016231",
		},
	})
	assert.NoError(t, err)

	aramex.printJSON(&result)
}

func TestTrackPickup(t *testing.T) {
	var aramex = New(&Config{
		IsLive:     false,
		ClientInfo: DefaultClientInfo,
	})
	result, err := aramex.TrackPickup(context.Background(), &PickupTrackingRequest{
		PickupID: "J073C31",
	})
	assert.NoError(t, err)

	aramex.printJSON(&result)
}
