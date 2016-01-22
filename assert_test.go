package gtester

import (
	"testing"
)

type Foo struct {
	Bar int
}

func TestEqual(t *testing.T) {
	if !Equal(`{"data":{"title":"The Litle Prince"}}`, `{"data":{"title":"The Litle Prince"}}`) {
		t.Error("can not detect two same string")
	}

	if !Equal(&Foo{1}, &Foo{1}) {
		t.Error("can not detect two same string")
	}
}
