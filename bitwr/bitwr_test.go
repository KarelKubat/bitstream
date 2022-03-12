package bitwr_test

import (
	"strings"
	"testing"

	"github.com/KarelKubat/bitstream"
	"github.com/KarelKubat/bitstream/bitwr"
)

const (
	plain = "hello world"
	bits  = "01101000 01100101 01101100 01101100 01101111 00100000 01110111 01101111 01110010 01101100 01100100"
)

func TestBitWriter(t *testing.T) {
	// Convert bits as string to real values
	var toSend []bitstream.Bit
	for _, c := range bits {
		if c == '0' {
			toSend = append(toSend, bitstream.Zero)
		} else if c == '1' {
			toSend = append(toSend, bitstream.One)
		}
	}

	// Send bits to the actual writer.
	recv := new(strings.Builder)
	wr := bitwr.New(recv)
	for _, b := range toSend {
		if err := wr.Write(b); err != nil {
			t.Errorf("wr.Write() = %v, want nil error", err)
		}
	}

	// Check
	if recv.String() != plain {
		t.Errorf("sent bits yield %q, want %q", recv.String(), plain)
	}

}
