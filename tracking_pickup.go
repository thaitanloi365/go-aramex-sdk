package aramex

import (
	"context"
	"encoding/xml"
	"time"
)

type pickupTrackingRequest struct {
	XMLName     xml.Name     `xml:"http://ws.aramex.net/ShippingAPI/v1/ PickupTrackingRequest"`
	ClientInfo  *ClientInfo  `xml:"ClientInfo" json:"ClientInfo"`
	Transaction *Transaction `xml:"Transaction" json:"Transaction"`
	Reference   string       `xml:"Reference" json:"Reference"`
}

type pickupTrackingResponse struct {
	XMLName               xml.Name             `xml:"http://ws.aramex.net/ShippingAPI/v1/ PickupTrackingResponse"`
	Transaction           *Transaction         `xml:"Transaction" json:"Transaction"`
	Notifications         *arrayOfNotification `xml:"Notifications" json:"Notifications"`
	HasErrors             bool                 `xml:"HasErrors" json:"HasErrors"`
	Entity                string               `xml:"Entity" json:"Entity"`
	Reference             string               `xml:"Reference" json:"Reference"`
	CollectionDate        string               `xml:"CollectionDate" json:"CollectionDate"`
	PickupDate            string               `xml:"PickupDate" json:"PickupDate"`
	LastStatus            string               `xml:"LastStatus" json:"LastStatus"`
	LastStatusDescription string               `xml:"LastStatusDescription" json:"LastStatusDescription"`
	CollectedWaybills     *arrayOfstring       `xml:"CollectedWaybills" json:"CollectedWaybills"`
}

// PickupTrackingRequest request
type PickupTrackingRequest struct {
	PickupID string
}

// PickupTrackingResponse response
type PickupTrackingResponse struct {
	Transaction           *Transaction
	Notifications         []*Notification
	HasErrors             bool
	Entity                string
	Reference             string
	CollectionDate        time.Time
	PickupDate            time.Time
	LastStatus            string
	LastStatusDescription string
	CollectedWaybills     []string
}

// TrackPickup track shipments
func (a *Aramex) TrackPickup(ctx context.Context, request *PickupTrackingRequest) (*PickupTrackingResponse, error) {
	var resp = new(pickupTrackingResponse)
	var req = &pickupTrackingRequest{
		ClientInfo: a.config.ClientInfo,
		Reference:  request.PickupID,
	}
	var err = a.clients[trackingService].CallContext(ctx, a.buildURL("TrackPickup"), req, resp)
	if err != nil {
		return nil, err
	}

	var response = &PickupTrackingResponse{
		Transaction:           resp.Transaction,
		HasErrors:             resp.HasErrors,
		Notifications:         resp.Notifications.Notification,
		Entity:                resp.Entity,
		CollectedWaybills:     resp.CollectedWaybills.Astring,
		LastStatus:            resp.LastStatus,
		Reference:             resp.Reference,
		LastStatusDescription: resp.LastStatusDescription,
	}

	if t, err := time.Parse("2006-01-02T15:04:05", resp.CollectionDate); err == nil {
		response.CollectionDate = t
	}

	if t, err := time.Parse("2006-01-02T15:04:05", resp.PickupDate); err == nil {
		response.PickupDate = t
	}

	return response, nil
}