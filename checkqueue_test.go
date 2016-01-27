package gtester

import (
	"fmt"
	"testing"
)

func TestCheckQueue(t *testing.T) {
	base := 0

	cq := NewCheckQueue()
	cq.Add(func() error {
		base++
		return nil
	}).Add(func() error {
		base++
		return fmt.Errorf("check2")
	}).Add(func() error {
		base++
		return fmt.Errorf("check3")
	}).Run()

	AssertEqual(t, len(cq.Checks), 3)
	AssertEqual(t, cq.Err.Error(), "check2")
	AssertEqual(t, base, 2)
}
