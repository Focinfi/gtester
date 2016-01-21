package httpmock

import (
	"io"
	"net/http"
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

func PUT(urlStr string, body io.Reader) *Recorder {
	return defaultClient.Dail("PUT", urlStr, body)
}

func DELETE(urlStr string, body io.Reader) *Recorder {
	return defaultClient.Dail("DELETE", urlStr, body)
}
