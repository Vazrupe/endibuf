package endibuf

import (
	"encoding/binary"
	"io"
)

// A Reader is Simply Endian Reader
type Reader struct {
	buf    io.ReadSeeker
	Endian binary.ByteOrder
}

// NewReader returns a new Reader
func NewReader(r io.ReadSeeker) *Reader {
	return &Reader{buf: r, Endian: DefaultEndian}
}

func (r *Reader) Read(p []byte) (int, error) {
	return r.buf.Read(p)
}

// Seek return offset
func (r *Reader) Seek(offset int64, whence int) (int64, error) {
	return r.buf.Seek(offset, whence)
}

// GetOffset return this buffer offset
func (r *Reader) GetOffset() int64 {
	offset, _ := r.Seek(0, 1)
	return offset
}
