package aramex

import (
	"context"
	"encoding/xml"

	"github.com/thaitanloi365/go-aramex-sdk/types"
)

type lastReservedShipmentNumberRangeRequest struct {
	XMLName      xml.Name           `xml:"http://ws.aramex.net/ShippingAPI/v1/ LastReservedShipmentNumberRangeRequest"`
	ClientInfo   *ClientInfo        `xml:"ClientInfo" json:"ClientInfo"`
	Transaction  *Transaction       `xml:"Transaction" json:"Transaction"`
	Entity       string             `xml:"Entity" json:"Entity"`
	ProductGroup types.ProductGroup `xml:"ProductGroup" json:"ProductGroup"`
}

type lastReservedShipmentNumberRangeResponse struct {
	XMLName       xml.Name             `xml:"http://ws.aramex.net/ShippingAPI/v1/ LastReservedShipmentNumberRangeResponse"`
	Transaction   *Transaction         `xml:"Transaction" json:"Transaction"`
	Notifications *arrayOfNotification `xml:"Notifications" json:"Notifications"`
	HasErrors     bool                 `xml:"HasErrors" json:"HasErrors"`
	FromWaybill   string               `xml:"FromWaybill" json:"FromWaybill"`
	ToWaybill     string               `xml:"ToWaybill" json:"ToWaybill"`
}

// LastReservedShipmentNumberRangeRequest request
type LastReservedShipmentNumberRangeRequest struct {
	ClientInfo   *ClientInfo        `xml:"ClientInfo" json:"ClientInfo"`
	Transaction  *Transaction       `xml:"Transaction" json:"Transaction"`
	Entity       string             `xml:"Entity" json:"Entity"`
	ProductGroup types.ProductGroup `xml:"ProductGroup" json:"ProductGroup"`
}

// LastReservedShipmentNumberRangeResponse response
type LastReservedShipmentNumberRangeResponse struct {
	Transaction   *Transaction  `xml:"Transaction" json:"Transaction"`
	Notifications Notifications `xml:"Notifications" json:"Notifications"`
	HasErrors     bool          `xml:"HasErrors" json:"HasErrors"`
	FromWaybill   string        `xml:"FromWaybill" json:"FromWaybill"`
	ToWaybill     string        `xml:"ToWaybill" json:"ToWaybill"`
}

// GetLastShipmentsNumbersRange last shipments
func (a *Aramex) GetLastShipmentsNumbersRange(ctx context.Context, request *LastReservedShipmentNumberRangeRequest) (*LastReservedShipmentNumberRangeResponse, error) {
	var resp = new(lastReservedShipmentNumberRangeResponse)
	var req = &lastReservedShipmentNumberRangeRequest{
		ClientInfo: a.config.ClientInfo,
	}
	if request != nil {
		req.Transaction = request.Transaction
		req.Entity = request.Entity
		req.ProductGroup = request.ProductGroup
	}

	if req.Entity == "" {
		req.Entity = a.config.ClientInfo.AccountEntity
	}

	if req.ProductGroup == "" {
		req.ProductGroup = a.config.DefaultSetting.ProductGroup
	}

	err := a.clients[shippingService].CallContext(ctx, a.buildURL("GetLastShipmentsNumbersRange"), req, resp)
	if err != nil {
		return nil, err
	}

	var response = &LastReservedShipmentNumberRangeResponse{
		Transaction:   resp.Transaction,
		HasErrors:     resp.HasErrors,
		FromWaybill:   resp.FromWaybill,
		ToWaybill:     resp.ToWaybill,
		Notifications: resp.Notifications.Notification,
	}
	return response, nil
}
