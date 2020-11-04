package aramex

import "fmt"

// Notification notification
type Notification struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
}

// Error error
func (n *Notification) Error() string {
	return fmt.Sprintf("%s - %s", n.Code, n.Message)
}
