package aramex

import (
	"context"
	"encoding/xml"
)

type addressValidationRequest struct {
	XMLName     xml.Name     `xml:"http://ws.aramex.net/ShippingAPI/v1/ AddressValidationRequest"`
	ClientInfo  *ClientInfo  `xml:"ClientInfo" json:"ClientInfo"`
	Transaction *Transaction `xml:"Transaction" json:"Transaction"`
	Address     *Address     `xml:"Address" json:"Address"`
}

type addressValidationResponse struct {
	XMLName            xml.Name             `xml:"http://ws.aramex.net/ShippingAPI/v1/ AddressValidationResponse"`
	Transaction        *Transaction         `xml:"Transaction" json:"Transaction"`
	Notifications      *arrayOfNotification `xml:"Notifications" json:"Notifications"`
	HasErrors          bool                 `xml:"HasErrors" json:"HasErrors"`
	SuggestedAddresses *arrayOfAddress      `xml:"SuggestedAddresses" json:"SuggestedAddresses"`
}

// AddressValidationRequest request
type AddressValidationRequest struct {
	Transaction *Transaction `xml:"Transaction" json:"Transaction"`
	Address     *Address     `xml:"Address" json:"Address"`
}

// AddressValidationResponse response
type AddressValidationResponse struct {
	Transaction        *Transaction  `xml:"Transaction" json:"Transaction"`
	Notifications      Notifications `xml:"Notifications" json:"Notifications"`
	HasErrors          bool          `xml:"HasErrors" json:"HasErrors"`
	SuggestedAddresses []*Address    `xml:"SuggestedAddresses" json:"SuggestedAddresses"`
}

// ValidateAddress validate address
func (a *Aramex) ValidateAddress(ctx context.Context, request *AddressValidationRequest) (*AddressValidationResponse, error) {
	var resp = new(addressValidationResponse)
	var req = &addressValidationRequest{
		ClientInfo:  a.config.ClientInfo,
		Transaction: request.Transaction,
		Address:     request.Address,
	}
	err := a.clients[locationService].CallContext(ctx, a.buildURL("ValidateAddress"), req, resp)
	if err != nil {
		return nil, err
	}

	var response = &AddressValidationResponse{
		Transaction:        resp.Transaction,
		HasErrors:          resp.HasErrors,
		Notifications:      resp.Notifications.Notification,
		SuggestedAddresses: resp.SuggestedAddresses.Address,
	}
	return response, nil
}
