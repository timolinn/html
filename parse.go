package html

import "io"

func Parse(r io.Reader) []byte {
	return []byte(`
        h1
            Hello Parser
            `)
}
