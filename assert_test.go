package gtester

import (
	"testing"
)

func TestEqual(t *testing.T) {
	if !Equal(`{"data":{"title":"The Litle Prince"}}`, `{"data":{"title":"The Litle Prince"}}`) {
		t.Error("can not detect two same string")
	}
}
