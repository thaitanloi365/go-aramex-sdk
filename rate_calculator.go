package aramex

import (
	"context"
	"encoding/xml"

	"github.com/thaitanloi365/go-aramex-sdk/types"
)

type rateCalculatorRequest struct {
	XMLName               xml.Name           `xml:"http://ws.aramex.net/ShippingAPI/v1/ RateCalculatorRequest"`
	ClientInfo            *ClientInfo        `xml:"ClientInfo" json:"ClientInfo"`
	Transaction           *Transaction       `xml:"Transaction" json:"Transaction"`
	OriginAddress         *Address           `xml:"OriginAddress" json:"OriginAddress"`
	DestinationAddress    *Address           `xml:"DestinationAddress" json:"DestinationAddress"`
	ShipmentDetails       *shipmentDetails   `xml:"ShipmentDetails" json:"ShipmentDetails"`
	PreferredCurrencyCode types.CurrencyCode `xml:"PreferredCurrencyCode" json:"PreferredCurrencyCode"`
}

type rateCalculatorResponse struct {
	XMLName       xml.Name             `xml:"http://ws.aramex.net/ShippingAPI/v1/ RateCalculatorResponse"`
	Transaction   *Transaction         `xml:"Transaction" json:"Transaction"`
	Notifications *arrayOfNotification `xml:"Notifications" json:"Notifications"`
	HasErrors     bool                 `xml:"HasErrors" json:"HasErrors"`
	TotalAmount   *Money               `xml:"TotalAmount" json:"TotalAmount"`
	RateDetails   *RateDetails         `xml:"RateDetails" json:"RateDetails"`
}

// RateDetails details
type RateDetails struct {
	Amount               float64 `xml:"Amount" json:"Amount"`
	OtherAmount1         float64 `xml:"OtherAmount1" json:"OtherAmount1"`
	OtherAmount2         float64 `xml:"OtherAmount2" json:"OtherAmount2"`
	OtherAmount3         float64 `xml:"OtherAmount3" json:"OtherAmount3"`
	OtherAmount4         float64 `xml:"OtherAmount4" json:"OtherAmount4"`
	OtherAmount5         float64 `xml:"OtherAmount5" json:"OtherAmount5"`
	TotalAmountBeforeTax float64 `xml:"TotalAmountBeforeTax" json:"TotalAmountBeforeTax"`
	TaxAmount            float64 `xml:"TaxAmount" json:"TaxAmount"`
}

// RateCalculatorRequest request
type RateCalculatorRequest struct {
	OriginAddress         *Address           `xml:"OriginAddress" json:"OriginAddress"`
	DestinationAddress    *Address           `xml:"DestinationAddress" json:"DestinationAddress"`
	ShipmentDetails       *ShipmentDetails   `xml:"ShipmentDetails" json:"ShipmentDetails"`
	PreferredCurrencyCode types.CurrencyCode `xml:"PreferredCurrencyCode" json:"PreferredCurrencyCode"`
}

// RateCalculatorResponse response
type RateCalculatorResponse struct {
	Transaction   *Transaction  `xml:"Transaction" json:"Transaction"`
	Notifications Notifications `xml:"Notifications" json:"Notifications"`
	HasErrors     bool          `xml:"HasErrors" json:"HasErrors"`
	TotalAmount   *Money        `xml:"TotalAmount" json:"TotalAmount"`
	RateDetails   *RateDetails  `xml:"RateDetails" json:"RateDetails"`
}

// CalculateRate calculate rate
func (a *Aramex) CalculateRate(ctx context.Context, request RateCalculatorRequest) (*RateCalculatorResponse, error) {
	var resp = new(rateCalculatorResponse)
	var req = &rateCalculatorRequest{
		ClientInfo:            a.config.ClientInfo,
		OriginAddress:         request.OriginAddress,
		DestinationAddress:    request.DestinationAddress,
		PreferredCurrencyCode: request.PreferredCurrencyCode,
		ShipmentDetails:       a.toShipmentDetailsRequest(request.ShipmentDetails),
	}
	var err = a.clients[rateService].CallContext(ctx, a.buildURL("CalculateRate"), req, resp)
	if err != nil {
		return nil, err
	}

	var response = &RateCalculatorResponse{
		Transaction:   resp.Transaction,
		HasErrors:     resp.HasErrors,
		Notifications: resp.Notifications.Notification,
		RateDetails:   resp.RateDetails,
		TotalAmount:   resp.TotalAmount,
	}

	return response, nil
}
