package protocol

import "testing"

func TestMax(t *testing.T) {
	if v := max(1, 2); v != 2 {
		t.Error("expect 2")
	}
	if v := max(2, 1); v != 2 {
		t.Error("expect 2")
	}
}