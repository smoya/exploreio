// S23: An uppercase writer.
//
//     $ echo "Hello Gophers" | go run main.go
//     HELLO GOPHERS
package main

import (
	"io"
	"log"
	"os"
)

// TODO: Implement UpperWriter, a reader that converts all Unicode letter mapped to their upper case. 6 lines.
// ...
// ...
// ...

// ...
// ...
// ...

func main() {
	if _, err := io.Copy(&UpperWriter{os.Stdout}, os.Stdin); err != nil {
		log.Fatal(err)
	}
}
