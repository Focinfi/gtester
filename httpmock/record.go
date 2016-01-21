package httpmock

import (
	"net/http/httptest"
)

type Recorder struct {
	*httptest.ResponseRecorder
}

func NewRecorder() *Recorder {
	return &Recorder{httptest.NewRecorder()}
}
