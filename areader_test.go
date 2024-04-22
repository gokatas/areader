package areader_test

import (
	"areader"
	"bytes"
	"testing"
)

func TestAEmitsThreeAs(t *testing.T) {
	want := []byte("AAA")
	got := make([]byte, 3)
	var a areader.A
	a.Read(got)
	if !bytes.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
