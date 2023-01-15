package public

import (
	"fmt"
	"net/url"
)

// GetReport grabs a report
func CheckReputation(url string) (string, error) {
	return NewDefaultRequest().CheckReputation(url)
}

func (api apiRequest) CheckReputation(check string) (string, error) {
	var reply string

	apiurl := fmt.Sprintf("https://api.urlquery.net/public/v1/reputation/check/?query=%s", url.QueryEscape(check))
	data, err := apiRequestHandle("GET", apiurl, nil, api.key)

	reply = string(data)

	return reply, err
}
