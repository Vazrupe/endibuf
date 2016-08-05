package endibuf

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

// ReadData is expand of 'binary.Read'
// input : valid type of binary.Read and others(*string, *float32, *flot64, []string, []float32, []float64)
func (r *Reader) ReadData(data interface{}) (err error) {
	switch data := data.(type) {
	case *string:
		*data, err = r.ReadCString()
	case []string:
		for i := range data {
			data[i], err = r.ReadCString()
			if err != nil {
				return
			}
		}
	case *float32:
		*data, err = r.ReadFloat32()
	case []float32:
		for i := range data {
			data[i], err = r.ReadFloat32()
			if err != nil {
				return
			}
		}
	case *float64:
		*data, err = r.ReadFloat64()
	case []float64:
		for i := range data {
			data[i], err = r.ReadFloat64()
			if err != nil {
				return
			}
		}
	case *int8, *int16, *int32, *int64,
		*uint8, *uint16, *uint32, *uint64,
		[]int8, []int16, []int32, []int64,
		[]uint8, []uint16, []uint32, []uint64:
		err = binary.Read(r, r.Endian, data)
	default:
		err = ErrInvalidVar
	}
	return
}

// ReadCString return string (remove last zero byte)
func (r *Reader) ReadCString() (line string, err error) {
	return r.ReadString(byte(0))
}

// ReadString return string (remove delim byte)
func (r *Reader) ReadString(delim byte) (line string, err error) {
	buffersize := 0x100
	buf := make([]byte, buffersize)
	line = ""

	var length int
	for {
		length, err = r.Read(buf)
		if err != nil {
			break
		}
		findDelim := bytes.IndexByte(buf, delim)
		found := findDelim > -1
		if found {
			noFirstByte := findDelim > 0
			if noFirstByte {
				line += string(buf[:findDelim])
			}
			delimNextPos := int64(-length + findDelim + 1)
			r.Seek(delimNextPos, 1)
			break
		} else {
			// not found delim byte
			line += string(buf)
		}
	}
	return
}

// ReadByte return one-byte
func (r *Reader) ReadByte() (value byte, err error) {
	data, err := r.ReadBytes(1)
	if err == nil {
		value = data[0]
	}
	return
}

// ReadBytes return bytes
func (r *Reader) ReadBytes(size int) (value []byte, err error) {
	value = make([]byte, size)
	_, err = r.Read(value)
	test := []interface{}{int32(0), int64(3)}
	fmt.Println(test)
	return
}

// ReadInt8 return int8
func (r *Reader) ReadInt8() (value int8, err error) {
	err = r.ReadData(&value)
	return
}

// ReadInt16 return int16
func (r *Reader) ReadInt16() (value int16, err error) {
	err = r.ReadData(&value)
	return
}

// ReadInt32 return int32
func (r *Reader) ReadInt32() (value int32, err error) {
	err = r.ReadData(&value)
	return
}

// ReadInt64 return int64
func (r *Reader) ReadInt64() (value int64, err error) {
	err = r.ReadData(&value)
	return
}

// ReadUint8 return uint8
func (r *Reader) ReadUint8() (value uint8, err error) {
	err = r.ReadData(&value)
	return
}

// ReadUint16 return uint16
func (r *Reader) ReadUint16() (value uint16, err error) {
	err = r.ReadData(&value)
	return
}

// ReadUint32 return uint32
func (r *Reader) ReadUint32() (value uint32, err error) {
	err = r.ReadData(&value)
	return
}

// ReadUint64 return uint64
func (r *Reader) ReadUint64() (value uint64, err error) {
	err = r.ReadData(&value)
	return
}

// ReadFloat32 return float32
func (r *Reader) ReadFloat32() (value float32, err error) {
	var tmp uint32
	err = binary.Read(r, r.Endian, &tmp)
	value = math.Float32frombits(tmp)
	return
}

// ReadFloat64 return float64
func (r *Reader) ReadFloat64() (value float64, err error) {
	var tmp uint64
	err = binary.Read(r, r.Endian, &tmp)
	value = math.Float64frombits(tmp)
	return
}
