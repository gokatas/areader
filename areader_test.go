package areader_test

import (
	"testing"

	"github.com/gokatas/areader"
)

// Adapted from https://go.googlesource.com/tour/+/master/reader/validate.go
func TestRead(t *testing.T) {
	var a areader.A
	b := make([]byte, 1024, 2048)
	i, o := 0, 0
	for ; i < 1<<20 && o < 1<<20; i++ { // test 1mb
		n, err := a.Read(b)
		for i, v := range b[:n] {
			if v != 'A' {
				t.Errorf("got byte %x at offset %v, want 'A'\n", v, o+i)
			}
		}
		o += n
		if err != nil {
			t.Errorf("read error: %v\n", err)
		}
	}
	if o == 0 {
		t.Errorf("read zero bytes after %d Read calls\n", i)
	}
}
