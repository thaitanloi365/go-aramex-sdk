package aramex

import (
	"context"
	"encoding/xml"

	"github.com/thaitanloi365/go-aramex-sdk/types"
)

// ProcessedPickup pickup
type ProcessedPickup struct {
	ID                 string               `xml:"ID" json:"ID"`
	GUID               string               `xml:"GUID" json:"GUID"`
	Reference1         string               `xml:"Reference1" json:"Reference1"`
	Reference2         string               `xml:"Reference2" json:"Reference2"`
	ProcessedShipments []*ProcessedShipment `xml:"ProcessedShipments" json:"ProcessedShipments"`
}

// ExistingShipment shipment
type ExistingShipment struct {
	Number       string             `xml:"Number" json:"Number"`
	OriginEntity string             `xml:"OriginEntity" json:"OriginEntity"`
	ProductGroup types.ProductGroup `xml:"ProductGroup" json:"ProductGroup"`
}

// PickupItemDetail pickup item detail
type PickupItemDetail struct {
	ProductGroup       types.ProductGroup `xml:"ProductGroup" json:"ProductGroup"`
	ProductType        types.ProductType  `xml:"ProductType" json:"ProductType"`
	NumberOfShipments  int32              `xml:"NumberOfShipments" json:"NumberOfShipments"`
	PackageType        string             `xml:"PackageType" json:"PackageType"`
	Payment            types.PaymentType  `xml:"Payment" json:"Payment"`
	ShipmentWeight     *Weight            `xml:"ShipmentWeight" json:"ShipmentWeight"`
	ShipmentVolume     *Volume            `xml:"ShipmentVolume" json:"ShipmentVolume"`
	NumberOfPieces     int32              `xml:"NumberOfPieces" json:"NumberOfPieces"`
	CashAmount         *Money             `xml:"CashAmount" json:"CashAmount"`
	ExtraCharges       *Money             `xml:"ExtraCharges" json:"ExtraCharges"`
	ShipmentDimensions *Dimensions        `xml:"ShipmentDimensions" json:"ShipmentDimensions"`
	Comments           string             `xml:"Comments" json:"Comments"`
}

// Pickup pickup
type Pickup struct {
	PickupAddress     *Address            `xml:"PickupAddress" json:"PickupAddress"`
	PickupContact     *Contact            `xml:"PickupContact" json:"PickupContact"`
	PickupLocation    string              `xml:"PickupLocation" json:"PickupLocation"`
	PickupDate        string              `xml:"PickupDate" json:"PickupDate"`
	ReadyTime         string              `xml:"ReadyTime" json:"ReadyTime"`
	LastPickupTime    string              `xml:"LastPickupTime" json:"LastPickupTime"`
	ClosingTime       string              `xml:"ClosingTime" json:"ClosingTime"`
	Comments          string              `xml:"Comments" json:"Comments"`
	Reference1        string              `xml:"Reference1" json:"Reference1"`
	Reference2        string              `xml:"Reference2" json:"Reference2"`
	Vehicle           types.VehicleType   `xml:"Vehicle" json:"Vehicle"`
	Shipments         []*Shipment         `xml:"Shipments" json:"Shipments"`
	PickupItems       []*PickupItemDetail `xml:"PickupItems" json:"PickupItems"`
	Status            types.PickupStatus  `xml:"Status" json:"Status"`
	ExistingShipments []*ExistingShipment `xml:"ExistingShipments" json:"ExistingShipments"`
	Branch            string              `xml:"Branch" json:"Branch"`
	RouteCode         string              `xml:"RouteCode" json:"RouteCode"`
	Dispatcher        int32               `xml:"Dispatcher" json:"Dispatcher"`
}

type arrayOfPickupItemDetail struct {
	PickupItemDetail []*PickupItemDetail `xml:"PickupItemDetail" json:"PickupItemDetail"`
}

type arrayOfExistingShipment struct {
	ExistingShipment []*ExistingShipment `xml:"ExistingShipment" json:"ExistingShipment"`
}

type pickup struct {
	PickupAddress     *Address                 `xml:"PickupAddress" json:"PickupAddress"`
	PickupContact     *Contact                 `xml:"PickupContact" json:"PickupContact"`
	PickupLocation    string                   `xml:"PickupLocation" json:"PickupLocation"`
	PickupDate        string                   `xml:"PickupDate,omitempty" json:"PickupDate"`
	ReadyTime         string                   `xml:"ReadyTime,omitempty" json:"ReadyTime"`
	LastPickupTime    string                   `xml:"LastPickupTime,omitempty" json:"LastPickupTime"`
	ClosingTime       string                   `xml:"ClosingTime,omitempty" json:"ClosingTime"`
	Comments          string                   `xml:"Comments" json:"Comments"`
	Reference1        string                   `xml:"Reference1" json:"Reference1"`
	Reference2        string                   `xml:"Reference2" json:"Reference2"`
	Vehicle           string                   `xml:"Vehicle" json:"Vehicle"`
	Shipments         *arrayOfShipment         `xml:"Shipments" json:"Shipments"`
	PickupItems       *arrayOfPickupItemDetail `xml:"PickupItems" json:"PickupItems"`
	Status            types.PickupStatus       `xml:"Status" json:"Status"`
	ExistingShipments *arrayOfExistingShipment `xml:"ExistingShipments" json:"ExistingShipments"`
	Branch            string                   `xml:"Branch" json:"Branch"`
	RouteCode         string                   `xml:"RouteCode" json:"RouteCode"`
	Dispatcher        int32                    `xml:"Dispatcher" json:"Dispatcher"`
}

type processedPickup struct {
	ID                 string                    `xml:"ID" json:"ID"`
	GUID               string                    `xml:"GUID" json:"GUID"`
	Reference1         string                    `xml:"Reference1" json:"Reference1"`
	Reference2         string                    `xml:"Reference2" json:"Reference2"`
	ProcessedShipments *arrayOfProcessedShipment `xml:"ProcessedShipments" json:"ProcessedShipments"`
}

type pickupCreationRequest struct {
	XMLName     xml.Name     `xml:"http://ws.aramex.net/ShippingAPI/v1/ PickupCreationRequest"`
	ClientInfo  *ClientInfo  `xml:"ClientInfo" json:"ClientInfo"`
	Transaction *Transaction `xml:"Transaction" json:"Transaction"`
	Pickup      *pickup      `xml:"Pickup" json:"Pickup"`
	LabelInfo   *LabelInfo   `xml:"LabelInfo" json:"LabelInfo"`
}

type pickupCreationResponse struct {
	XMLName         xml.Name             `xml:"http://ws.aramex.net/ShippingAPI/v1/ PickupCreationResponse"`
	Transaction     *Transaction         `xml:"Transaction" json:"Transaction"`
	Notifications   *arrayOfNotification `xml:"Notifications" json:"Notifications"`
	HasErrors       bool                 `xml:"HasErrors" json:"HasErrors"`
	ProcessedPickup *processedPickup     `xml:"ProcessedPickup" json:"ProcessedPickup"`
}

// PickupCreationRequest request
type PickupCreationRequest struct {
	Pickup    *Pickup    `xml:"Pickup" json:"Pickup"`
	LabelInfo *LabelInfo `xml:"LabelInfo" json:"LabelInfo"`
}

// PickupCreationResponse response
type PickupCreationResponse struct {
	Transaction     *Transaction     `xml:"Transaction" json:"Transaction"`
	Notifications   Notifications    `xml:"Notifications" json:"Notifications"`
	HasErrors       bool             `xml:"HasErrors" json:"HasErrors"`
	ProcessedPickup *ProcessedPickup `xml:"ProcessedPickup" json:"ProcessedPickup"`
}

// CreatePickup create pickup
func (a *Aramex) CreatePickup(ctx context.Context, request *PickupCreationRequest) (*PickupCreationResponse, error) {
	var resp = new(pickupCreationResponse)
	var req = &pickupCreationRequest{
		ClientInfo: a.config.ClientInfo,
		LabelInfo:  request.LabelInfo,
		Pickup: &pickup{
			PickupAddress:  request.Pickup.PickupAddress,
			Branch:         request.Pickup.Branch,
			ClosingTime:    request.Pickup.ClosingTime,
			Comments:       request.Pickup.Comments,
			Dispatcher:     request.Pickup.Dispatcher,
			LastPickupTime: request.Pickup.LastPickupTime,
			PickupContact:  request.Pickup.PickupContact,
			PickupDate:     request.Pickup.PickupDate,
			PickupLocation: request.Pickup.PickupLocation,
			Reference1:     request.Pickup.Reference1,
			Reference2:     request.Pickup.Reference2,
			ReadyTime:      request.Pickup.ReadyTime,
			RouteCode:      request.Pickup.RouteCode,
			Status:         request.Pickup.Status,
		},
	}
	if len(request.Pickup.ExistingShipments) > 0 {
		req.Pickup.ExistingShipments = &arrayOfExistingShipment{
			ExistingShipment: []*ExistingShipment{},
		}
		for _, es := range request.Pickup.ExistingShipments {
			req.Pickup.ExistingShipments.ExistingShipment = append(req.Pickup.ExistingShipments.ExistingShipment, a.toExistingPickupRequest(es))
		}
	}

	if req.LabelInfo == nil && a.config.DefaultSetting != nil {
		req.LabelInfo = a.config.DefaultSetting.LabelInfo
	}

	if len(request.Pickup.PickupItems) > 0 {
		req.Pickup.PickupItems = &arrayOfPickupItemDetail{
			PickupItemDetail: []*PickupItemDetail{},
		}

		for _, item := range request.Pickup.PickupItems {
			var pickupItem = a.toPickupItemRequest(item)
			req.Pickup.PickupItems.PickupItemDetail = append(req.Pickup.PickupItems.PickupItemDetail, pickupItem)
		}

	}

	if len(request.Pickup.Shipments) > 0 {
		req.Pickup.Shipments = &arrayOfShipment{
			Shipment: []*shipment{},
		}
		for _, s := range request.Pickup.Shipments {
			var sm = a.toShipmentRequest(s)
			req.Pickup.Shipments.Shipment = append(req.Pickup.Shipments.Shipment, sm)
		}
	}
	switch request.Pickup.Vehicle {
	case types.VehicleTypeMedium:
		req.Pickup.Vehicle = "Medium (regular card or small van)"
	case types.VehicleTypeLarge:
		req.Pickup.Vehicle = "Large (van)"
	default:
		req.Pickup.Vehicle = "Small (regular bike or car)"
	}

	var err = a.clients[shippingService].CallContext(ctx, a.buildURL("CreatePickup"), req, resp)
	if err != nil {
		return nil, err
	}

	var response = &PickupCreationResponse{
		Transaction:   resp.Transaction,
		HasErrors:     resp.HasErrors,
		Notifications: resp.Notifications.Notification,
	}

	if resp.ProcessedPickup != nil {
		response.ProcessedPickup = &ProcessedPickup{
			GUID:       resp.ProcessedPickup.GUID,
			ID:         resp.ProcessedPickup.ID,
			Reference1: resp.ProcessedPickup.Reference1,
			Reference2: resp.ProcessedPickup.Reference2,
		}

		if resp.ProcessedPickup.ProcessedShipments != nil {
			for _, shipment := range resp.ProcessedPickup.ProcessedShipments.ProcessedShipment {
				response.ProcessedPickup.ProcessedShipments = append(response.ProcessedPickup.ProcessedShipments, &ProcessedShipment{
					ForeignHAWB:         shipment.ForeignHAWB,
					HasErrors:           shipment.HasErrors,
					ID:                  shipment.ID,
					Notifications:       shipment.Notifications.Notification,
					Reference1:          shipment.Reference1,
					Reference2:          shipment.Reference2,
					Reference3:          shipment.Reference3,
					ShipmentDetails:     shipment.ShipmentDetails,
					ShipmentAttachments: shipment.ShipmentAttachments.ProcessedShipmentAttachment,
					ShipmentLabel:       shipment.ShipmentLabel,
				})
			}

		}
	}

	return response, nil
}
