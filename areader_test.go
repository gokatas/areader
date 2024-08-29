package areader_test

import (
	"bytes"
	"testing"

	"github.com/gokatas/areader"
)

func TestRead(t *testing.T) {
	var a areader.A

	buf := make([]byte, 10)

	a.Read(buf)

	want := bytes.Repeat([]byte{'A'}, len(buf))
	if !bytes.Equal(buf, want) {
		t.Errorf("unexpected data in buffer: got %v, want %v", buf, want)
	}
}
