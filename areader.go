// Package areader contains a custom type called A that satisfies the io.Reader
// interface. The A's Read method reads len(p) bytes into p. Each byte that is
// read represents the ASCII character 'A'.
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
