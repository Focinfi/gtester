package httpmock

import (
	"net/http/httptest"
)

type Recoder struct {
	*httptest.ResponseRecorder
}

func NewRecorder() *recoder {
	return &recoder{httptest.NewRecorder()}
}
