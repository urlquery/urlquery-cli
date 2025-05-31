# urlquery-cli

A command-line tool for interacting with [urlquery.net](https://urlquery.net), allowing you to submit URLs for sandbox analysis, check URL reputations, and download reports, screenshots, and resources.

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
urlquery-cli config set useragent "MyScanner/1.0"
urlquery-cli config set access "private"
```

---

## Usage

### Submit a URL

```bash
urlquery-cli submit https://example.com
```

You can configure visibility and user-agent via config or flags.

### Check submission status

```bash
urlquery-cli submit status <queue_id>
```

### Check URL reputation

```bash
urlquery-cli reputation https://example.com
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

---

## Examples

```bash
# Submit a URL for scanning
urlquery-cli submit https://test.com

# Check if the scan is finished
urlquery-cli submit status <queue_id>

# Get scan results
urlquery-cli report <report_id> report
```

---

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
