package aramex

import (
	"context"
	"encoding/xml"

	"github.com/thaitanloi365/go-aramex-sdk/types"
)

type statesFetchingRequest struct {
	XMLName     xml.Name          `xml:"http://ws.aramex.net/ShippingAPI/v1/ StatesFetchingRequest"`
	ClientInfo  *ClientInfo       `xml:"ClientInfo" json:"ClientInfo"`
	Transaction *Transaction      `xml:"Transaction" json:"Transaction"`
	CountryCode types.CountryCode `xml:"CountryCode" json:"CountryCode"`
}

type statesFetchingResponse struct {
	XMLName       xml.Name             `xml:"http://ws.aramex.net/ShippingAPI/v1/ StatesFetchingResponse"`
	Transaction   *Transaction         `xml:"Transaction" json:"Transaction"`
	Notifications *arrayOfNotification `xml:"Notifications" json:"Notifications"`
	HasErrors     bool                 `xml:"HasErrors" json:"HasErrors"`
	States        *arrayOfState        `xml:"States" json:"States"`
}

// StatesFetchingRequest request
type StatesFetchingRequest struct {
	Transaction *Transaction      `xml:"Transaction" json:"Transaction"`
	CountryCode types.CountryCode `xml:"CountryCode" json:"CountryCode"`
}

// StatesFetchingResponse response
type StatesFetchingResponse struct {
	Transaction   *Transaction    `xml:"Transaction" json:"Transaction"`
	Notifications []*Notification `xml:"Notifications" json:"Notifications"`
	HasErrors     bool            `xml:"HasErrors" json:"HasErrors"`
	States        []*State        `xml:"States" json:"States"`
}

// FetchStates fetch states
func (a *Aramex) FetchStates(ctx context.Context, request *StatesFetchingRequest) (*StatesFetchingResponse, error) {
	var resp = new(statesFetchingResponse)
	var req = &statesFetchingRequest{
		ClientInfo:  a.config.ClientInfo,
		Transaction: request.Transaction,
		CountryCode: request.CountryCode,
	}
	err := a.clients[locationService].CallContext(ctx, a.buildURL("FetchStates"), req, resp)
	if err != nil {
		return nil, err
	}

	var response = &StatesFetchingResponse{
		Transaction:   resp.Transaction,
		HasErrors:     resp.HasErrors,
		Notifications: resp.Notifications.Notification,
		States:        resp.States.State,
	}
	return response, nil
}
