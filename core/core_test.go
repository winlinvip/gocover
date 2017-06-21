package core

import "testing"

func TestMin(t *testing.T) {
	if v := min(1, 2); v != 1 {
		t.Error("expect 1")
	}
	if v := min(2, 1); v != 1 {
		t.Error("expect 1")
	}
}