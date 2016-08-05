package endibuf

// Offset-base read methods

// ReadDataFromOffset is offset-base read method of ReadData
func (r *Reader) ReadDataFromOffset(offset int64, data interface{}) (err error) {
	origin := r.GetOffset()
	r.Seek(offset, 0)
	err = r.ReadData(data)
	r.Seek(origin, 0)
	return
}

// ReadCStringFromOffset is offset-base read method of ReadString
func (r *Reader) ReadCStringFromOffset(offset int64) (line string, err error) {
	origin := r.GetOffset()
	r.Seek(offset, 0)
	line, err = r.ReadCString()
	r.Seek(origin, 0)
	return
}

// ReadStringFromOffset is offset-base read method of ReadString
func (r *Reader) ReadStringFromOffset(offset int64, delim byte) (line string, err error) {
	origin := r.GetOffset()
	r.Seek(offset, 0)
	line, err = r.ReadString(delim)
	r.Seek(origin, 0)
	return
}

// ReadByteFromOffset is offset-base read method of ReadByte
func (r *Reader) ReadByteFromOffset(offset int64) (value byte, err error) {
	origin := r.GetOffset()
	r.Seek(offset, 0)
	value, err = r.ReadByte()
	r.Seek(origin, 0)
	return
}

// ReadBytesFromOffset is offset-base read method of ReadBytes
func (r *Reader) ReadBytesFromOffset(offset int64, size int) (value []byte, err error) {
	origin := r.GetOffset()
	r.Seek(offset, 0)
	value, err = r.ReadBytes(size)
	r.Seek(origin, 0)
	return
}

// ReadInt8FromOffset is offset-base read method of ReadInt8
func (r *Reader) ReadInt8FromOffset(offset int64) (value int8, err error) {
	origin := r.GetOffset()
	r.Seek(offset, 0)
	value, err = r.ReadInt8()
	r.Seek(origin, 0)
	return
}

// ReadInt16FromOffset is offset-base read method of ReadInt16
func (r *Reader) ReadInt16FromOffset(offset int64) (value int16, err error) {
	origin := r.GetOffset()
	r.Seek(offset, 0)
	value, err = r.ReadInt16()
	r.Seek(origin, 0)
	return
}

// ReadInt32FromOffset is offset-base read method of ReadInt32
func (r *Reader) ReadInt32FromOffset(offset int64) (value int32, err error) {
	origin := r.GetOffset()
	r.Seek(offset, 0)
	value, err = r.ReadInt32()
	r.Seek(origin, 0)
	return
}

// ReadInt64FromOffset is offset-base read method of ReadInt64
func (r *Reader) ReadInt64FromOffset(offset int64) (value int64, err error) {
	origin := r.GetOffset()
	r.Seek(offset, 0)
	value, err = r.ReadInt64()
	r.Seek(origin, 0)
	return
}

// ReadUint8FromOffset is offset-base read method of ReadUint8
func (r *Reader) ReadUint8FromOffset(offset int64) (value uint8, err error) {
	origin := r.GetOffset()
	r.Seek(offset, 0)
	value, err = r.ReadUint8()
	r.Seek(origin, 0)
	return
}

// ReadUint16FromOffset is offset-base read method of ReadUint16
func (r *Reader) ReadUint16FromOffset(offset int64) (value uint16, err error) {
	origin := r.GetOffset()
	r.Seek(offset, 0)
	value, err = r.ReadUint16()
	r.Seek(origin, 0)
	return
}

// ReadUint32FromOffset is offset-base read method of ReadUint32
func (r *Reader) ReadUint32FromOffset(offset int64) (value uint32, err error) {
	origin := r.GetOffset()
	r.Seek(offset, 0)
	value, err = r.ReadUint32()
	r.Seek(origin, 0)
	return
}

// ReadUint64FromOffset is offset-base read method of ReadUint64
func (r *Reader) ReadUint64FromOffset(offset int64) (value uint64, err error) {
	origin := r.GetOffset()
	r.Seek(offset, 0)
	value, err = r.ReadUint64()
	r.Seek(origin, 0)
	return
}

// ReadFloat32FromOffset is offset-base read method of ReadFloat32
func (r *Reader) ReadFloat32FromOffset(offset int64) (value float32, err error) {
	origin := r.GetOffset()
	r.Seek(offset, 0)
	value, err = r.ReadFloat32()
	r.Seek(origin, 0)
	return
}

// ReadFloat64FromOffset is offset-base read method of ReadFloat64
func (r *Reader) ReadFloat64FromOffset(offset int64) (value float64, err error) {
	origin := r.GetOffset()
	r.Seek(offset, 0)
	value, err = r.ReadFloat64()
	r.Seek(origin, 0)
	return
}
