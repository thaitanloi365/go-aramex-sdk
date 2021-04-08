package aramex

import (
	"context"
	"encoding/xml"
)

type countryFetchingRequest struct {
	XMLName     xml.Name     `xml:"http://ws.aramex.net/ShippingAPI/v1/ CountriesFetchingRequest"`
	ClientInfo  *ClientInfo  `xml:"ClientInfo" json:"ClientInfo"`
	Transaction *Transaction `xml:"Transaction" json:"Transaction"`
}

type countriesFetchingResponse struct {
	XMLName       xml.Name             `xml:"http://ws.aramex.net/ShippingAPI/v1/ CountriesFetchingResponse"`
	Transaction   *Transaction         `xml:"Transaction,omitempty" json:"Transaction,omitempty"`
	Notifications *arrayOfNotification `xml:"Notifications,omitempty" json:"Notifications,omitempty"`
	HasErrors     bool                 `xml:"HasErrors,omitempty" json:"HasErrors,omitempty"`
	Countries     *arrayOfCountry      `xml:"Countries,omitempty" json:"Countries,omitempty"`
}

// CountriesFetchingRequest request
type CountriesFetchingRequest struct {
	Transaction *Transaction `xml:"Transaction" json:"Transaction"`
}

// CountriesFetchingResponse response
type CountriesFetchingResponse struct {
	Transaction   *Transaction  `xml:"Transaction" json:"Transaction"`
	Notifications Notifications `xml:"Notifications" json:"Notifications"`
	HasErrors     bool          `xml:"HasErrors" json:"HasErrors"`
	Countries     []*Country    `xml:"Countries" json:"Countries"`
}

// FetchCountries fetch countries
func (a *Aramex) FetchCountries(ctx context.Context, request *CountriesFetchingRequest) (*CountriesFetchingResponse, error) {
	var resp = new(countriesFetchingResponse)
	var req = &countryFetchingRequest{
		ClientInfo: a.config.ClientInfo,
	}

	if request != nil {
		req.Transaction = request.Transaction
	}
	err := a.clients[locationService].CallContext(ctx, a.buildURL("FetchCountries"), req, resp)
	if err != nil {
		return nil, err
	}

	var response = &CountriesFetchingResponse{
		Transaction:   resp.Transaction,
		HasErrors:     resp.HasErrors,
		Notifications: resp.Notifications.Notification,
		Countries:     resp.Countries.Country,
	}
	return response, nil
}
