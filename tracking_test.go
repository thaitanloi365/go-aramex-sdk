package aramex

import (
	"context"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestTrackShipments(t *testing.T) {
	var aramex = New(&Config{
		IsLive:     false,
		ClientInfo: DefaultClientInfo,
	})
	result, err := aramex.TrackShipments(context.Background(), &ShipmentTrackingRequest{
		Shipments: []string{
			// "41118182136",
		},
	})
	assert.NoError(t, err)

	if len(result.Notifications) > 0 {
		var e = errors.Wrap(result.Notifications[0], "sdfsdf")
		fmt.Println("err", e)
	}
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
