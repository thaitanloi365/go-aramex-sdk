package aramex

import (
	"context"
	"encoding/xml"

	"github.com/thaitanloi365/go-aramex-sdk/types"
)

type citiesFetchingRequest struct {
	XMLName        xml.Name          `xml:"http://ws.aramex.net/ShippingAPI/v1/ CitiesFetchingRequest"`
	ClientInfo     *ClientInfo       `xml:"ClientInfo" json:"ClientInfo"`
	Transaction    *Transaction      `xml:"Transaction" json:"Transaction"`
	CountryCode    types.CountryCode `xml:"CountryCode" json:"CountryCode"`
	State          string            `xml:"State" json:"State"`
	NameStartsWith string            `xml:"NameStartsWith" json:"NameStartsWith"`
}

type citiesFetchingResponse struct {
	XMLName       xml.Name             `xml:"http://ws.aramex.net/ShippingAPI/v1/ CitiesFetchingResponse"`
	Transaction   *Transaction         `xml:"Transaction" json:"Transaction"`
	Notifications *arrayOfNotification `xml:"Notifications" json:"Notifications"`
	HasErrors     bool                 `xml:"HasErrors" json:"HasErrors"`
	Cities        *arrayOfstring       `xml:"Cities" json:"Cities"`
}

// CitiesFetchingRequest request
type CitiesFetchingRequest struct {
	Transaction    *Transaction      `xml:"Transaction" json:"Transaction"`
	CountryCode    types.CountryCode `xml:"CountryCode" json:"CountryCode"`
	State          string            `xml:"State" json:"State"`
	NameStartsWith string            `xml:"NameStartsWith" json:"NameStartsWith"`
}

// CitiesFetchingResponse response
type CitiesFetchingResponse struct {
	Transaction   *Transaction  `xml:"Transaction" json:"Transaction"`
	Notifications Notifications `xml:"Notifications" json:"Notifications"`
	HasErrors     bool          `xml:"HasErrors" json:"HasErrors"`
	Cities        []string      `xml:"Cities" json:"Cities"`
}

// FetchCities fetch cities
func (a *Aramex) FetchCities(ctx context.Context, request *CitiesFetchingRequest) (*CitiesFetchingResponse, error) {
	var resp = new(citiesFetchingResponse)
	var req = &citiesFetchingRequest{
		ClientInfo:     a.config.ClientInfo,
		Transaction:    request.Transaction,
		CountryCode:    request.CountryCode,
		NameStartsWith: request.NameStartsWith,
		State:          request.State,
	}
	err := a.clients[locationService].CallContext(ctx, a.buildURL("FetchCities"), req, resp)
	if err != nil {
		return nil, err
	}

	var response = &CitiesFetchingResponse{
		Transaction:   resp.Transaction,
		HasErrors:     resp.HasErrors,
		Notifications: resp.Notifications.Notification,
		Cities:        resp.Cities.Astring,
	}
	return response, nil
}
