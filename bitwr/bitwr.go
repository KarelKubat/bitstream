package bitwr

import (
	"io"

	"github.com/KarelKubat/bitstream"
)

type Writer struct {
	Writer io.Writer
	buf    []byte
	bitPos int
}

func New(w io.Writer) (bitwriter *Writer) {
	return &Writer{
		Writer: w,
		buf:    []byte{0},
		bitPos: 7,
	}
}

func (w *Writer) Write(bit bitstream.Bit) (err error) {
	if _, err := w.Flush(); err != nil {
		return err
	}
	w.buf[0] |= byte(bit) << w.bitPos
	w.bitPos--
	w.Flush()
	return nil
}

func (w *Writer) Flush() (flushed bool, err error) {
	if w.bitPos >= 0 {
		return false, nil
	}
	if _, err := w.Writer.Write(w.buf); err != nil {
		return false, err
	}
	w.bitPos = 7
	w.buf[0] = 0
	return true, nil
}
