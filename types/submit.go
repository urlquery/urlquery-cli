package types

import (
	"encoding/json"
)

type Submit struct {
	Url       string   `json:"url"`
	Tags      []string `json:"tags"`
	UserAgent string   `json:"useragent"`
	Referer   string   `json:"referer"`
}

func (j Submit) String() string {
	b, _ := json.MarshalIndent(j, "", "  ")

	return string(b)
}

type Queue struct {
	ReportID string `json:"report_id"`
	QueueID  string `json:"queue_id"`
	Status   string `json:"status"`
}
