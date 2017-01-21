// S04: Decompressors are filters can be chained.
//
//     $ cat hello.txt.gz | go run main.go
//     Don't just check errors, handle them gracefully.
package main

import (
	"compress/gzip"
	"io"
	"os"

	"github.com/prometheus/common/log"
)

func main() {
	// TODO: Read compressed input from stdin and pass it to Stdout.
	// TODO: without using a byte slice (7 lines, including error handling).
	reader, err := gzip.NewReader(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		log.Fatal(err)
	}
}
