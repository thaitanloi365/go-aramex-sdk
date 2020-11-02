package types

import (
	"encoding/xml"
	"time"
)

const ctLayout = "2006-01-02T15:04:05"

// CustomTime custom
type CustomTime struct {
	time.Time
}

// UnmarshalXML unmarshal
func (c *CustomTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, _ := time.Parse(ctLayout, v)
	*c = CustomTime{parse}
	return nil
}

// UnmarshalXMLAttr unmarshal
func (c *CustomTime) UnmarshalXMLAttr(attr xml.Attr) error {
	parse, _ := time.Parse(ctLayout, attr.Value)
	*c = CustomTime{parse}
	return nil
}
