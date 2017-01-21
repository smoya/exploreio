// S05: An command line image converter.
//
// Read PNG from standard input, write JPG to standard output.
//
// Inspired by Donovan, Kernighan (2016), page 287.
// https://creativecommons.org/licenses/by-nc-sa/4.0/
//
//    $ cat gopherbw.png | go run main.go > gopherbw.jpg
//    $ ls -lah gopherbw.jpg
//    ...
//    ... 542541 Jan 20 17:08 gopherbw.jpg
//
package main

import (
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"
)

// toJPG converts gif or PNG to JPG.
func toJPG(in io.Reader, out io.Writer) error {
	// TODO: Read the image, encode the image. 5 lines with error handling.
	img, _, err := image.Decode(in)
	if err != nil {
		log.Fatal(err)
	}

	return jpeg.Encode(out, img, nil)
}

func main() {
	if err := toJPG(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
