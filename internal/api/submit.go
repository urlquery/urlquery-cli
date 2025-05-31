package api

import (
	"fmt"
	"strings"
)

func Submit(submit SubmitJob) (*QueuedJob, error) {
	return DefaultClient.Submit(submit)
}

func (api httpClient) Submit(submit SubmitJob) (*QueuedJob, error) {
	var queued_job QueuedJob

	endpoint := "/public/v1/submit/url"
	resp, err := api.DoRequest("POST", endpoint, strings.NewReader(submit.String()))
	if err != nil {
		return nil, err
	}

	err = DecodeResponse(resp, &queued_job)
	return &queued_job, err
}

func QueueStatus(queue_id string) (*QueuedJob, error) {
	return DefaultClient.QueueStatus(queue_id)
}

func (api httpClient) QueueStatus(queue_id string) (*QueuedJob, error) {
	var queued_job QueuedJob

	endpoint := fmt.Sprintf("/public/v1/submit/status/%s", queue_id)
	resp, err := api.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	err = DecodeResponse(resp, &queued_job)
	return &queued_job, err
}
