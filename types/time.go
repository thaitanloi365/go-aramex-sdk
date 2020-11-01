package types

import (
	"fmt"
	"time"
)

const ctLayout = "2006-01-02T15:04:05"

// CustomTime custom
type CustomTime time.Time

// MarshalJSON marshal
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(ct).Format(ctLayout))
	fmt.Println(stamp)
	return []byte(stamp), nil
}
