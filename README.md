# urlquery-cli


A simple tool for interacting with urlquery.net from the command line. urlquery is a automated online service for analysing websites for malicious and suspicous content.

Currently submitting, retriving a report and checking reputation on URLs are supported through the public API



# Usage

```
Usage:
  urlquery-cli [command]

Available Commands:
  help        Help about any command
  rep-check   Check the reputation of a URL or IP
  report      Report a string
  screenshot  Download the screenshot of a report
  submit      Submit a URL

Flags:
      --apikey string   API key
  -h, --help            help for urlquery-cli
      --output string   Output (stdout, file) (default "stdout")

Use "urlquery-cli [command] --help" for more information about a command.
```

## Commands

### submit
Submitting a URL
```
$ urlquery-cli submit urlquery.net
URL submitted!
Current status: processing
```

Submission finished processing
```
$ urlquery-cli submit urlquery.net
URL submitted!
Report finished:
  - 688a87d0-0314-45d8-8676-ebcb7cb63dd8
  - https://urlquery.net/report/688a87d0-0314-45d8-8676-ebcb7cb63dd8

To grab the report run:
   urlquery-cli report 688a87d0-0314-45d8-8676-ebcb7cb63dd8
```

### report
Retrives a report and prints a breif summary

```
$ urlquery-cli report c3577be1-0260-45e8-87b9-e4e60292168b
Report URL: https://urlquery.net/report/c3577be1-0260-45e8-87b9-e4e60292168b
 --- 
URL:  google[.]com
IP:   142.250.74.78 (United States)
Tags: 
Alerts
  IDS:       0
  urlquery:  0
  blocklist: 0
```

### rep-check

Checking reputation on a URL
```
$ urlquery-cli rep-check enareuoasrec.com/all/login.php
{"url":"enareuoasrec.com/all/login.php","verdict":"phishing","details":null}
```


# Installation

```
go install github.com/urlquery/urlquery-cli@latest
```

# License

urlquery-cli is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/urlquery/urlquery-cli/blob/master/LICENSE.txt)

