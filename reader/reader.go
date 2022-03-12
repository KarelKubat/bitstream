package bitrd

import (
	"io"

	"github.com/KarelKubat/bitstream"
)

type Reader struct {
	Reader io.Reader
	buf    []byte
	bitPos int
}

func New(r io.Reader) *Reader {
	return &Reader{
		Reader: r,
		buf:    []byte{0},
		bitPos: -1,
	}
}

func (r *Reader) Read() (bit bitstream.Bit, err error) {
	if r.bitPos < 0 {
		var n int
		n, err = r.Reader.Read(r.buf)
		if err != nil {
			return bitstream.Zero, err
		}
		r.bitPos = 7
	}
	mask := 1 << r.bitPos
	if r.buf[0]&byte(mask) == 0 {
		bit = bitstream.Zero
	} else {
		bit = bitstream.One
	}
	r.bitPos--
	return bit, nil
}
