# urlquery-cli v0.1

A command-line interface for interacting with [urlquery.net](https://urlquery.net), allowing you to submit URLs for analysis, check their reputation, and download scan reports or artifacts.

## Features

- **Submit URLs** for sandbox analysis  
- **Check reputation** of URLs  
- **Download reports**, screenshots, domain graphs, and resources
- **API key** authentication  
- Configurable defaults  

## Installation

```sh
git clone https://github.com/urlquery/urlquery-cli.git
cd urlquery-cli
go build -o urlquery-cli
```

## Configuration

Set configuration options:

```sh
urlquery-cli config set apikey <your-api-key>
urlquery-cli config set access public         # Options: public, restricted, private
urlquery-cli config set useragent "Custom UserAgent/1.0"
urlquery-cli config set output ./downloads
```

Unset a config value:

```sh
urlquery-cli config unset useragent
```

## Usage

### Submit a URL

```sh
urlquery-cli submit http://urlquery.net
```

### Check submit status

```sh
urlquery-cli submit status <queue_id>
```

### Fetch a report

```sh
urlquery-cli report <report_id> report         # JSON report
urlquery-cli report <report_id> screenshot     # Screenshot image
urlquery-cli report <report_id> domain_graph   # Domain graph
urlquery-cli report <report_id> resource <hash> # Download resource file
```

### Check URL reputation

```sh
urlquery-cli reputation http://example.com
```

## Help

Each command has a built-in `--help`:

```sh
urlquery-cli report --help
```

## License

MIT
