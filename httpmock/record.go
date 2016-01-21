package httpmock

import (
	"net/http/httptest"
)

type Recoder struct {
	*httptest.ResponseRecorder
}

func NewRecorder() *Recoder {
	return &Recoder{httptest.NewRecorder()}
}
