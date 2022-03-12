package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/KarelKubat/bitstream/bitrd"
)

func main() {
	r := bitrd.New(os.Stdin)

	nread := 0
	for {
		bit, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("read failed: %v\n", err)
		}
		nread++
		fmt.Print(bit)
		if nread%8 == 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
