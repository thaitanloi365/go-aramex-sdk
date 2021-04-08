package aramex

import "github.com/thaitanloi365/go-aramex-sdk/types"

// ClientInfo client info
type ClientInfo struct {
	UserName           string `xml:"UserName,omitempty" json:"UserName,omitempty"`
	Password           string `xml:"Password,omitempty" json:"Password,omitempty"`
	Version            string `xml:"Version,omitempty" json:"Version,omitempty"`
	AccountNumber      string `xml:"AccountNumber,omitempty" json:"AccountNumber,omitempty"`
	AccountPin         string `xml:"AccountPin,omitempty" json:"AccountPin,omitempty"`
	AccountEntity      string `xml:"AccountEntity,omitempty" json:"AccountEntity,omitempty"`
	AccountCountryCode string `xml:"AccountCountryCode,omitempty" json:"AccountCountryCode,omitempty"`
}

// LabelInfo label info
type LabelInfo struct {
	ReportID   int
	ReportType types.ReportType
}

// Transaction transaction
type Transaction struct {
	Reference1 string `xml:"Reference1,omitempty" json:"Reference1,omitempty"`
	Reference2 string `xml:"Reference2,omitempty" json:"Reference2,omitempty"`
	Reference3 string `xml:"Reference3,omitempty" json:"Reference3,omitempty"`
	Reference4 string `xml:"Reference4,omitempty" json:"Reference4,omitempty"`
	Reference5 string `xml:"Reference5,omitempty" json:"Reference5,omitempty"`
}

// arrayOfNotification notifications
type arrayOfNotification struct {
	Notification Notifications `xml:"Notification" json:"Notification"`
}

// Dimensions dimension
type Dimensions struct {
	Length float64             `xml:"Length" json:"Length"`
	Width  float64             `xml:"Width" json:"Width"`
	Height float64             `xml:"Height" json:"Height"`
	Unit   types.DimensionUnit `xml:"Unit" json:"Unit"`
}

// Weight weight
type Weight struct {
	Unit  types.WeightUnit `xml:"Unit" json:"Unit"`
	Value float64          `xml:"Value" json:"Value"`
}

// Volume volume
type Volume struct {
	Unit  types.VolumeUnit `xml:"Unit" json:"Unit"`
	Value float64          `xml:"Value" json:"Value"`
}

// Money money
type Money struct {
	CurrencyCode types.CurrencyCode `xml:"CurrencyCode" json:"CurrencyCode"`
	Value        float64            `xml:"Value" json:"Value"`
}

// Party party
type Party struct {
	Reference1    string   `xml:"Reference1" json:"Reference1"`
	Reference2    string   `xml:"Reference2" json:"Reference2"`
	AccountNumber string   `xml:"AccountNumber,omitempty" json:"AccountNumber"`
	PartyAddress  *Address `xml:"PartyAddress" json:"PartyAddress"`
	Contact       *Contact `xml:"Contact" json:"Contact"`
}

// Address address
type Address struct {
	Line1               string            `xml:"Line1" json:"Line1"`
	Line2               string            `xml:"Line2" json:"Line2"`
	Line3               string            `xml:"Line3" json:"Line3"`
	City                string            `xml:"City" json:"City"`
	StateOrProvinceCode string            `xml:"StateOrProvinceCode,omitempty" json:"StateOrProvinceCode,omitempty"`
	PostCode            string            `xml:"PostCode" json:"PostCode"`
	CountryCode         types.CountryCode `xml:"CountryCode" json:"CountryCode"`
	Longitude           float64           `xml:"Longitude,omitempty" json:"Longitude,omitempty"`
	Latitude            float64           `xml:"Latitude,omitempty" json:"Latitude,omitempty"`
	BuildingNumber      string            `xml:"BuildingNumber,omitempty" json:"BuildingNumber,omitempty"`
	BuildingName        string            `xml:"BuildingName,omitempty" json:"BuildingName,omitempty"`
	Floor               string            `xml:"Floor,omitempty" json:"Floor,omitempty"`
	Apartment           string            `xml:"Apartment,omitempty" json:"Apartment,omitempty"`
	POBox               string            `xml:"POBox,omitempty" json:"POBox"`
	Description         string            `xml:"Description" json:"Description"`
}

// Contact contact
type Contact struct {
	Department      string `xml:"Department" json:"Department"`
	PersonName      string `xml:"PersonName" json:"PersonName"`
	Title           string `xml:"Title" json:"Title"`
	CompanyName     string `xml:"CompanyName" json:"CompanyName"`
	PhoneNumber1    string `xml:"PhoneNumber1" json:"PhoneNumber1"`
	PhoneNumber1Ext string `xml:"PhoneNumber1Ext" json:"PhoneNumber1Ext"`
	PhoneNumber2    string `xml:"PhoneNumber2" json:"PhoneNumber2"`
	PhoneNumber2Ext string `xml:"PhoneNumber2Ext" json:"PhoneNumber2Ext"`
	FaxNumber       string `xml:"FaxNumber" json:"FaxNumber"`
	CellPhone       string `xml:"CellPhone" json:"CellPhone"`
	EmailAddress    string `xml:"EmailAddress" json:"EmailAddress"`
	Type            string `xml:"Type" json:"Type"`
}

// Country count
type Country struct {
	Code                       string         `xml:"Code" json:"Code"`
	Name                       string         `xml:"Name" json:"Name"`
	IsoCode                    string         `xml:"IsoCode" json:"IsoCode"`
	StateRequired              bool           `xml:"StateRequired" json:"StateRequired"`
	PostCodeRequired           bool           `xml:"PostCodeRequired" json:"PostCodeRequired"`
	PostCodeRegex              *arrayOfstring `xml:"PostCodeRegex" json:"PostCodeRegex"`
	InternationalCallingNumber string         `xml:"InternationalCallingNumber" json:"InternationalCallingNumber"`
}

// State state
type State struct {
	Code string `xml:"Code" json:"Code"`
	Name string `xml:"Name" json:"Name"`
}

// Office office
type Office struct {
	Entity            string   `xml:"Entity" json:"Entity"`
	EntityDescription string   `xml:"EntityDescription" json:"EntityDescription"`
	OfficeType        string   `xml:"OfficeType" json:"OfficeType"`
	Address           *Address `xml:"Address" json:"Address"`
	Telephone         string   `xml:"Telephone" json:"Telephone"`
	WorkingDays       string   `xml:"WorkingDays" json:"WorkingDays"`
	WorkingHours      string   `xml:"WorkingHours" json:"WorkingHours"`
	Longtitude        float64  `xml:"Longtitude" json:"Longtitude"`
	Latitude          float64  `xml:"Latitude" json:"Latitude"`
}

type arrayOfState struct {
	State []*State `xml:"State" json:"State"`
}

type arrayOfOffice struct {
	Office []*Office `xml:"Office" json:"Office"`
}

// arrayOfstring array string
type arrayOfstring struct {
	Astring []string `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays string" json:"string"`
}

type arrayOfAddress struct {
	Address []*Address `xml:"Address" json:"Address"`
}

type arrayOfCountry struct {
	Country []*Country `xml:"Country" json:"Country"`
}
