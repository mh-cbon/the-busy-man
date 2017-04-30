package wish

import (
	"testing"
)

func TestParse(t *testing.T) {
	w, _ := Parse("y:l+k")
	want := "y"
	got := w.Plugin
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	if w.HasShade("l") == false {
		t.Error("missing shade l")
	}
	if w.HasShade("k") == false {
		t.Error("missing shade k")
	}
	if w.HasShade("hhh") {
		t.Error("wrong shade hhh")
	}
}
