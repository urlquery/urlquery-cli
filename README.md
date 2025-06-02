# urlquery-cli

A command-line tool for interacting with [urlquery.net](https://urlquery.net), allowing you to submit URLs for analysis, check URL reputations, and download reports, screenshots, and resources.

## Features

- Submit URLs for threat analysis
- Check reputation of URLs
- Retrieve scan results including:
  - JSON reports
  - Screenshots
  - Domain graphs
  - Specific resource files by hash
- Custom user-agent and access control
- Configurable via CLI flags or config file

---

## Installation

```bash
go install github.com/urlquery/urlquery-cli@latest
```

Or clone and build manually:

```bash
git clone https://github.com/urlquery/urlquery-cli.git
cd urlquery-cli
go build -o urlquery-cli .
```

---

## Configuration

Set your API key (required):

```bash
urlquery-cli config set apikey <your-api-key>
```

Optional settings:

```bash
urlquery-cli config set useragent "curl/7.81.0"
urlquery-cli config set access "private"
```

---

## Usage

### Submit a URL

```bash
urlquery-cli submit https://urlquery.net
```

You can configure visibility and user-agent via config or flags.

### Check submission status

```bash
urlquery-cli submit status <queue_id>
```

### Check URL reputation

```bash
urlquery-cli reputation google.com
```

### Retrieve scan results

```bash
urlquery-cli report <report_id> report
urlquery-cli report <report_id> screenshot
urlquery-cli report <report_id> domain_graph
urlquery-cli report <report_id> resource <hash>
```

You can specify an output directory with `--output`:

```bash
urlquery-cli report <report_id> screenshot --output ./downloads
```

Get a quick summary of the data with `--summary`:

```bash
urlquery-cli report <report_id> report --summary
```


---

## Examples

```bash
urlquery-cli help
```
```console
A command-line interface for querying and analyzing URLs via the urlquery API.

Usage:
  urlquery-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      View or edit CLI config (e.g., API key, default output)
  help        Help about any command
  report      Fetch report details or download artifacts
  reputation  Check the reputation of a URL.
  search      Search for reports in urlquery.net.
  submit      Submit a URL for analysis.

Flags:
      --apikey string   API Key (can also be set via config file or URLQUERY_APIKEY env var)
      --config string   Path to config file (default is $HOME/.urlquery-cli.yaml)
  -h, --help            help for urlquery-cli
      --output string   Location to store downloaded data (reports, screenshots, files)
      --summary         Show a summary output instead of full json

Use "urlquery-cli [command] --help" for more information about a command.
```

Submit URL
```bash
urlquery-cli submit https://urlquery.net
```

Get report
```bash
urlquery-cli report 5e085255-6d43-4dfb-a2cf-add81f84a67d report
```

Get summary of a report
```bash
urlquery-cli report 5e085255-6d43-4dfb-a2cf-add81f84a67d report --summary
```
```console
📝 Report Summary:  5e085255-6d43-4dfb-a2cf-add81f84a67d
🔗 Submitted URL:   t7.news.rs-email.com/r/?bid=613924440&cid=DM93150&cm_mmc=IT-EM-_-RSN_20180228-_-DM93150-_-FOUR_PROD_URL_B&id=h2497be58,a5f6514,a5f6523&p1=lezandieolivier.com/cache/authenticate/USgShXfEEslS/Y2FuZGVyc29uQHNsdXJwbWFpbC5uZXQ=
🔗 Final URL:       2ssvy.mibkenns.es/YMVRLKQSXTGOWVCAJMQKJFDUf3i98w1o2c00a3nujoply?MUYECGEORKSMUTFVOZXCAUAWBLJML
📄 Webpage Title:   Enter Safe Account
🚨 Detections:      62
🏷️  Tags:            microsoft phishing suspicious tycoon
🌐 HTTP Requests:   52

🌍 Domain Summary:
FQDN                                                                Registered   First Seen    Last Seen   RX Bytes   TX Bytes Alerts
t7.news.rs-email.com                                                2010-06-08   2025-06-02   2025-06-02      733 B      690 B      0
get.geojs.io                                                        2017-02-18   2017-03-30   2025-05-29     1.5 kB      491 B      0
challenges.cloudflare.com                                           2009-02-17   2021-10-20   2025-05-28      97 kB      916 B      0
3ce1yhutflrpv0rnuhyguvmwd6rdozj3epkgmyilpruy88w1fm1usua.zpjgkd.es      unknown   2025-06-02   2025-06-02     1.2 kB      662 B      1
digq0.ugyqwmm.es                                                       unknown   2025-06-02   2025-06-02      566 B      450 B      1
objects.githubusercontent.com                                       2014-02-06   2021-11-01   2025-05-28      11 kB      891 B      0
2ssvy.mibkenns.es                                                      unknown   2025-06-02   2025-06-02     1.0 MB      35 kB     50
cdnjs.cloudflare.com                                                2009-02-17   2012-05-23   2025-05-28     247 kB     2.3 kB      0
ok4static.oktacdn.com                                               2014-11-11   2018-06-15   2025-05-28     268 kB     2.0 kB      0
unpkg.com                                                           2016-01-06   2016-01-07   2025-05-28     6.3 kB     1.3 kB      0
github.com                                                          2007-10-09   2016-07-13   2025-05-28      15 kB      456 B      0
lezandieolivier.com                                                 2020-11-26   2025-06-02   2025-06-02      324 B      634 B      0
code.jquery.com                                                     2005-12-10   2012-05-21   2025-05-28     270 kB     1.3 kB      0
---
```

## Development

Build locally:

```bash
go build -o urlquery-cli .
```

Run a command:

```bash
./urlquery-cli help
```

---

## License

MIT © [urlquery.net](https://urlquery.net)
