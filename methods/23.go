package main

import (
	"io"
	"os"
	// "fmt"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read (b []byte) (int, error) {
	n, err := rot13.r.Read(b)

	var (
		from = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
		to = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	)

	for i := 0; i < n; i++ {
		if index := strings.Index(from, string(b[i])); index != -1 {
			b[i] = to[index]
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}