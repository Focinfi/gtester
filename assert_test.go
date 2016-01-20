package gtester

import (
	"testing"
)

func TestEqual(t *testing.T) {
	if !isEqual("hello", "hello") {
		t.Error("can not detect two same string")
	}
}
