package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}


func (rot13reader rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13reader.r.Read(b)
	if err == nil {
		for i:=0; i<n; i++ {
			if b[i] >= 'A'&& b[i] <= 'M' || b[i] >= 'a'&& b[i] <= 'm' {
				b[i] += 13;
			} else if b[i] >= 'Z'&& b[i] >= 'N' || b[i] >= 'z'&& b[i] <= 'n' {
				b[i] -= 13;
			}
		}
	}
	return n,err
}


func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

