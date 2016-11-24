package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type FirstWordReader struct {
	R io.Reader
}

// START OMIT

func (r FirstWordReader) Read(p []byte) (n int, err error) {
	v, err := ioutil.ReadAll(r.R)
	if err != nil {
		return 0, err
	}
	n = len(v)
	for i, c := range v {
		if c == ' ' {
			n = i
			break
		}
	}
	copy(p, v[:n])
	return n, io.EOF
}

func main() {
	//f, _ := os.Open("./read.txt")
	//fwr := FirstWordReader{f}
	fwr := FirstWordReader{strings.NewReader("Gophers in Avocode!")}
	v, _ := ioutil.ReadAll(fwr)
	fmt.Printf("FirstWordReader returned %q", string(v))
}

// END OMIT
