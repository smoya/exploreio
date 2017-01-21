// S24a: A counting reader, count the number of bytes read in total.
//
//     $ cat main.go | go run main.go
//     n (io.Copy) = 550, n (CountingReader) = 550
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync/atomic"
)

type CountingReader struct {
	r     io.Reader
	count uint64
}

func (c *CountingReader) Read(p []byte) (n int, err error) {
	n, err = c.r.Read(p)
	atomic.AddUint64(&c.count, uint64(n))
	return
}

func (c CountingReader) Count() uint64 {
	return atomic.LoadUint64(&c.count)
}

func main() {
	cr := &CountingReader{r: os.Stdin}
	n, err := io.Copy(ioutil.Discard, cr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("n (io.Copy) = %d, n (CountingReader) = %d\n", n, cr.Count())
}
