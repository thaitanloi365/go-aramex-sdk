package aramex

import (
	"context"
	"testing"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/thaitanloi365/go-aramex-sdk/types"
)

func TestShippingGetLastShipmentsNumbersRange(t *testing.T) {
	var aramex = New(&Config{
		IsLive:     false,
		ClientInfo: DefaultClientInfo,
	})
	result, err := aramex.GetLastShipmentsNumbersRange(context.Background(), nil)
	assert.NoError(t, err)

	aramex.printJSON(&result)
}

func TestShippingReserveShipmentNumberRange(t *testing.T) {
	var aramex = New(&Config{
		IsLive:     false,
		ClientInfo: DefaultClientInfo,
	})
	result, err := aramex.ReserveShipmentNumberRange(context.Background(), nil)
	assert.NoError(t, err)

	aramex.printJSON(&result)
}

func TestShippingScheduleDelivery(t *testing.T) {
	var aramex = New(&Config{
		IsLive:     false,
		ClientInfo: DefaultClientInfo,
	})
	result, err := aramex.ScheduleDelivery(context.Background(), &ScheduledDeliveryRequest{})
	assert.NoError(t, err)

	aramex.printJSON(&result)
}

func TestShippingPrintLabel(t *testing.T) {
	var aramex = New(&Config{
		IsLive:     false,
		ClientInfo: DefaultClientInfo,
	})
	result, err := aramex.PrintLabel(context.Background(), &LabelPrintingRequest{
		ShipmentNumber: "42489016194",
	})
	assert.NoError(t, err)

	aramex.printJSON(&result)
}

func TestShippingCreateShipments(t *testing.T) {
	var aramex = New(&Config{
		IsLive:     false,
		ClientInfo: DefaultClientInfo,
	})

	// f, err := os.Open("/Users/triluong/Downloads/sample.pdf")
	// if err != nil {
	// 	panic(err)
	// }

	// Read entire JPG into byte slice.
	// reader := bufio.NewReader(f)
	// content, _ := ioutil.ReadAll(reader)

	// // Encode as base64.
	// encoded := base64.StdEncoding.EncodeToString(content)

	// bytes, err := ioutil.ReadFile("/Users/relia/Desktop/shipping-services-api-manual.pdf")
	// if err != nil {
	// 	panic(err)
	// }

	var rawData = []byte(`
	{
    "Shipments": [
        {
            "Reference1": "OR-TEST12-123456",
            "Reference2": "",
            "Reference3": "",
            "Shipper": {
                "Reference1": "",
                "Reference2": "",
                "AccountNumber": "20016",
                "PartyAddress": {
                    "Line1": " 76 Shenton, 76 Shenton Way, 079119",
                    "Line2": "",
                    "Line3": "",
                    "City": "Singapore",
                    "PostCode": "079119",
                    "CountryCode": "SG",
                    "POBox": "",
                    "Description": ""
                },
                "Contact": {
                    "Department": "",
                    "PersonName": "Tri Test",
                    "Title": "",
                    "CompanyName": "",
                    "PhoneNumber1": "+6565325277",
                    "PhoneNumber1Ext": "",
                    "PhoneNumber2": "",
                    "PhoneNumber2Ext": "",
                    "FaxNumber": "",
                    "CellPhone": "+6565325277",
                    "EmailAddress": "tritest@mailinator.com",
                    "Type": ""
                }
            },
            "Consignee": {
                "Reference1": "",
                "Reference2": "",
                "AccountNumber": "",
                "PartyAddress": {
                    "Line1": " 515, Pesiaran Sultan Salahuddin, 50480",
                    "Line2": "",
                    "Line3": "",
                    "City": "Kuala Lumpur",
                    "StateOrProvinceCode": "Kuala Lumpur",
                    "PostCode": "50480",
                    "CountryCode": "MY",
                    "POBox": "",
                    "Description": ""
                },
                "Contact": {
                    "Department": "",
                    "PersonName": "test_padad das c_01",
                    "Title": "",
                    "CompanyName": "test_padad das c_01",
                    "PhoneNumber1": "+606313412341",
                    "PhoneNumber1Ext": "",
                    "PhoneNumber2": "",
                    "PhoneNumber2Ext": "",
                    "FaxNumber": "",
                    "CellPhone": "+606313412341",
                    "EmailAddress": "tritest@mailinator.com",
                    "Type": ""
                }
            },
            "ShippingDateTime": "2022-03-02T07:00:00",
            "DueDate": "2022-03-02T07:00:00",
            "Comments": "",
            "PickupLocation": "Security",
            "OperationsInstructions": "conten as  dt_01 / xxx",
            "AccountingInstrcutions": "",
            "Details": {
                "ActualWeight": {
                    "Unit": "KG",
                    "Value": 1
                },
                "ChargeableWeight": {
                    "Unit": "KG",
                    "Value": 1
                },
                "DescriptionOfGoods": "",
                "GoodsOriginCountry": "SG",
                "NumberOfPieces": 2,
                "ProductGroup": "EXP",
                "ProductType": "PPX",
                "PaymentType": "",
                "PaymentOptions": "",
				"CustomsValueAmount": 12,
                "CashAdditionalAmountDescription": "",
                "ContainsDangerousGoods": false
            },
            "Attachments": [
                {
                    "FileName": "invoice_20220125-000001",
                    "FileExtension": ".pdf"
                }
            ],
            "ForeignHAWB": "",
            "PickupGUID": "",
            "Number": ""
        }
    ],
    "LabelInfo": {
        "ReportID": 9201,
        "ReportType": "URL"
    }
}
	`)

	var params *ShipmentCreationRequest

	jsoniter.Unmarshal(rawData, &params)

	aramex.printJSON(&params)

	result, err := aramex.CreateShipments(context.Background(), params)

	if err != nil {
		panic(err)
	}

	aramex.printJSON(&result)
}

func TestShippingCreatePickupFromExistingShipments(t *testing.T) {
	var aramex = New(&Config{
		IsLive:     false,
		ClientInfo: DefaultClientInfo,
	})

	loc, _ := time.LoadLocation("Asia/Singapore")

	var current = time.Now()
	var readyTime = time.Date(current.Year(), current.Month(), current.Day(), 18, 0, 0, 0, loc)
	if readyTime.Before(current) {
		readyTime = readyTime.AddDate(0, 0, 1)
	}
	var pickupDate = current
	if pickupDate.Before(readyTime) {
		pickupDate = readyTime.Add(time.Hour * 2)
	}
	var lastPickupTime = readyTime.AddDate(0, 0, 1)
	var closingTime = readyTime.AddDate(0, 0, 1)
	result, err := aramex.CreatePickup(context.Background(), &PickupCreationRequest{
		Pickup: &Pickup{
			ClosingTime:    closingTime.Format("2006-01-02T15:04:05"),
			PickupDate:     pickupDate.Format("2006-01-02T15:04:05"),
			LastPickupTime: lastPickupTime.Format("2006-01-02T15:04:05"),
			ReadyTime:      readyTime.Format("2006-01-02T15:04:05"),
			Status:         types.PickupStatusReady,
			PickupLocation: "At reception",
			Reference1:     "001122",
			PickupContact: &Contact{
				PersonName:   "Loi",
				PhoneNumber1: "+12345678910",
				CellPhone:    "+12345678910",
				Title:        "Pick up",
				CompanyName:  "Relia",
			},
			PickupAddress: &Address{
				City:        "Singapore",
				CountryCode: types.CountryCodeSG,
				PostCode:    "139945",
			},
			PickupItems: []*PickupItemDetail{
				{
					NumberOfPieces:    1,
					NumberOfShipments: 1,
					ShipmentWeight: &Weight{
						Unit:  types.WeightUnitKG,
						Value: 0.5,
					},
					ShipmentVolume: &Volume{
						Unit:  types.VolumeUnitCm3,
						Value: 10,
					},
				},
			},
		},
	})
	assert.NoError(t, err)

	aramex.printJSON(&result)

}
func TestShippingCreatePickup(t *testing.T) {
	var aramex = New(&Config{
		IsLive:     false,
		ClientInfo: DefaultClientInfo,
	})

	loc, _ := time.LoadLocation("Asia/Singapore")

	var current = time.Now()
	var readyTime = time.Date(current.Year(), current.Month(), current.Day(), 18, 0, 0, 0, loc)
	if readyTime.Before(current) {
		readyTime = readyTime.AddDate(0, 0, 1)
	}
	var pickupDate = current
	if pickupDate.Before(readyTime) {
		pickupDate = readyTime.Add(time.Hour * 2)
	}
	var lastPickupTime = readyTime.AddDate(0, 0, 1)
	var closingTime = readyTime.AddDate(0, 0, 1)
	result, err := aramex.CreatePickup(context.Background(), &PickupCreationRequest{
		Pickup: &Pickup{
			ClosingTime:    closingTime.Format("2006-01-02T15:04:05"),
			PickupDate:     pickupDate.Format("2006-01-02T15:04:05"),
			LastPickupTime: lastPickupTime.Format("2006-01-02T15:04:05"),
			ReadyTime:      readyTime.Format("2006-01-02T15:04:05"),
			Status:         types.PickupStatusReady,
			PickupLocation: "At reception",
			Reference1:     "001122",
			PickupContact: &Contact{
				PersonName:   "Loi",
				PhoneNumber1: "+12345678910",
				CellPhone:    "+12345678910",
				Title:        "Pick up",
				CompanyName:  "Relia",
			},
			PickupAddress: &Address{
				City:        "Singapore",
				CountryCode: types.CountryCodeSG,
				PostCode:    "139945",
			},
			PickupItems: []*PickupItemDetail{
				{
					NumberOfPieces:    1,
					NumberOfShipments: 1,
					ShipmentWeight: &Weight{
						Unit:  types.WeightUnitKG,
						Value: 0.5,
					},
					ShipmentVolume: &Volume{
						Unit:  types.VolumeUnitCm3,
						Value: 10,
					},
				},
			},

			Shipments: []*Shipment{
				{
					Shipper: &Party{
						Contact: &Contact{
							PersonName:   "Thai  Loi",
							PhoneNumber1: "+6568727288",
							CellPhone:    "+6568727288",
						},
						PartyAddress: &Address{
							City:        "Singapore",
							CountryCode: types.CountryCodeSG,
							PostCode:    "139945",
						},
					},
					Consignee: &Party{
						Contact: &Contact{
							PersonName:   "My My",
							PhoneNumber1: "+6568727288",
							CellPhone:    "+6568727288",
							CompanyName:  "Relia",
							EmailAddress: "thaitanloi365@gmail.com",
						},
						PartyAddress: &Address{
							Line1:       "540 Airport Road Paya Lebar, 539938, Singapore",
							City:        "Singapore",
							CountryCode: types.CountryCodeSG,
							PostCode:    "139945",
						},
					},
					Details: &ShipmentDetails{
						ActualWeight: &Weight{
							Unit:  types.WeightUnitKG,
							Value: 0.5,
						},
						NumberOfPieces:     1,
						DescriptionOfGoods: "Parcel",
						GoodsOriginCountry: types.CountryCodeSG,
						CashOnDeliveryAmount: &Money{
							CurrencyCode: types.CurrencyCodeSGD,
							Value:        12.4,
						},
					},

					DueDate: time.Now().AddDate(0, 0, 3).Format("2006-01-02T15:04:05"),
				},
			},
		},
	})
	assert.NoError(t, err)

	aramex.printJSON(&result)
}

func TestCancelPickup(t *testing.T) {

}
