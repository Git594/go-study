package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (n int, err error) {
	for {
		count, errors := rot.r.Read(b)
		rot13(b, count)
		if errors == io.EOF {
			break
		}
		n += count
	}
	return n, err
}

func rot13(b []byte, n int) {
	for i := 0; i < n; i++ {
		v := b[i]
		if v >= 'a' && v <= 'm' || v >= 'A' && v <= 'M' {
			b[i] = v + 13
		} else if v >= 'n' && v <= 'z' || v >= 'N' && v <= 'Z' {
			b[i] = v - 13
		}
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
