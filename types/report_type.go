package types

// ReportType type
type ReportType string

// All report types
var (
	ReportTypeURL ReportType = "URL"
	ReportTypeRPT ReportType = "RPT"
)

func (t ReportType) String() string {
	return string(t)
}
