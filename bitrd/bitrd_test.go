package bitrd_test

import (
	"io"
	"strings"
	"testing"

	"github.com/KarelKubat/bitstream"
	"github.com/KarelKubat/bitstream/bitrd"
)

const (
	plain = "hello world"
	bits  = "01101000 01100101 01101100 01101100 01101111 00100000 01110111 01101111 01110010 01101100 01100100"
)

func TestBitreader(t *testing.T) {
	// Convert bits as string to real values
	var wantBits []bitstream.Bit
	for _, c := range bits {
		if c == '0' {
			wantBits = append(wantBits, bitstream.Zero)
		} else if c == '1' {
			wantBits = append(wantBits, bitstream.One)
		}
	}

	// Collect bits from the actual reader.
	rd := bitrd.New(strings.NewReader(plain))
	pos := 0
	for {
		b, err := rd.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("bitrd.Read() = _,%v, unexpected error", err)
		}
		if b != wantBits[pos] {
			t.Errorf("bitrd.Read() at pos %v = %v, want %v", pos, b, wantBits[pos])
		}
		pos++
	}
	if pos != len(wantBits) {
		t.Errorf("bitrd.Read()s yielded %v values, want %v", pos, len(wantBits))
	}
}
