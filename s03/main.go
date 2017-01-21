// S03: Stdin is a file, too.
//
// A filter that does nothing.
//
//     $ cat hello.txt | go run main.go
//     Cgo is not Go.
package main

import (
	"io"
	"os"
	"github.com/prometheus/common/log"
)

func main() {
	_, err := io.Copy(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Read input from stdin and pass it to Stdout.
	// TODO: without using a byte slice (3 lines, including error handling).
}
