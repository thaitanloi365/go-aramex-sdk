package aramex

import (
	"context"
	"encoding/xml"

	"github.com/thaitanloi365/go-aramex-sdk/types"
)

type labelPrintingRequest struct {
	XMLName        xml.Name           `xml:"http://ws.aramex.net/ShippingAPI/v1/ LabelPrintingRequest"`
	ClientInfo     *ClientInfo        `xml:"ClientInfo" json:"ClientInfo"`
	Transaction    *Transaction       `xml:"Transaction" json:"Transaction"`
	ShipmentNumber string             `xml:"ShipmentNumber" json:"ShipmentNumber"`
	ProductGroup   types.ProductGroup `xml:"ProductGroup" json:"ProductGroup"`
	OriginEntity   string             `xml:"OriginEntity" json:"OriginEntity"`
	LabelInfo      *LabelInfo         `xml:"LabelInfo" json:"LabelInfo"`
}

type labelPrintingResponse struct {
	XMLName        xml.Name             `xml:"http://ws.aramex.net/ShippingAPI/v1/ LabelPrintingResponse"`
	Transaction    *Transaction         `xml:"Transaction" json:"Transaction"`
	Notifications  *arrayOfNotification `xml:"Notifications" json:"Notifications"`
	HasErrors      bool                 `xml:"HasErrors" json:"HasErrors"`
	ShipmentNumber string               `xml:"ShipmentNumber" json:"ShipmentNumber"`
	ShipmentLabel  *ShipmentLabel       `xml:"ShipmentLabel" json:"ShipmentLabel"`
}

// LabelPrintingRequest request
type LabelPrintingRequest struct {
	ShipmentNumber string             `xml:"ShipmentNumber" json:"ShipmentNumber"`
	ProductGroup   types.ProductGroup `xml:"ProductGroup" json:"ProductGroup"`
	OriginEntity   string             `xml:"OriginEntity" json:"OriginEntity"`
	LabelInfo      *LabelInfo         `xml:"LabelInfo" json:"LabelInfo"`
}

// LabelPrintingResponse response
type LabelPrintingResponse struct {
	Transaction    *Transaction   `xml:"Transaction" json:"Transaction"`
	Notifications  Notifications  `xml:"Notifications" json:"Notifications"`
	HasErrors      bool           `xml:"HasErrors" json:"HasErrors"`
	ShipmentNumber string         `xml:"ShipmentNumber" json:"ShipmentNumber"`
	ShipmentLabel  *ShipmentLabel `xml:"ShipmentLabel" json:"ShipmentLabel"`
}

// PrintLabel create shipments
func (a *Aramex) PrintLabel(ctx context.Context, request *LabelPrintingRequest) (*LabelPrintingResponse, error) {
	var resp = new(labelPrintingResponse)
	var req = &labelPrintingRequest{
		ClientInfo:     a.config.ClientInfo,
		LabelInfo:      request.LabelInfo,
		OriginEntity:   request.OriginEntity,
		ProductGroup:   request.ProductGroup,
		ShipmentNumber: request.ShipmentNumber,
	}

	if req.LabelInfo == nil && a.config.DefaultSetting != nil {
		req.LabelInfo = a.config.DefaultSetting.LabelInfo
	}

	var err = a.clients[shippingService].CallContext(ctx, a.buildURL("PrintLabel"), req, resp)
	if err != nil {
		return nil, err
	}

	var response = &LabelPrintingResponse{
		Transaction:    resp.Transaction,
		HasErrors:      resp.HasErrors,
		Notifications:  resp.Notifications.Notification,
		ShipmentLabel:  resp.ShipmentLabel,
		ShipmentNumber: resp.ShipmentNumber,
	}

	return response, nil
}
