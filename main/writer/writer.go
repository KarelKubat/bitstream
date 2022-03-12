package main

import (
	"io"
	"log"
	"os"

	"github.com/KarelKubat/bitstream"
	"github.com/KarelKubat/bitstream/bitwr"
)

func main() {
	w := bitwr.New(os.Stdout)
	for {
		buf := []byte{0}
		_, err := os.Stdin.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		if buf[0] == '0' {
			w.Write(bitstream.Zero)
		} else if buf[0] == '1' {
			w.Write(bitstream.One)
		}
	}
}
