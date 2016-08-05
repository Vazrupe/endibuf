package endibuf

import (
	"encoding/binary"
	"errors"
)

// ErrInvalidVar is '*Reader.ReadData' function invalid type errors
var ErrInvalidVar = errors.New("invalid variable")

// DefaultEndian is default byte-order of NewReader
var DefaultEndian = binary.BigEndian
