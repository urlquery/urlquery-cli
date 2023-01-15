package public

import (
	"fmt"
	"strings"
	"urlquery-cli/types"
)

//-------------------------------
// SubmitURL
//-------------------------------

func SubmitURL(url string) (string, error) {
	return NewDefaultRequest().SubmitURL(url)
}

func (api apiRequest) SubmitURL(url string) (string, error) {
	var reply string
	var submit types.Submit

	submit.Url = url

	apiurl := "https://api.urlquery.net/public/v1/submit/url"
	data, err := apiRequestHandle("POST", apiurl, strings.NewReader(submit.String()), api.key)

	reply = string(data)
	return reply, err
}

//-------------------------------
// GetQueueStatus
//-------------------------------

func GetQueueStatus(queue_id string) (string, error) {
	return NewDefaultRequest().GetQueueStatus(queue_id)
}

func (api apiRequest) GetQueueStatus(queue_id string) (string, error) {
	var reply string

	url := fmt.Sprintf("https://api.urlquery.net/public/v1/submit/status/%s", queue_id)
	data, err := apiRequestHandle("GET", url, nil, api.key)

	reply = string(data)

	return reply, err
}
