package aramex

import (
	"context"
	"encoding/xml"
)

type pickupCancelationRequest struct {
	XMLName     xml.Name     `xml:"http://ws.aramex.net/ShippingAPI/v1/ PickupCancelationRequest"`
	ClientInfo  *ClientInfo  `xml:"ClientInfo" json:"ClientInfo"`
	Transaction *Transaction `xml:"Transaction" json:"Transaction"`
	PickupGUID  string       `xml:"PickupGUID" json:"PickupGUID"`
	Comments    string       `xml:"Comments" json:"Comments"`
}

type pickupCancelationResponse struct {
	XMLName       xml.Name             `xml:"http://ws.aramex.net/ShippingAPI/v1/ PickupCancelationResponse"`
	Transaction   *Transaction         `xml:"Transaction" json:"Transaction"`
	Notifications *arrayOfNotification `xml:"Notifications" json:"Notifications"`
	HasErrors     bool                 `xml:"HasErrors" json:"HasErrors"`
	Message       string               `xml:"Message" json:"Message"`
}

// PickupCancelationRequest request
type PickupCancelationRequest struct {
	Transaction *Transaction `xml:"Transaction" json:"Transaction"`
	PickupGUID  string       `xml:"PickupGUID" json:"PickupGUID"`
	Comments    string       `xml:"Comments" json:"Comments"`
}

// PickupCancelationResponse response
type PickupCancelationResponse struct {
	Transaction   *Transaction  `xml:"Transaction" json:"Transaction"`
	Notifications Notifications `xml:"Notifications" json:"Notifications"`
	HasErrors     bool          `xml:"HasErrors" json:"HasErrors"`
	Message       string        `xml:"Message" json:"Message"`
}

// CancelPickup cancel pickup
func (a *Aramex) CancelPickup(ctx context.Context, request *PickupCancelationRequest) (*PickupCancelationResponse, error) {
	var resp = new(pickupCancelationResponse)
	var req = &pickupCancelationRequest{
		ClientInfo:  a.config.ClientInfo,
		Comments:    request.Comments,
		PickupGUID:  request.PickupGUID,
		Transaction: request.Transaction,
	}
	err := a.clients[shippingService].CallContext(ctx, a.buildURL("CancelPickup"), req, resp)
	if err != nil {
		return nil, err
	}

	var response = &PickupCancelationResponse{
		Transaction:   resp.Transaction,
		HasErrors:     resp.HasErrors,
		Message:       resp.Message,
		Notifications: resp.Notifications.Notification,
	}
	return response, nil
}
