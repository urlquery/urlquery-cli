package public

import (
	"fmt"
)

//-------------------------------
// GetReport API
//-------------------------------

// GetReport grabs a report
func GetReport(report_id string) (string, error) {
	return NewDefaultRequest().GetReport(report_id)
}

func (api apiRequest) GetReport(report_id string) (string, error) {
	var reply string

	err := verifyReportID(report_id)

	if err == nil {
		var data []byte

		url := fmt.Sprintf("https://api.urlquery.net/public/v1/report/%s", report_id)
		data, err = apiRequestHandle("GET", url, nil, api.key)

		reply = string(data)
	}

	return reply, err
}

//-------------------------------
// GetScreenshot
//-------------------------------

func GetScreenshot(report_id string) ([]byte, error) {
	return NewDefaultRequest().GetScreenshot(report_id)
}

func (api apiRequest) GetScreenshot(report_id string) ([]byte, error) {
	var reply []byte

	err := verifyReportID(report_id)

	if err == nil {
		url := fmt.Sprintf("https://api.urlquery.net/public/v1/report/%s/screenshot", report_id)
		reply, err = apiRequestHandle("GET", url, nil, api.key)
	}

	return reply, err
}

//-------------------------------
// GetDomainGraph
//-------------------------------

func GetDomainGraph(report_id string) ([]byte, error) {
	return NewDefaultRequest().GetDomainGraph(report_id)
}

func (api apiRequest) GetDomainGraph(report_id string) ([]byte, error) {
	var reply []byte

	err := verifyReportID(report_id)

	if err == nil {
		url := fmt.Sprintf("https://api.urlquery.net/public/v1/report/%s/domain_graph", report_id)
		reply, err = apiRequestHandle("GET", url, nil, api.key)
	}

	return reply, err
}
