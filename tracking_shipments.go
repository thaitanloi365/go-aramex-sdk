package aramex

import (
	"context"
	"encoding/xml"
	"time"
)

type arrayOfTrackingResult struct {
	TrackingResult []*trackingResult `xml:"TrackingResult" json:"TrackingResult"`
}

type arrayOfKeyValueOfstringArrayOfTrackingResultmFAkxlpY struct {
	KeyValueOfstringArrayOfTrackingResultmFAkxlpY []struct {
		Key   string                 `xml:"Key" json:"Key"`
		Value *arrayOfTrackingResult `xml:"Value" json:"Value"`
	} `xml:"KeyValueOfstringArrayOfTrackingResultmFAkxlpY" json:"KeyValueOfstringArrayOfTrackingResultmFAkxlpY"`
}

// TrackingResult tracking result
type trackingResult struct {
	WaybillNumber     string `xml:"WaybillNumber" json:"WaybillNumber"`
	UpdateCode        string `xml:"UpdateCode" json:"UpdateCode"`
	UpdateDescription string `xml:"UpdateDescription" json:"UpdateDescription"`
	UpdateDateTime    string `xml:"UpdateDateTime" json:"UpdateDateTime"`

	Comments         string `xml:"Comments" json:"Comments"`
	ProblemCode      string `xml:"ProblemCode" json:"ProblemCode"`
	GrossWeight      string `xml:"GrossWeight" json:"GrossWeight"`
	ChargeableWeight string `xml:"ChargeableWeight" json:"ChargeableWeight"`
	WeightUnit       string `xml:"WeightUnit" json:"WeightUnit"`
}

// TrackingResult tracking result
type TrackingResult struct {
	WaybillNumber     string    `xml:"WaybillNumber" json:"WaybillNumber"`
	UpdateCode        string    `xml:"UpdateCode" json:"UpdateCode"`
	UpdateDescription string    `xml:"UpdateDescription" json:"UpdateDescription"`
	UpdateDateTime    time.Time `xml:"UpdateDateTime" json:"UpdateDateTime"`

	Comments         string `xml:"Comments" json:"Comments"`
	ProblemCode      string `xml:"ProblemCode" json:"ProblemCode"`
	GrossWeight      string `xml:"GrossWeight" json:"GrossWeight"`
	ChargeableWeight string `xml:"ChargeableWeight" json:"ChargeableWeight"`
	WeightUnit       string `xml:"WeightUnit" json:"WeightUnit"`
}

type shipmentTrackingRequest struct {
	XMLName                   xml.Name       `xml:"http://ws.aramex.net/ShippingAPI/v1/ ShipmentTrackingRequest"`
	ClientInfo                *ClientInfo    `xml:"ClientInfo" json:"ClientInfo"`
	Transaction               *Transaction   `xml:"Transaction" json:"Transaction"`
	Shipments                 *arrayOfstring `xml:"Shipments" json:"Shipments"`
	GetLastTrackingUpdateOnly bool           `xml:"GetLastTrackingUpdateOnly" json:"GetLastTrackingUpdateOnly"`
}

type shipmentTrackingResponse struct {
	XMLName             xml.Name                                              `xml:"http://ws.aramex.net/ShippingAPI/v1/ ShipmentTrackingResponse"`
	Transaction         *Transaction                                          `xml:"Transaction" json:"Transaction"`
	Notifications       *arrayOfNotification                                  `xml:"Notifications" json:"Notifications"`
	HasErrors           bool                                                  `xml:"HasErrors" json:"HasErrors"`
	TrackingResults     *arrayOfKeyValueOfstringArrayOfTrackingResultmFAkxlpY `xml:"TrackingResults" json:"TrackingResults"`
	NonExistingWaybills *arrayOfstring                                        `xml:"NonExistingWaybills" json:"NonExistingWaybills"`
}

// ShipmentTrackingRequest request
type ShipmentTrackingRequest struct {
	GetLastTrackingUpdateOnly bool
	Shipments                 []string
}

// ShipmentTrackingResponse response
type ShipmentTrackingResponse struct {
	Transaction         *Transaction
	Notifications       []*Notification
	HasErrors           bool
	TrackingResults     map[string][]*TrackingResult
	NonExistingWaybills []string
}

// TrackShipments track shipments
func (a *Aramex) TrackShipments(ctx context.Context, request *ShipmentTrackingRequest) (*ShipmentTrackingResponse, error) {
	var resp = new(shipmentTrackingResponse)
	var req = &shipmentTrackingRequest{
		ClientInfo:                a.config.ClientInfo,
		GetLastTrackingUpdateOnly: request.GetLastTrackingUpdateOnly,
		Shipments: &arrayOfstring{
			Astring: request.Shipments,
		},
	}
	var err = a.clients[trackingService].CallContext(ctx, a.buildURL("TrackShipments"), req, resp)
	if err != nil {
		return nil, err
	}

	var response = &ShipmentTrackingResponse{
		Transaction:         resp.Transaction,
		HasErrors:           resp.HasErrors,
		NonExistingWaybills: resp.NonExistingWaybills.Astring,
		Notifications:       resp.Notifications.Notification,
		TrackingResults:     make(map[string][]*TrackingResult),
	}

	for _, v := range resp.TrackingResults.KeyValueOfstringArrayOfTrackingResultmFAkxlpY {

		var results = []*TrackingResult{}
		for _, value := range v.Value.TrackingResult {
			results = append(results, &TrackingResult{
				WaybillNumber:     value.WaybillNumber,
				UpdateCode:        value.UpdateCode,
				UpdateDescription: value.UpdateDescription,
				Comments:          value.Comments,
				ProblemCode:       value.ProblemCode,
				GrossWeight:       value.GrossWeight,
				ChargeableWeight:  value.ChargeableWeight,
				WeightUnit:        value.WeightUnit,
				UpdateDateTime:    a.parseTime(value.UpdateDateTime),
			})

		}
		response.TrackingResults[v.Key] = results
	}

	return response, nil
}
