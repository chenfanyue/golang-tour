package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M':
		b += 13
	case 'M' < b && b <= 'Z':
		b -= 13
	case 'a' <= b && b <= 'm':
		b += 13
	case 'm' < b && b <= 'z':
		b -= 13
	}
	return b
}

func (p *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = p.r.Read(b)

	for i, v := range b {
		b[i] = rot13(v)
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
