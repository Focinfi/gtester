package httpmock

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type httpMockClient struct {
	host    string
	handler http.Handler
}

func (client *httpMockClient) listenAndServe(host string, handler http.Handler) {
	client.host = host
	client.handler = handler
}

func (client *httpMockClient) Dail(mehtod string, urlStr string, body io.Reader) *Recorder {
	request, _ := http.NewRequest(mehtod, urlStr, body)
	response := NewRecorder()
	client.handler.ServeHTTP(response, request)
	return response
}

func (client *httpMockClient) DailWithForm(method string, urlStr string, form map[string]interface{}) (*Recorder, error) {
	params := url.Values{}
	for key := range form {
		switch value := form[key].(type) {
		case string:
			params.Set(key, value)
		case int:
			params.Set(key, strconv.Itoa(value))
		default:
			return nil, fmt.Errorf("has unexcepted type of value")
		}
	}

	body := ioutil.NopCloser(strings.NewReader(params.Encode()))
	request, _ := http.NewRequest("POST", urlStr, body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	response := NewRecorder()
	client.handler.ServeHTTP(response, request)
	return response, nil
}

var defaultClient = &httpMockClient{}

func ListenAndServe(host string, handler http.Handler) {
	defaultClient.listenAndServe(host, handler)
}

func GET(urlStr string, body io.Reader) *Recorder {
	return defaultClient.Dail("GET", urlStr, body)
}

func POST(urlStr string, body io.Reader) *Recorder {
	return defaultClient.Dail("POST", urlStr, body)
}

func POSTForm(urlStr string, form map[string]interface{}) (*Recorder, error) {
	return defaultClient.DailWithForm("POST", urlStr, form)
}

func PUT(urlStr string, form map[string]interface{}) (*Recorder, error) {
	return defaultClient.DailWithForm("PUT", urlStr, form)
}

func PATCH(urlStr string, form map[string]interface{}) (*Recorder, error) {
	return defaultClient.DailWithForm("PATCH", urlStr, form)
}

func DELETE(urlStr string, body io.Reader) *Recorder {
	return defaultClient.Dail("DELETE", urlStr, body)
}
