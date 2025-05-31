package api

import (
	"fmt"
)

func Search(query string, limit int, offset int) (*SearchReportResponse, error) {
	return DefaultClient.Search(query, limit, offset)
}

func (api httpClient) Search(query string, limit int, offset int) (*SearchReportResponse, error) {
	var reply SearchReportResponse

	endpoint := fmt.Sprintf("/public/v1/search/reports/?query=%s&limit=%d&offset=%d", query, limit, offset)
	resp, err := api.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	err = DecodeResponse(resp, &reply)
	return &reply, err
}
