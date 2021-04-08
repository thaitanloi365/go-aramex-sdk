package aramex

import (
	"context"
	"encoding/xml"
)

type arrayOfProcessedShipmentHold struct {
	ProcessedShipmentHold []*ProcessedShipmentHold `xml:"ProcessedShipmentHold" json:"ProcessedShipmentHold"`
}

type arrayOfShipmentHoldDetails struct {
	ShipmentHoldDetails []*ShipmentHoldDetails `xml:"ShipmentHoldDetails" json:"ShipmentHoldDetails"`
}

type holdCreationRequest struct {
	XMLName       xml.Name                    `xml:"http://ws.aramex.net/ShippingAPI/v1/ HoldCreationRequest"`
	ClientInfo    *ClientInfo                 `xml:"ClientInfo" json:"ClientInfo"`
	Transaction   *Transaction                `xml:"Transaction" json:"Transaction"`
	ShipmentHolds *arrayOfShipmentHoldDetails `xml:"ShipmentHolds" json:"ShipmentHolds"`
}

type holdCreationResponse struct {
	XMLName                xml.Name                      `xml:"http://ws.aramex.net/ShippingAPI/v1/ HoldCreationResponse"`
	Transaction            *Transaction                  `xml:"Transaction" json:"Transaction"`
	Notifications          *arrayOfNotification          `xml:"Notifications" json:"Notifications"`
	HasErrors              bool                          `xml:"HasErrors" json:"HasErrors"`
	ProcessedShipmentHolds *arrayOfProcessedShipmentHold `xml:"ProcessedShipmentHolds" json:"ProcessedShipmentHolds"`
}

type processedShipmentHold struct {
	ID            string               `xml:"ID" json:"ID"`
	HasErrors     bool                 `xml:"HasErrors" json:"HasErrors"`
	Notifications *arrayOfNotification `xml:"Notifications" json:"Notifications"`
}

// ShipmentHoldDetails details
type ShipmentHoldDetails struct {
	ShipmentNumber string `xml:"ShipmentNumber" json:"ShipmentNumber"`
	Comment        string `xml:"Comment" json:"Comment"`
}

// ProcessedShipmentHold hold shipment
type ProcessedShipmentHold struct {
	ID            string        `xml:"ID" json:"ID"`
	HasErrors     bool          `xml:"HasErrors" json:"HasErrors"`
	Notifications Notifications `xml:"Notifications" json:"Notifications"`
}

// HoldCreationRequest request
type HoldCreationRequest struct {
	Transaction   *Transaction           `xml:"Transaction" json:"Transaction"`
	ShipmentHolds []*ShipmentHoldDetails `xml:"ShipmentHolds" json:"ShipmentHolds"`
}

// HoldCreationResponse response
type HoldCreationResponse struct {
	Transaction            *Transaction             `xml:"Transaction" json:"Transaction"`
	Notifications          Notifications            `xml:"Notifications" json:"Notifications"`
	HasErrors              bool                     `xml:"HasErrors" json:"HasErrors"`
	ProcessedShipmentHolds []*ProcessedShipmentHold `xml:"ProcessedShipmentHolds" json:"ProcessedShipmentHolds"`
}

// HoldShipments hold shipments
func (a *Aramex) HoldShipments(ctx context.Context, request *HoldCreationRequest) (*HoldCreationResponse, error) {
	var resp = new(holdCreationResponse)
	var req = &holdCreationRequest{
		ClientInfo: a.config.ClientInfo,
	}

	if len(request.ShipmentHolds) > 0 {
		req.ShipmentHolds = &arrayOfShipmentHoldDetails{
			ShipmentHoldDetails: request.ShipmentHolds,
		}
	}
	err := a.clients[shippingService].CallContext(ctx, a.buildURL("GetLastShipmentsNumbersRange"), req, resp)
	if err != nil {
		return nil, err
	}

	var response = &HoldCreationResponse{
		Transaction:            resp.Transaction,
		HasErrors:              resp.HasErrors,
		Notifications:          resp.Notifications.Notification,
		ProcessedShipmentHolds: resp.ProcessedShipmentHolds.ProcessedShipmentHold,
	}
	return response, nil
}
