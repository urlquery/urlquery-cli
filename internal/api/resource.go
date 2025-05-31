package api

import (
	"fmt"
	"io"
)

func GetResource(report_id string, hash string) {
	DefaultClient.GetResource(report_id, hash)
}

func (api httpClient) GetResource(report_id string, hash string) ([]byte, error) {

	endpoint := fmt.Sprintf("/public/v1/report/%s/resource/%s", report_id, hash)
	resp, err := api.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	return data, err
}
