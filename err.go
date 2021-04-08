package aramex

import (
	"fmt"
	"strings"
)

// Notification notification
type Notification struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
}

// Error error
func (n *Notification) Error() string {
	return fmt.Sprintf("%s - %s", n.Code, n.Message)
}

type Notifications []*Notification

func (notifications Notifications) Error() string {
	var listStatusDescription = []string{}
	for _, n := range notifications {
		listStatusDescription = append(listStatusDescription, fmt.Sprintf("%s - %s", n.Code, n.Message))
	}
	var statusDescription = strings.Join(listStatusDescription, ";")

	return statusDescription
}

func (notifications Notifications) IsInvalidPostcode() bool {
	for _, notif := range notifications {
		if notif.Code == "ERR06" {
			return true
		}
	}
	return false
}

func (notifications Notifications) HasErrorCode(errorCode string) bool {
	for _, notif := range notifications {
		if notif.Code == errorCode {
			return true
		}
	}
	return false
}
