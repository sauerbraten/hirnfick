package script

import "io"

type Cleaner struct {
	input io.Reader
}

func (c *Cleaner) Read(p []byte) (n int, err error) {
	for n < len(p) {
		buf := make([]byte, len(p)-n, len(p)-n)
		_, err = c.input.Read(buf)
		if err != nil {
			return
		}

		for _, b := range buf {
			switch b {
			case '>', '<', '+', '-', '.', ',', '[', ']':
				p[n] = b
				n++
			}
		}
	}

	return
}
