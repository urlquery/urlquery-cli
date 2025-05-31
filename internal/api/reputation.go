package api

import (
	"fmt"
	"net/url"
)

type ReputationResult struct {
	Url     string `json:"url"`
	Verdict string `json:"verdict"`
}

func CheckReputation(query string) (*ReputationResult, error) {
	return DefaultClient.CheckReputation(query)
}

func (api httpClient) CheckReputation(query string) (*ReputationResult, error) {
	var reply ReputationResult

	endpoint := fmt.Sprintf("/public/v1/reputation/check/?query=%s", url.QueryEscape(query))
	resp, err := api.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	err = DecodeResponse(resp, &reply)
	return &reply, err
}
