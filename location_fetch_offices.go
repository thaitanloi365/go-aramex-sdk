package aramex

import (
	"context"
	"encoding/xml"

	"github.com/thaitanloi365/go-aramex-sdk/types"
)

type officesFetchingRequest struct {
	XMLName     xml.Name          `xml:"http://ws.aramex.net/ShippingAPI/v1/ OfficesFetchingRequest"`
	ClientInfo  *ClientInfo       `xml:"ClientInfo" json:"ClientInfo"`
	Transaction *Transaction      `xml:"Transaction" json:"Transaction"`
	CountryCode types.CountryCode `xml:"CountryCode" json:"CountryCode"`
}

type officesFetchingResponse struct {
	XMLName       xml.Name             `xml:"http://ws.aramex.net/ShippingAPI/v1/ OfficesFetchingResponse"`
	Transaction   *Transaction         `xml:"Transaction" json:"Transaction"`
	Notifications *arrayOfNotification `xml:"Notifications" json:"Notifications"`
	HasErrors     bool                 `xml:"HasErrors" json:"HasErrors"`
	Offices       *arrayOfOffice       `xml:"Offices" json:"Offices"`
}

// OfficesFetchingRequest request
type OfficesFetchingRequest struct {
	Transaction *Transaction      `xml:"Transaction" json:"Transaction"`
	CountryCode types.CountryCode `xml:"CountryCode" json:"CountryCode"`
}

// OfficesFetchingResponse response
type OfficesFetchingResponse struct {
	Transaction   *Transaction  `xml:"Transaction" json:"Transaction"`
	Notifications Notifications `xml:"Notifications" json:"Notifications"`
	HasErrors     bool          `xml:"HasErrors" json:"HasErrors"`
	Offices       []*Office     `xml:"Offices" json:"Offices"`
}

// FetchOffices fetch offices
func (a *Aramex) FetchOffices(ctx context.Context, request *OfficesFetchingRequest) (*OfficesFetchingResponse, error) {
	var resp = new(officesFetchingResponse)
	var req = &officesFetchingRequest{
		ClientInfo:  a.config.ClientInfo,
		Transaction: request.Transaction,
		CountryCode: request.CountryCode,
	}
	err := a.clients[locationService].CallContext(ctx, a.buildURL("FetchOffices"), req, resp)
	if err != nil {
		return nil, err
	}

	var response = &OfficesFetchingResponse{
		Transaction:   resp.Transaction,
		HasErrors:     resp.HasErrors,
		Notifications: resp.Notifications.Notification,
		Offices:       resp.Offices.Office,
	}
	return response, nil
}
