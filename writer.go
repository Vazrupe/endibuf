package endibuf

import (
	"encoding/binary"
	"io"
)

// A Writer is Simply Endian Writer
type Writer struct {
	buf    io.WriteSeeker
	Endian binary.ByteOrder
}

// NewWriter returns a new Writer
func NewWriter(w io.WriteSeeker) *Writer {
	return &Writer{buf: w, Endian: DefaultEndian}
}

func (w *Writer) Write(p []byte) (int, error) {
	return w.buf.Write(p)
}

// Seek return offset
func (w *Writer) Seek(offset int64, whence int) (int64, error) {
	return w.buf.Seek(offset, whence)
}

// GetOffset return this buffer offset
func (w *Writer) GetOffset() int64 {
	offset, _ := w.Seek(0, 1)
	return offset
}
