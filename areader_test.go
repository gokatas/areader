package areader_test

import (
	"bytes"
	"testing"

	"github.com/gokatas/areader"
)

func TestRead(t *testing.T) {
	want := bytes.Repeat([]byte{'A'}, 10)

	var a areader.A
	got := make([]byte, 10)
	a.Read(got)

	if !bytes.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
