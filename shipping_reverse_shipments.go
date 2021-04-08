package aramex

import (
	"context"
	"encoding/xml"

	"github.com/thaitanloi365/go-aramex-sdk/types"
)

type reserveRangeRequest struct {
	XMLName      xml.Name           `xml:"http://ws.aramex.net/ShippingAPI/v1/ ReserveRangeRequest"`
	ClientInfo   *ClientInfo        `xml:"ClientInfo" json:"ClientInfo"`
	Transaction  *Transaction       `xml:"Transaction" json:"Transaction"`
	Entity       string             `xml:"Entity" json:"Entity"`
	ProductGroup types.ProductGroup `xml:"ProductGroup" json:"ProductGroup"`
	Count        int32              `xml:"Count" json:"Count"`
}

type reserveRangeResponse struct {
	XMLName       xml.Name             `xml:"http://ws.aramex.net/ShippingAPI/v1/ ReserveRangeResponse"`
	Transaction   *Transaction         `xml:"Transaction" json:"Transaction"`
	Notifications *arrayOfNotification `xml:"Notifications" json:"Notifications"`
	HasErrors     bool                 `xml:"HasErrors" json:"HasErrors"`
	FromWaybill   string               `xml:"FromWaybill" json:"FromWaybill"`
	ToWaybill     string               `xml:"ToWaybill" json:"ToWaybill"`
}

// ReserveRangeRequest request
type ReserveRangeRequest struct {
	Transaction  *Transaction       `xml:"Transaction" json:"Transaction"`
	Entity       string             `xml:"Entity" json:"Entity"`
	ProductGroup types.ProductGroup `xml:"ProductGroup" json:"ProductGroup"`
	Count        int32              `xml:"Count" json:"Count"`
}

// ReserveRangeResponse response
type ReserveRangeResponse struct {
	Transaction   *Transaction  `xml:"Transaction" json:"Transaction"`
	Notifications Notifications `xml:"Notifications" json:"Notifications"`
	HasErrors     bool          `xml:"HasErrors" json:"HasErrors"`
	FromWaybill   string        `xml:"FromWaybill" json:"FromWaybill"`
	ToWaybill     string        `xml:"ToWaybill" json:"ToWaybill"`
}

// ReserveShipmentNumberRange reverse shipments
func (a *Aramex) ReserveShipmentNumberRange(ctx context.Context, request *ReserveRangeRequest) (*ReserveRangeResponse, error) {
	var resp = new(reserveRangeResponse)
	var req = &reserveRangeRequest{
		ClientInfo: a.config.ClientInfo,
	}
	if request != nil {
		req.Count = request.Count
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

	if req.Count == 0 {
		req.Count = 10000 // Max is 10000
	}

	err := a.clients[shippingService].CallContext(ctx, a.buildURL("ReserveShipmentNumberRange"), req, resp)
	if err != nil {
		return nil, err
	}

	var response = &ReserveRangeResponse{
		Transaction:   resp.Transaction,
		HasErrors:     resp.HasErrors,
		FromWaybill:   resp.FromWaybill,
		ToWaybill:     resp.ToWaybill,
		Notifications: resp.Notifications.Notification,
	}
	return response, nil
}
