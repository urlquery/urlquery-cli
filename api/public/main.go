package public

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var key string

func init() {
}

func SetDefaultApiKey(apikey string) {
	key = apikey
}

type apiRequest struct {
	key string
}

func NewRequest(key string) apiRequest {
	var r apiRequest
	r.key = key
	return r
}

func NewDefaultRequest() apiRequest {
	var r apiRequest
	r.key = key
	return r
}

func (a apiRequest) WithApiKey(key string) apiRequest {
	a.key = key
	return a
}

func verifyReportID(report_id string) error {
	verify := "^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"

	ok, _ := regexp.Match(verify, []byte(strings.ToUpper(report_id)))

	if ok == false {
		return fmt.Errorf("Invalid Report ID (%s)", report_id)
	}

	return nil
}

// getReader returns a appropriate reader based on the HTTP response header encoding
func getReader(resp *http.Response) io.ReadCloser {
	var reader io.ReadCloser
	reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, _ = gzip.NewReader(resp.Body)
	}
	return reader
}

// apiRequestHandle handles generic API request setup, returns JSON-data in a byte array
func apiRequestHandle(method string, url string, body io.Reader, apikey string) ([]byte, error) {
	var data []byte
	var err error
	var reader io.ReadCloser

	var netClient = &http.Client{
		Timeout: time.Second * 20,
	}

	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("User-Agent", "urlquery-cli")
	req.Header.Add("Accept-Encoding", "gzip")
	req.Close = true

	if apikey != "" {
		req.Header.Add("X-APIKEY", apikey)
	}

	resp, err := netClient.Do(req)
	if resp != nil {
		reader = getReader(resp)
		defer reader.Close()
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("Not found")
	}

	data, err = ioutil.ReadAll(reader)

	return data, err
}
