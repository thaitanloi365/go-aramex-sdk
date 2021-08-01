package aramex

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hooklift/gowsdl/soap"
	"github.com/thaitanloi365/go-aramex-sdk/types"
)

// Aramex aramex
type Aramex struct {
	client  *ClientInfo
	config  *Config
	logger  Logger
	baseURL string
	clients map[service]*soap.Client
}

// New instance
func New(config *Config) *Aramex {
	if config.DefaultSetting == nil {
		config.DefaultSetting = DefaultSetting
	}

	var services = []service{
		locationService,
		rateService,
		shippingService,
		trackingService,
	}
	var clients = map[service]*soap.Client{}
	for _, ser := range services {
		clients[ser] = soap.NewClient(config.getEndpoint(ser))
	}

	// http://ws.aramex.net/ShippingAPI.V2/Shipping/Service_1_0.svc?wsdl
	var instance = &Aramex{
		config:  config,
		client:  config.ClientInfo,
		clients: clients,
		baseURL: "http://ws.aramex.net/ShippingAPI/v1/Service_1_0",
	}

	if config.Logger != nil {
		instance.logger = config.Logger
	} else {
		instance.logger = log.New(os.Stdout, "\r\n", 0)
	}

	return instance

}

// GetConfig get config
func (a *Aramex) GetConfig() *Config {
	return a.config
}

func (a *Aramex) buildURL(path string) string {
	return fmt.Sprintf("%s/%s", a.baseURL, path)
}

func (a *Aramex) printXML(in interface{}) {
	data, _ := xml.MarshalIndent(&in, "", "    ")
	a.logger.Printf("%s", data)
}

func (a *Aramex) printJSON(in interface{}) {
	data, _ := json.MarshalIndent(&in, "", "    ")
	a.logger.Printf("%s", data)
}

func (a *Aramex) toShipmentDetailsRequest(s *ShipmentDetails) *shipmentDetails {
	var sm = &shipmentDetails{
		Dimensions:                      s.Dimensions,
		ActualWeight:                    s.ActualWeight,
		CashAdditionalAmount:            s.CashOnDeliveryAmount,
		CashAdditionalAmountDescription: s.CashAdditionalAmountDescription,
		CashOnDeliveryAmount:            s.CashOnDeliveryAmount,
		ChargeableWeight:                s.ChargeableWeight,
		CollectAmount:                   s.CollectAmount,
		ContainsDangerousGoods:          s.ContainsDangerousGoods,
		CustomsValueAmount:              s.CustomsValueAmount,
		DeliveryInstructions:            s.DeliveryInstructions,
		DescriptionOfGoods:              s.DescriptionOfGoods,
		GoodsOriginCountry:              s.GoodsOriginCountry.String(),
		InsuranceAmount:                 s.InsuranceAmount,
		NumberOfPieces:                  s.NumberOfPieces,
		PaymentOptions:                  s.PaymentOptions.String(),
		PaymentType:                     s.PaymentType.String(),
		ProductGroup:                    s.ProductGroup.String(),
		ProductType:                     s.ProductType.String(),
		Services:                        strings.Join(s.Services.ListString(), ","),
	}

	if s.ChargeableWeight == nil {
		sm.ChargeableWeight = s.ActualWeight

	}

	if s.Dimensions == nil {
		sm.Dimensions = &Dimensions{
			Height: 0,
			Length: 0,
			Width:  0,
			Unit:   types.DimensionUnitCM,
		}
	}

	if len(s.AdditionalProperties) > 0 {
		sm.AdditionalProperties = &arrayOfAdditionalProperty{
			AdditionalProperty: s.AdditionalProperties,
		}
	}
	if sm.PaymentOptions == "" {
		sm.PaymentOptions = a.config.DefaultSetting.PaymentOptions.String()
	}

	if sm.PaymentType == "" {
		sm.PaymentType = a.config.DefaultSetting.PaymentType.String()
	}

	if sm.ProductGroup == "" {
		sm.ProductGroup = a.config.DefaultSetting.ProductGroup.String()
	}

	if sm.ProductType == "" {
		sm.ProductType = a.config.DefaultSetting.ProductType.String()
	}

	if sm.Services == "" {
		sm.Services = strings.Join(a.config.DefaultSetting.Services.ListString(), ",")
	}

	if len(s.Items) > 0 {
		sm.Items = &arrayOfShipmentItem{
			ShipmentItem: []*shipmentItem{},
		}
		for _, item := range s.Items {
			var it = &shipmentItem{
				Comments:         item.Comments,
				CommodityCode:    item.CommodityCode,
				ContainerNumber:  item.ContainerNumber,
				CountryOfOrigin:  item.CountryOfOrigin,
				CustomsValue:     item.CustomsValue,
				GoodsDescription: item.GoodsDescription,
				PackageType:      item.PackageType,
				Quantity:         item.Quantity,
				Reference:        item.Reference,
				Weight:           item.Weight,
			}
			if len(item.PiecesDimensions) > 0 {
				it.PiecesDimensions = &arrayOfDimensions{
					Dimensions: item.PiecesDimensions,
				}
			}
			sm.Items.ShipmentItem = append(sm.Items.ShipmentItem, it)
		}
	}

	return sm
}

func (a *Aramex) toShipmentRequest(s *Shipment) *shipment {
	var sm = &shipment{
		AccountingInstrcutions: s.AccountingInstrcutions,
		Comments:               s.Comments,
		Consignee:              s.Consignee,
		DueDate:                s.DueDate,
		ForeignHAWB:            s.ForeignHAWB,
		Number:                 s.Number,
		OperationsInstructions: s.OperationsInstructions,
		PickupGUID:             s.PickupGUID,
		PickupLocation:         s.PickupLocation,
		Reference1:             s.Reference1,
		Reference2:             s.Reference2,
		Reference3:             s.Reference3,
		Shipper:                s.Shipper,
		ShippingDateTime:       s.ShippingDateTime,
		ThirdParty:             s.ThirdParty,
	}
	if len(s.Attachments) > 0 {
		sm.Attachments = &arrayOfAttachment{
			Attachment: s.Attachments,
		}
	}

	if s.ScheduledDelivery != nil {
		sm.ScheduledDelivery = &scheduledDelivery{
			PreferredDeliveryDate: s.ScheduledDelivery.PreferredDeliveryDate,
			PreferredDeliveryTime: s.ScheduledDelivery.PreferredDeliveryTime,
		}
	}

	if s.Shipper.AccountNumber == "" {
		s.Shipper.AccountNumber = a.config.ClientInfo.AccountNumber
	}

	if s.Details != nil {
		sm.Details = a.toShipmentDetailsRequest(s.Details)
	}

	return sm
}

func (a *Aramex) toPickupItemRequest(item *PickupItemDetail) *PickupItemDetail {
	var pickupItem = &PickupItemDetail{
		ProductGroup:       item.ProductGroup,
		PackageType:        item.PackageType,
		Payment:            item.Payment,
		ProductType:        item.ProductType,
		CashAmount:         item.CashAmount,
		Comments:           item.Comments,
		ExtraCharges:       item.ExtraCharges,
		NumberOfPieces:     item.NumberOfPieces,
		NumberOfShipments:  item.NumberOfShipments,
		ShipmentDimensions: item.ShipmentDimensions,
		ShipmentVolume:     item.ShipmentVolume,
		ShipmentWeight:     item.ShipmentWeight,
	}

	if pickupItem.ProductGroup == "" {
		pickupItem.ProductGroup = a.config.DefaultSetting.ProductGroup
	}

	if pickupItem.Payment == "" {
		pickupItem.Payment = a.config.DefaultSetting.PaymentType
	}

	if pickupItem.ProductType == "" {
		pickupItem.ProductType = a.config.DefaultSetting.ProductType
	}

	if pickupItem.ShipmentDimensions == nil {
		pickupItem.ShipmentDimensions = &Dimensions{
			Height: 0,
			Length: 0,
			Unit:   types.DimensionUnitCM,
			Width:  0,
		}
	}

	if pickupItem.CashAmount == nil {
		pickupItem.CashAmount = &Money{
			CurrencyCode: types.CurrencyCodeSGD,
			Value:        0,
		}
	}

	if pickupItem.ExtraCharges == nil {
		pickupItem.ExtraCharges = &Money{
			CurrencyCode: types.CurrencyCodeSGD,
			Value:        0,
		}
	}

	return pickupItem
}

func (a *Aramex) toExistingPickupRequest(s *ExistingShipment) *ExistingShipment {
	var em = &ExistingShipment{
		Number:       s.Number,
		OriginEntity: s.OriginEntity,
		ProductGroup: s.ProductGroup,
	}

	if em.ProductGroup == "" {
		em.ProductGroup = a.config.DefaultSetting.ProductGroup
	}

	if em.OriginEntity == "" {
		em.OriginEntity = a.client.AccountEntity
	}

	return em
}
