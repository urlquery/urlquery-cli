package types

import (
	"encoding/json"
)

// Verdict can have the following values
//    malware, phishing, fraud, suspicoius

type ReputationCheck struct {
	Url     string           `json:"url"`
	Verdict string           `json:"verdict"`
	Details *ReputationEntry `json:"details"`
}

func (r ReputationCheck) String() string {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(b)
}

type ReputationEntry struct {
	Timestamp      string   `json:"timestamp"`
	Url            URL      `json:"url"`
	Verdict        string   `json:"verdict"`
	PhishingTarget string   `json:"phishing_target"`
	MalwareFamily  string   `json:"malware_family"`
	Alert          string   `json:"alert"`
	Tags           []string `json:"tags"`
	ReportID       string   `json:"report_id"`
}

func (r ReputationEntry) String() string {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
