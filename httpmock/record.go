package httpmock

import (
	"encoding/json"
	"net/http/httptest"
)

type Recorder struct {
	*httptest.ResponseRecorder
}

func NewRecorder() *Recorder {
	return &Recorder{httptest.NewRecorder()}
}

func (r Recorder) JSON() interface{} {
	respStr := r.Body.String()
	var respJson interface{}
	if err := json.Unmarshal([]byte(respStr), &respJson); err != nil {
		return nil
	} else {
		return respJson
	}
}
