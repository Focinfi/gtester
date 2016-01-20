package gtester

import (
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

func (client *httpMockClient) get(urlStr string, response *recoder) {
	request, _ := http.NewRequest("GET", urlStr, nil)
	client.handler.ServeHTTP(response, request)
}

var defaultClient = &httpMockClient{}

func ListenAndServe(host string, handler http.Handler) {
	defaultClient.listenAndServe(host, handler)
}

func Get(urlStr string, response *recoder) {
	defaultClient.get(urlStr, response)
}
