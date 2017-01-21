// S20: Read interface.
//
//     $ go run main.go
//     0
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Empty is an reader that returns nothing.
type Empty struct{}

func (e Empty) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

// TODO: Implement the Read interface, always return EOF. 3 lines.
// ...
// ...
// ...

func main() {
	n, err := io.Copy(os.Stdout, &Empty{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
