package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/urlquery/urlquery-cli/internal/api"
)

// Template for report summary
const reportSummaryTemplate = `
📝 Report ID     : {{.ID}}
📝 Created       : {{.Date}}
🔗 Submitted URL : {{.Url.Addr}}
🌐 IP      		 : {{ .Ip.Addr }} {{ countryFlag .Ip.CountryCode }}
🔗 Final URL     : {{.Final.Url.Addr}}
📄 Webpage Title : {{.Final.Title}}
🚨 Detections    : {{.Stats.AlertCount.Urlquery}}
🏷️  Tags         : {{join .Tags " "}}
🌐 HTTP Requests : {{len .HttpTransactions}}

🌍 Domain Summary:
{{- if .Summary}}
FQDN                                                      | Registered   | First Seen | Last Seen  | RX Bytes   | TX Bytes   | Alerts
{{- range .Summary}}
{{printf "%-57s | %-12s | %-10s | %-10s | %-10s | %-10s | %6d" .Fqdn .DomainRegistered (formatDate .FirstSeen) (formatDate .LastSeen) (humanizeBytes .ReceivedData) (humanizeBytes .SentData) .AlertCount}}
{{- end}}
{{- else}}
No domain summary available.
{{- end}}

🚨 URLQuery Detections:
{{- if .Sensors.UrlQueryAlerts }}
{{- range .Sensors.UrlQueryAlerts }}
   └─ {{ .Alert }}
{{- end }}
{{- end }}

🌍 HTTP Transactions:
{{ range .HttpTransactions }}
─────────────────────────────────────────────────────────────────────────────
🔗 URL       : {{ .Url.Schema }}://{{ .Url.Addr }}
🌐 IP        : {{ .Ip.Addr }} {{ countryFlag .Ip.CountryCode }}
🌐 ASN       : #{{ .Ip.ASN }} {{ .Ip.AS }}
📡 Method    : {{ .Request.Method }}
📥 Status    : {{ .Response.StatusCode }} {{ .Response.StatusText }}
⏱️  Duration  : {{ .TotalTimeUsed }}ms
🔐 Security  : {{ .SecurityState }}

📦 Response Content:
   └─ Size     : {{ humanizeBytes .Response.Content.Size }}
   └─ MIME     : {{ .Response.Content.MimeType }}
   └─ Magic    : {{ .Response.Content.Magic }}
   └─ MD5      : {{ .Response.Content.Md5 }}
   └─ SHA1     : {{ .Response.Content.Sha1 }}
   └─ SHA256   : {{ .Response.Content.Sha256 }}
   └─ SHA512   : {{ .Response.Content.Sha512 }}

🚨 Detections:
{{- if or .Alerts.IDSAlerts .Alerts.AnalyzerAlerts .Alerts.UrlqueryAlerts }}
{{- range .Alerts.IDSAlerts }}
   └─ IDS: {{ .Alert }}
{{ end -}}
{{- range .Alerts.AnalyzerAlerts }}
   └─ Analyzer: {{ .Alert }}
{{ end -}}
{{- range .Alerts.UrlqueryAlerts }}
   └─ Urlquery: {{ .Alert }}
{{ end -}}
{{- else }}
   └─ None
{{- end }}

{{ end }}
`

// Custom template functions
var templateFunctions = template.FuncMap{
	"join": strings.Join,
	"formatDate": func(dateStr string) string {
		t, err := time.Parse(time.RFC3339, dateStr)
		if err != nil {
			return dateStr
		}
		return t.Format("2006-01-02")
	},
	"humanizeBytes": func(bytes int) string {
		return humanize.Bytes(uint64(bytes))
	},
	"countryFlag": func(code string) string {
		code = strings.ToUpper(code)
		if len(code) != 2 {
			return ""
		}
		r1 := rune(code[0]) + 127397 // 'A' → 🇦 (U+1F1E6)
		r2 := rune(code[1]) + 127397
		return string([]rune{r1, r2})
	},
}

// SummarizeReport generates a formatted summary of a report using templates
func SummarizeReport(report *api.Report) string {
	// Parse the template
	tmpl, err := template.New("report").Funcs(templateFunctions).Parse(reportSummaryTemplate)
	if err != nil {
		return fmt.Sprintf("Error parsing template: %v", err)
	}

	// Create a buffer to store the template output
	var buf bytes.Buffer

	// Execute the template with the report data
	if err := tmpl.Execute(&buf, report); err != nil {
		return fmt.Sprintf("Error executing template: %v", err)
	}

	return buf.String()
}
