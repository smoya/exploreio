// S07: package io contains many useful readers.
//
//    $ cat hello.txt | go run main.go
//    Reflection is never clear.
package main

import (
	"io"
	"os"
	"github.com/prometheus/common/log"
)

func main() {
	// TODO: Only read the first 27 bytes of stdin. 3 (or 6) lines with error handling.
	r := io.LimitReader(os.Stdin, 27)
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}
