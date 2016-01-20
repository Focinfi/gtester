package gtester

import (
	"net/http/httptest"
)

type recoder struct {
	*httptest.ResponseRecorder
}

func NewRecorder() *recoder {
	return &recoder{httptest.NewRecorder()}
}
