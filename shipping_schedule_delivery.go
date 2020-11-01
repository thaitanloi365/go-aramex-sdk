package aramex

import (
	"context"
	"encoding/xml"

	"github.com/thaitanloi365/go-aramex-sdk/types"
)

type scheduledDeliveryRequest struct {
	XMLName           xml.Name           `xml:"http://ws.aramex.net/ShippingAPI/v1/ ScheduledDeliveryRequest"`
	ClientInfo        *ClientInfo        `xml:"ClientInfo" json:"ClientInfo"`
	Transaction       *Transaction       `xml:"Transaction" json:"Transaction"`
	Address           *Address           `xml:"Address" json:"Address"`
	ScheduledDelivery *scheduledDelivery `xml:"ScheduledDelivery" json:"ScheduledDelivery"`
	ProductGroup      types.ProductGroup `xml:"ProductGroup" json:"ProductGroup"`
	ShipmentNumber    string             `xml:"ShipmentNumber" json:"ShipmentNumber"`
	Entity            string             `xml:"Entity" json:"Entity"`
	ConsigneePhone    string             `xml:"ConsigneePhone" json:"ConsigneePhone"`
	ShipperNumber     string             `xml:"ShipperNumber" json:"ShipperNumber"`
	ShipperReference  string             `xml:"ShipperReference" json:"ShipperReference"`
	Reference1        string             `xml:"Reference1" json:"Reference1"`
	Reference2        string             `xml:"Reference2" json:"Reference2"`
	Reference3        string             `xml:"Reference3" json:"Reference3"`
}

type scheduledDeliveryResponse struct {
	XMLName       xml.Name             `xml:"http://ws.aramex.net/ShippingAPI/v1/ ScheduledDeliveryResponse"`
	Transaction   *Transaction         `xml:"Transaction" json:"Transaction"`
	Notifications *arrayOfNotification `xml:"Notifications" json:"Notifications"`
	HasErrors     bool                 `xml:"HasErrors" json:"HasErrors"`
	Id            int32                `xml:"Id" json:"Id"`
}

// ScheduledDeliveryRequest request
type ScheduledDeliveryRequest struct {
	XMLName           xml.Name           `xml:"http://ws.aramex.net/ShippingAPI/v1/ ScheduledDeliveryRequest"`
	ClientInfo        *ClientInfo        `xml:"ClientInfo" json:"ClientInfo"`
	Transaction       *Transaction       `xml:"Transaction" json:"Transaction"`
	Address           *Address           `xml:"Address" json:"Address"`
	ScheduledDelivery *ScheduledDelivery `xml:"ScheduledDelivery" json:"ScheduledDelivery"`
	ProductGroup      types.ProductGroup `xml:"ProductGroup" json:"ProductGroup"`
	ShipmentNumber    string             `xml:"ShipmentNumber" json:"ShipmentNumber"`
	Entity            string             `xml:"Entity" json:"Entity"`
	ConsigneePhone    string             `xml:"ConsigneePhone" json:"ConsigneePhone"`
	ShipperNumber     string             `xml:"ShipperNumber" json:"ShipperNumber"`
	ShipperReference  string             `xml:"ShipperReference" json:"ShipperReference"`
	Reference1        string             `xml:"Reference1" json:"Reference1"`
	Reference2        string             `xml:"Reference2" json:"Reference2"`
	Reference3        string             `xml:"Reference3" json:"Reference3"`
}

// ScheduledDeliveryResponse response
type ScheduledDeliveryResponse struct {
	Transaction   *Transaction    `xml:"Transaction" json:"Transaction"`
	Notifications []*Notification `xml:"Notifications" json:"Notifications"`
	HasErrors     bool            `xml:"HasErrors" json:"HasErrors"`
	Id            int32           `xml:"Id" json:"Id"`
}

// ScheduleDelivery schedule delivery
func (a *Aramex) ScheduleDelivery(ctx context.Context, request *ScheduledDeliveryRequest) (*ScheduledDeliveryResponse, error) {
	var resp = new(scheduledDeliveryResponse)
	var req = &scheduledDeliveryRequest{
		ClientInfo:       a.config.ClientInfo,
		Address:          request.Address,
		ConsigneePhone:   request.ConsigneePhone,
		Entity:           request.Entity,
		ProductGroup:     request.ProductGroup,
		Reference1:       request.Reference1,
		Reference2:       request.Reference2,
		Reference3:       request.Reference3,
		ShipmentNumber:   request.ShipmentNumber,
		ShipperNumber:    request.ShipperNumber,
		ShipperReference: request.ShipperReference,
		Transaction:      request.Transaction,
	}

	if request.ScheduledDelivery != nil {
		req.ScheduledDelivery = &scheduledDelivery{
			PreferredDeliveryDate: request.ScheduledDelivery.PreferredDeliveryDate,
			PreferredDeliveryTime: request.ScheduledDelivery.PreferredDeliveryTime,
		}
	}

	if req.Entity == "" {
		req.Entity = a.config.ClientInfo.AccountEntity
	}

	if req.ProductGroup == "" {
		req.ProductGroup = a.config.DefaultSetting.ProductGroup
	}

	err := a.clients[shippingService].CallContext(ctx, a.buildURL("ScheduleDelivery"), req, resp)
	if err != nil {
		return nil, err
	}

	var response = &ScheduledDeliveryResponse{
		Transaction:   resp.Transaction,
		HasErrors:     resp.HasErrors,
		Notifications: resp.Notifications.Notification,
		Id:            resp.Id,
	}
	return response, nil
}
