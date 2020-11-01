package aramex

import (
	"context"
	"encoding/xml"
	"time"

	"github.com/thaitanloi365/go-aramex-sdk/types"
)

// ShipmentLabel labels
type ShipmentLabel struct {
	LabelURL          string `xml:"LabelURL" json:"LabelURL"`
	LabelFileContents []byte `xml:"LabelFileContents" json:"LabelFileContents"`
}

// ProcessedShipmentAttachment attachment
type ProcessedShipmentAttachment struct {
	Name string `xml:"Name" json:"Name"`
	Type string `xml:"Type" json:"Type"`
	Url  string `xml:"Url" json:"Url"`
}

// Attachment attachment
type Attachment struct {
	FileName      string `xml:"FileName" json:"FileName"`
	FileExtension string `xml:"FileExtension" json:"FileExtension"`
	FileContents  []byte `xml:"FileContents" json:"FileContents"`
}

// ShipmentDetails details
type ShipmentDetails struct {
	Dimensions                      *Dimensions             `xml:"Dimensions" json:"Dimensions"`
	ActualWeight                    *Weight                 `xml:"ActualWeight" json:"ActualWeight"`
	ChargeableWeight                *Weight                 `xml:"ChargeableWeight" json:"ChargeableWeight"`
	DescriptionOfGoods              string                  `xml:"DescriptionOfGoods" json:"DescriptionOfGoods"`
	GoodsOriginCountry              types.CountryCode       `xml:"GoodsOriginCountry" json:"GoodsOriginCountry"`
	NumberOfPieces                  int32                   `xml:"NumberOfPieces" json:"NumberOfPieces"`
	ProductGroup                    types.ProductGroup      `xml:"ProductGroup" json:"ProductGroup"`
	ProductType                     types.ProductType       `xml:"ProductType" json:"ProductType"`
	PaymentType                     types.PaymentType       `xml:"PaymentType" json:"PaymentType"`
	PaymentOptions                  types.PaymentTypeOption `xml:"PaymentOptions" json:"PaymentOptions"`
	Services                        types.ListServiceCode   `xml:"Services" json:"Services"`
	CustomsValueAmount              *Money                  `xml:"CustomsValueAmount" json:"CustomsValueAmount"`
	CashOnDeliveryAmount            *Money                  `xml:"CashOnDeliveryAmount" json:"CashOnDeliveryAmount"`
	InsuranceAmount                 *Money                  `xml:"InsuranceAmount" json:"InsuranceAmount"`
	CashAdditionalAmount            *Money                  `xml:"CashAdditionalAmount" json:"CashAdditionalAmount"`
	CashAdditionalAmountDescription string                  `xml:"CashAdditionalAmountDescription" json:"CashAdditionalAmountDescription"`
	CollectAmount                   *Money                  `xml:"CollectAmount" json:"CollectAmount"`
	Items                           []*ShipmentItem         `xml:"Items" json:"Items"`
	DeliveryInstructions            *DeliveryInstructions   `xml:"DeliveryInstructions" json:"DeliveryInstructions"`
	AdditionalProperties            []*AdditionalProperty   `xml:"AdditionalProperties" json:"AdditionalProperties"`
	ContainsDangerousGoods          bool                    `xml:"ContainsDangerousGoods" json:"ContainsDangerousGoods"`
}

// AdditionalProperty additional
type AdditionalProperty struct {
	CategoryName string `xml:"CategoryName" json:"CategoryName"`
	Name         string `xml:"Name" json:"Name"`
	Value        string `xml:"Value" json:"Value"`
}

// ShipmentItem shipment item
type ShipmentItem struct {
	PackageType string `xml:"PackageType" json:"PackageType"`

	Quantity int32 `xml:"Quantity" json:"Quantity"`

	Weight *Weight `xml:"Weight" json:"Weight"`

	Comments string `xml:"Comments" json:"Comments"`

	Reference string `xml:"Reference" json:"Reference"`

	PiecesDimensions []*Dimensions `xml:"PiecesDimensions" json:"PiecesDimensions"`

	CommodityCode string `xml:"CommodityCode" json:"CommodityCode"`

	GoodsDescription string `xml:"GoodsDescription" json:"GoodsDescription"`

	CountryOfOrigin string `xml:"CountryOfOrigin" json:"CountryOfOrigin"`

	CustomsValue *Money `xml:"CustomsValue" json:"CustomsValue"`

	ContainerNumber string `xml:"ContainerNumber" json:"ContainerNumber"`
}

// ProcessedShipmentDetails processed shipment
type ProcessedShipmentDetails struct {
	Origin               string  `xml:"Origin" json:"Origin"`
	Destination          string  `xml:"Destination" json:"Destination"`
	ChargeableWeight     *Weight `xml:"ChargeableWeight" json:"ChargeableWeight"`
	DescriptionOfGoods   string  `xml:"DescriptionOfGoods" json:"DescriptionOfGoods"`
	GoodsOriginCountry   string  `xml:"GoodsOriginCountry" json:"GoodsOriginCountry"`
	NumberOfPieces       int32   `xml:"NumberOfPieces" json:"NumberOfPieces"`
	ProductGroup         string  `xml:"ProductGroup" json:"ProductGroup"`
	ProductType          string  `xml:"ProductType" json:"ProductType"`
	PaymentType          string  `xml:"PaymentType" json:"PaymentType"`
	PaymentOptions       string  `xml:"PaymentOptions" json:"PaymentOptions"`
	CustomsValueAmount   *Money  `xml:"CustomsValueAmount" json:"CustomsValueAmount"`
	CashOnDeliveryAmount *Money  `xml:"CashOnDeliveryAmount" json:"CashOnDeliveryAmount"`
	InsuranceAmount      *Money  `xml:"InsuranceAmount" json:"InsuranceAmount"`
	CashAdditionalAmount *Money  `xml:"CashAdditionalAmount" json:"CashAdditionalAmount"`
	CollectAmount        *Money  `xml:"CollectAmount" json:"CollectAmount"`
	Services             string  `xml:"Services" json:"Services"`
}

type scheduledDelivery struct {
	PreferredDeliveryDate             time.Time `xml:"PreferredDeliveryDate" json:"PreferredDeliveryDate"`
	PreferredDeliveryTimeFrame_x0020_ string    `xml:"PreferredDeliveryTimeFrame_x0020_" json:"PreferredDeliveryTimeFrame_x0020_"`
	PreferredDeliveryTime             string    `xml:"PreferredDeliveryTime" json:"PreferredDeliveryTime"`
}

// DeliveryInstructions instructions
type DeliveryInstructions struct {
	Option    string `xml:"Option" json:"Option"`
	Reference string `xml:"Reference" json:"Reference"`
}

type arrayOfDimensions struct {
	Dimensions []*Dimensions `xml:"Dimensions" json:"Dimensions"`
}

type shipmentItem struct {
	PackageType string `xml:"PackageType" json:"PackageType"`

	Quantity int32 `xml:"Quantity" json:"Quantity"`

	Weight *Weight `xml:"Weight" json:"Weight"`

	Comments string `xml:"Comments" json:"Comments"`

	Reference string `xml:"Reference" json:"Reference"`

	PiecesDimensions *arrayOfDimensions `xml:"PiecesDimensions" json:"PiecesDimensions"`

	CommodityCode string `xml:"CommodityCode" json:"CommodityCode"`

	GoodsDescription string `xml:"GoodsDescription" json:"GoodsDescription"`

	CountryOfOrigin string `xml:"CountryOfOrigin" json:"CountryOfOrigin"`

	CustomsValue *Money `xml:"CustomsValue" json:"CustomsValue"`

	ContainerNumber string `xml:"ContainerNumber" json:"ContainerNumber"`
}

type arrayOfShipmentItem struct {
	ShipmentItem []*shipmentItem `xml:"ShipmentItem" json:"ShipmentItem"`
}

type arrayOfAdditionalProperty struct {
	AdditionalProperty []*AdditionalProperty `xml:"AdditionalProperty" json:"AdditionalProperty"`
}

type arrayOfProcessedShipmentAttachment struct {
	ProcessedShipmentAttachment []*ProcessedShipmentAttachment `xml:"ProcessedShipmentAttachment" json:"ProcessedShipmentAttachment"`
}

type arrayOfAttachment struct {
	Attachment []*Attachment `xml:"Attachment" json:"Attachment"`
}

type shipmentDetails struct {
	Dimensions                      *Dimensions                `xml:"Dimensions" json:"Dimensions"`
	ActualWeight                    *Weight                    `xml:"ActualWeight" json:"ActualWeight"`
	ChargeableWeight                *Weight                    `xml:"ChargeableWeight" json:"ChargeableWeight"`
	DescriptionOfGoods              string                     `xml:"DescriptionOfGoods" json:"DescriptionOfGoods"`
	GoodsOriginCountry              string                     `xml:"GoodsOriginCountry" json:"GoodsOriginCountry"`
	NumberOfPieces                  int32                      `xml:"NumberOfPieces" json:"NumberOfPieces"`
	ProductGroup                    string                     `xml:"ProductGroup" json:"ProductGroup"`
	ProductType                     string                     `xml:"ProductType" json:"ProductType"`
	PaymentType                     string                     `xml:"PaymentType" json:"PaymentType"`
	PaymentOptions                  string                     `xml:"PaymentOptions" json:"PaymentOptions"`
	CustomsValueAmount              *Money                     `xml:"CustomsValueAmount" json:"CustomsValueAmount"`
	CashOnDeliveryAmount            *Money                     `xml:"CashOnDeliveryAmount" json:"CashOnDeliveryAmount"`
	InsuranceAmount                 *Money                     `xml:"InsuranceAmount" json:"InsuranceAmount"`
	CashAdditionalAmount            *Money                     `xml:"CashAdditionalAmount" json:"CashAdditionalAmount"`
	CashAdditionalAmountDescription string                     `xml:"CashAdditionalAmountDescription" json:"CashAdditionalAmountDescription"`
	CollectAmount                   *Money                     `xml:"CollectAmount" json:"CollectAmount"`
	Services                        string                     `xml:"Services" json:"Services"`
	Items                           *arrayOfShipmentItem       `xml:"Items" json:"Items"`
	DeliveryInstructions            *DeliveryInstructions      `xml:"DeliveryInstructions" json:"DeliveryInstructions"`
	AdditionalProperties            *arrayOfAdditionalProperty `xml:"AdditionalProperties" json:"AdditionalProperties"`
	ContainsDangerousGoods          bool                       `xml:"ContainsDangerousGoods" json:"ContainsDangerousGoods"`
}

type shipment struct {
	Reference1             string             `xml:"Reference1" json:"Reference1"`
	Reference2             string             `xml:"Reference2" json:"Reference2"`
	Reference3             string             `xml:"Reference3" json:"Reference3"`
	Shipper                *Party             `xml:"Shipper" json:"Shipper"`
	Consignee              *Party             `xml:"Consignee" json:"Consignee"`
	ThirdParty             *Party             `xml:"ThirdParty" json:"ThirdParty"`
	ShippingDateTime       time.Time          `xml:"ShippingDateTime" json:"ShippingDateTime"`
	DueDate                time.Time          `xml:"DueDate" json:"DueDate"`
	Comments               string             `xml:"Comments" json:"Comments"`
	PickupLocation         string             `xml:"PickupLocation" json:"PickupLocation"`
	OperationsInstructions string             `xml:"OperationsInstructions" json:"OperationsInstructions"`
	AccountingInstrcutions string             `xml:"AccountingInstrcutions" json:"AccountingInstrcutions"`
	Details                *shipmentDetails   `xml:"Details" json:"Details"`
	Attachments            *arrayOfAttachment `xml:"Attachments" json:"Attachments"`
	ForeignHAWB            string             `xml:"ForeignHAWB" json:"ForeignHAWB"`
	TransportType_x0020_   int32              `xml:"TransportType_x0020_" json:"TransportType_x0020_"`
	PickupGUID             string             `xml:"PickupGUID" json:"PickupGUID"`
	Number                 string             `xml:"Number" json:"Number"`
	ScheduledDelivery      *scheduledDelivery `xml:"ScheduledDelivery" json:"ScheduledDelivery"`
}

type arrayOfShipment struct {
	Shipment []*shipment `xml:"Shipment" json:"Shipment"`
}

type processedShipment struct {
	ID                  string                              `xml:"ID" json:"ID"`
	Reference1          string                              `xml:"Reference1" json:"Reference1"`
	Reference2          string                              `xml:"Reference2" json:"Reference2"`
	Reference3          string                              `xml:"Reference3" json:"Reference3"`
	ForeignHAWB         string                              `xml:"ForeignHAWB" json:"ForeignHAWB"`
	HasErrors           bool                                `xml:"HasErrors" json:"HasErrors"`
	Notifications       *arrayOfNotification                `xml:"Notifications" json:"Notifications"`
	ShipmentLabel       *ShipmentLabel                      `xml:"ShipmentLabel" json:"ShipmentLabel"`
	ShipmentDetails     *ProcessedShipmentDetails           `xml:"ShipmentDetails" json:"ShipmentDetails"`
	ShipmentAttachments *arrayOfProcessedShipmentAttachment `xml:"ShipmentAttachments" json:"ShipmentAttachments"`
}

type arrayOfProcessedShipment struct {
	ProcessedShipment []*processedShipment `xml:"ProcessedShipment" json:"ProcessedShipment"`
}

type shipmentCreationRequest struct {
	XMLName xml.Name `xml:"http://ws.aramex.net/ShippingAPI/v1/ ShipmentCreationRequest"`

	ClientInfo *ClientInfo `xml:"ClientInfo" json:"ClientInfo"`

	Transaction *Transaction `xml:"Transaction" json:"Transaction"`

	Shipments *arrayOfShipment `xml:"Shipments" json:"Shipments"`

	LabelInfo *LabelInfo `xml:"LabelInfo" json:"LabelInfo"`
}

type shipmentCreationResponse struct {
	XMLName xml.Name `xml:"http://ws.aramex.net/ShippingAPI/v1/ ShipmentCreationResponse"`

	Transaction *Transaction `xml:"Transaction" json:"Transaction"`

	Notifications *arrayOfNotification `xml:"Notifications" json:"Notifications"`

	HasErrors bool `xml:"HasErrors" json:"HasErrors"`

	Shipments *arrayOfProcessedShipment `xml:"Shipments" json:"Shipments"`
}

// ProcessedShipment shipment
type ProcessedShipment struct {
	ID                  string                         `xml:"ID" json:"ID"`
	Reference1          string                         `xml:"Reference1" json:"Reference1"`
	Reference2          string                         `xml:"Reference2" json:"Reference2"`
	Reference3          string                         `xml:"Reference3" json:"Reference3"`
	ForeignHAWB         string                         `xml:"ForeignHAWB" json:"ForeignHAWB"`
	HasErrors           bool                           `xml:"HasErrors" json:"HasErrors"`
	Notifications       []*Notification                `xml:"Notifications" json:"Notifications"`
	ShipmentLabel       *ShipmentLabel                 `xml:"ShipmentLabel" json:"ShipmentLabel"`
	ShipmentDetails     *ProcessedShipmentDetails      `xml:"ShipmentDetails" json:"ShipmentDetails"`
	ShipmentAttachments []*ProcessedShipmentAttachment `xml:"ShipmentAttachments" json:"ShipmentAttachments"`
}

// ScheduledDelivery scheduled
type ScheduledDelivery struct {
	PreferredDeliveryDate time.Time `xml:"PreferredDeliveryDate" json:"PreferredDeliveryDate"`
	PreferredDeliveryTime string    `xml:"PreferredDeliveryTime" json:"PreferredDeliveryTime"`
}

// Shipment shipment
type Shipment struct {
	Reference1             string             `xml:"Reference1" json:"Reference1"`
	Reference2             string             `xml:"Reference2" json:"Reference2"`
	Reference3             string             `xml:"Reference3" json:"Reference3"`
	Shipper                *Party             `xml:"Shipper" json:"Shipper"`
	Consignee              *Party             `xml:"Consignee" json:"Consignee"`
	ThirdParty             *Party             `xml:"ThirdParty" json:"ThirdParty"`
	ShippingDateTime       time.Time          `xml:"ShippingDateTime" json:"ShippingDateTime"`
	DueDate                time.Time          `xml:"DueDate" json:"DueDate"`
	Comments               string             `xml:"Comments" json:"Comments"`
	PickupLocation         string             `xml:"PickupLocation" json:"PickupLocation"`
	OperationsInstructions string             `xml:"OperationsInstructions" json:"OperationsInstructions"`
	AccountingInstrcutions string             `xml:"AccountingInstrcutions" json:"AccountingInstrcutions"`
	Details                *ShipmentDetails   `xml:"Details" json:"Details"`
	Attachments            []*Attachment      `xml:"Attachments" json:"Attachments"`
	ForeignHAWB            string             `xml:"ForeignHAWB" json:"ForeignHAWB"`
	PickupGUID             string             `xml:"PickupGUID" json:"PickupGUID"`
	Number                 string             `xml:"Number" json:"Number"`
	ScheduledDelivery      *ScheduledDelivery `xml:"ScheduledDelivery" json:"ScheduledDelivery"`
}

// ShipmentCreationRequest request
type ShipmentCreationRequest struct {
	Shipments []*Shipment `xml:"Shipments" json:"Shipments"`
	LabelInfo *LabelInfo  `xml:"LabelInfo" json:"LabelInfo"`
}

// ShipmentCreationResponse response
type ShipmentCreationResponse struct {
	Transaction   *Transaction         `xml:"Transaction" json:"Transaction"`
	Notifications []*Notification      `xml:"Notifications" json:"Notifications"`
	HasErrors     bool                 `xml:"HasErrors" json:"HasErrors"`
	Shipments     []*ProcessedShipment `xml:"Shipments" json:"Shipments"`
}

// CreateShipments create shipments
func (a *Aramex) CreateShipments(ctx context.Context, request *ShipmentCreationRequest) (*ShipmentCreationResponse, error) {
	var resp = new(shipmentCreationResponse)
	var req = &shipmentCreationRequest{
		ClientInfo: a.config.ClientInfo,
		LabelInfo:  request.LabelInfo,
	}

	if req.LabelInfo == nil && a.config.DefaultSetting != nil {
		req.LabelInfo = a.config.DefaultSetting.LabelInfo
	}

	if len(request.Shipments) > 0 {
		req.Shipments = &arrayOfShipment{
			Shipment: []*shipment{},
		}
		for _, s := range request.Shipments {
			var sm = a.toShipmentRequest(s)
			req.Shipments.Shipment = append(req.Shipments.Shipment, sm)
		}
	}

	var err = a.clients[shippingService].CallContext(ctx, a.buildURL("CreateShipments"), req, resp)
	if err != nil {
		return nil, err
	}

	var response = &ShipmentCreationResponse{
		Transaction:   resp.Transaction,
		HasErrors:     resp.HasErrors,
		Notifications: resp.Notifications.Notification,
	}

	if resp.Shipments != nil {
		response.Shipments = []*ProcessedShipment{}
		for _, sm := range resp.Shipments.ProcessedShipment {
			var psm = &ProcessedShipment{
				ForeignHAWB:     sm.ForeignHAWB,
				ID:              sm.ID,
				HasErrors:       sm.HasErrors,
				Reference1:      sm.Reference1,
				Reference2:      sm.Reference2,
				Reference3:      sm.Reference3,
				ShipmentDetails: sm.ShipmentDetails,
				ShipmentLabel:   sm.ShipmentLabel,
			}
			if sm.Notifications != nil {
				psm.Notifications = sm.Notifications.Notification
			}

			if sm.ShipmentAttachments != nil {
				psm.ShipmentAttachments = sm.ShipmentAttachments.ProcessedShipmentAttachment
			}
			response.Shipments = append(response.Shipments, psm)
		}
	}

	return response, nil
}
