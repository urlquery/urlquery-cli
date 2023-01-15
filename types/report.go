package types

type ReportOverview struct {
	ID      string   `json:"report_id"`
	Version int      `json:"version"`
	Tags    []string `json:"tags"`
	Date    string   `json:"date"`
	Url     URL      `json:"url"`
	Ip      IP       `json:"ip"`

	Settings ReportSettings `json:"settings"`

	UrlQueryAlertCount  int `json:"urlquery_alert_count"`
	IDSAlertCount       int `json:"ids_alert_count"`
	BlocklistAlertCount int `json:"blocklist_alert_count"`

	NetworkSensors  []IDSSensor     `json:"ids_sensors"`
	UrlQueryAlerts  []UrlqueryAlert `json:"urlquery_alerts"`
	BlocklistAlerts []Blocklist     `json:"blocklists"`
}

type ReportSettings struct {
	UserAgent string `json:"useragent"`
	Referer   string `json:"referer"`
}

type URL struct {
	Addr   string `json:"addr"`
	Domain string `json:"domain"`
	FQDN   string `json:"fqdn"`
	TLD    string `json:"tld"`
}

type IP struct {
	Addr        string `json:"addr"`
	CountryCode string `json:"cc"`
	Country     string `json:"country"`
	ASN         int    `json:"asn"`
	AS          string `json:"as"`
}

type UrlqueryAlert struct {
	SensorName string   `json:"sensor_name"`
	Alert      string   `json:"alert"`
	Severity   int      `json:"severity"`
	Confidence string   `json:"confidence"`
	Verdict    string   `json:"verdict"`
	Comment    string   `json:"comment"`
	Tags       []string `json:"tags"`
}

type IDSAlert struct {
	SensorName string `json:"sensor_name"`
	Timestamp  string `json:"timestamp"`
	IpDst      IP     `json:"ip_dst"`
	IpSrc      IP     `json:"ip_src"`
	Severity   int    `json:"severity"`
	Alert      string `json:"alert"`
}

type IDSSensor struct {
	SensorName  string `json:"sensor_name"`
	Description string `json:"description"`

	Alerts []IDSAlert `json:"alerts"`
}

type Blocklist struct {
	BlocklistName string `json:"blocklist_name"`
	SensorName    string `json:"sensor_name"`
	Description   string `json:"description"`
	Link          string `json:"link"`
	Type          string `json:"type"`

	Alerts []BlocklistAlert `json:"alerts"`
}

type BlocklistAlert struct {
	SensorName  string `json:"sensor_name"`
	Type        string `json:"type"`
	Description string `json:"description"`

	ScanDate string             `json:"scan_date"`
	Trigger  string             `json:"trigger"`
	Severity int                `json:"severity"`
	Verdict  string             `json:"verdict"`
	Comment  string             `json:"comment"`
	Link     *string            `json:"link"`
	Meta     *map[string]string `json:"meta"`
}
