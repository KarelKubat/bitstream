package bitstream_test

import (
	"io"
	"strings"
	"testing"

	"github.com/KarelKubat/bitstream"
	"github.com/KarelKubat/bitstream/bitrd"
	"github.com/KarelKubat/bitstream/bitwr"
)

const lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation
ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit
in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat
non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

func TestReadWrite(t *testing.T) {
	// Convert lorem to 0s and 1s
	bitRd := bitrd.New(strings.NewReader(lorem))
	var bits []bitstream.Bit
	for {
		bit, err := bitRd.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatal(err)
		}
		bits = append(bits, bit)
	}

	// Convert bits back to a string.
	receiver := new(strings.Builder)
	bitWr := bitwr.New(receiver)
	for _, b := range bits {
		if err := bitWr.Write(b); err != nil {
			t.Fatal(err)
		}
	}
	bitWr.Flush()
	got := receiver.String()

	// Compare.
	if got != lorem {
		t.Errorf("pipe test: got\n  %q\nbut want\n  %q\n", got, lorem)
	}
}
