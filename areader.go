// Package areader implements io.Reader interface with the A type. The A's Read
// method reads up to len(p) bytes representing the ASCII character 'A' into p.
//
// Adapted from https://tour.golang.org/methods/22.
package areader

type A struct{}

func (A) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}
