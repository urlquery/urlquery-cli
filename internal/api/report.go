package api

import (
	"fmt"
	"io"
)

func GetReport(report_id string) {
	DefaultClient.GetReport(report_id)

}

func (api httpClient) GetReport(report_id string) (*Report, error) {
	var reply Report

	endpoint := fmt.Sprintf("/public/v1/report/%s", report_id)
	resp, err := api.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	err = DecodeResponse(resp, &reply)
	return &reply, err
}

func (api httpClient) GetScreenshot(report_id string) ([]byte, error) {

	endpoint := fmt.Sprintf("/public/v1/report/%s/screenshot", report_id)
	resp, err := api.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	return data, err
}

func (api httpClient) GetDomainGraph(report_id string) ([]byte, error) {

	endpoint := fmt.Sprintf("/public/v1/report/%s/domain_graph", report_id)
	resp, err := api.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	return data, err
}
