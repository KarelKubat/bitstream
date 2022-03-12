package bitrd

import (
	"io"

	"github.com/KarelKubat/bitstream"
)

// Reader us the receiver of a `bitrd`, holding the underlying standard `io.Reader`.
type Reader struct {
	Reader io.Reader
	buf    []byte
	bitPos int
}

// New instantiates a new `Reader`. The passed-in `io.Reader` is used a the source of bytes.
func New(r io.Reader) (bitreader *Reader) {
	return &Reader{
		Reader: r,
		buf:    []byte{0},
		bitPos: -1,
	}
}

// Read fetches the next bit from the underlying `io.Reader` and returns it as a `bitstream.Bit`,
// having the value `bitstream.Zero` or `bitstream.One`. Any error return indicates a fault of the
// underlying `io.Reader` which is responsible for fetching separate bytes.
func (r *Reader) Read() (bit bitstream.Bit, err error) {
	if r.bitPos < 0 {
		_, err = r.Reader.Read(r.buf)
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
