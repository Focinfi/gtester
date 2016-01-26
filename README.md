## gtester

`gtester` is a box of utils for testing Golang codes.

### Install
`go get github.com/Focinfi/gtester`

### Usage
---

#### mockhttp package

`mockhttp` provide a container for a http.Handler, and a series of methods to access this Handler and return a reponse for testing.

Simple example

```go
import (
  "encoding/json"
  "net/http"
  "github.com/Focinfi/gtester/httpmock"
)

func main() {
  respJson, _ := json.Marshal(map[string]string{"hello": "world"})

  handler := func(wr http.ResponseWriter, req *http.Request) {
    wr.WriteHeader(http.StatusOK)
    fmt.Fprint(wr, string(respJson))
  }
  mux := http.NewServeMux()
  mux.HandleFunc("/hello", handler)

  // set mux as http.Handler to httpmock for later testing
  httpmock.ListenAndServe("hello.com", mux)

  // request a GET /hello/1 with a empty body
  response := httpmock.GET("/hello/1", nil)

  // we can check the response to test the mux
  response.Code == http.StatusOK // true
  response.Body.String() == string(respJson) // true
}
```

Besides `httpmock.GET()`, httpmock has other convenient methods to make requesting easier, use a `map[string]interface{}` as parameters or form values. Note that the interface{} only accepts `string` and `int`, ortherwise error will not be nil.

```go
PATCH(urlStr string, form map[string]interface{}) (*Recorder, error)
POSTForm(urlStr string, form map[string]interface{}) (*Recorder, error)
PUT(urlStr string, form map[string]interface{}) (*Recorder, error)
```

Also, if you wanna use a customized body, use the `POST()`

```go
POST(urlStr string, body io.Reader) *Recorder
DELETE(urlStr string, body io.Reader) *Recorder
```

