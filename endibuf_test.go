package endibuf

import (
	"encoding/binary"
	"io/ioutil"
	"math"
	"os"
	"reflect"
	"testing"
)

var cases = []struct {
	data   []interface{}
	endian binary.ByteOrder
	middle []byte
}{
	{ // BigEndian Case
		[]interface{}{
			int8(0x01), int16(0x3224), int32(0x20723630), int64(0x3472568424695484),
			uint8(0x70), uint16(0x8089), uint32(0x63434234), uint64(0x4785173264736748),
			string("test"),
		},
		binary.BigEndian,
		[]byte("\x01" + "\x32\x24" + "\x20\x72\x36\x30" + "\x34\x72\x56\x84\x24\x69\x54\x84" +
			"\x70" + "\x80\x89" + "\x63\x43\x42\x34" + "\x47\x85\x17\x32\x64\x73\x67\x48" +
			"test\x00"),
	},
	{ // LittleEndian Case
		[]interface{}{
			int8(0x48), int16(0x4726), int32(0x12247686), int64(0x1253255435453463),
			uint8(0x84), uint16(0x8089), uint32(0x57257562), uint64(0x5463435153414244),
			string("test"),
		},
		binary.LittleEndian,
		[]byte("\x48" + "\x26\x47" + "\x86\x76\x24\x12" + "\x63\x34\x45\x35\x54\x25\x53\x12" +
			"\x84" + "\x89\x80" + "\x62\x75\x25\x57" + "\x44\x42\x41\x53\x51\x43\x63\x54" +
			"test\x00"),
	},
	{ // Minus Case
		[]interface{}{
			int8(-0x48), int16(-0x4726), int32(-0x12247686), int64(-0x1253255435453463),
		},
		binary.BigEndian,
		[]byte("\xB8" + "\xB8\xDA" + "\xED\xDB\x89\x7A" + "\xED\xAC\xDA\xAB\xCA\xBA\xCB\x9D"),
	},
	{ // Float Case
		[]interface{}{
			float32(math.Pi), math.Pi,
		},
		binary.BigEndian,
		[]byte("\x40\x49\x0f\xdb" + "\x40\x09\x21\xfb\x54\x44\x2d\x18"),
	},
}

// TestEndianBuffer is Reader/Writer test with Tempfiles
func TestEndianBuffer(t *testing.T) {
	for cIdx, c := range cases {
		types := make([]interface{}, len(c.data))

		tempfile, err := ioutil.TempFile("", "endibuf_")
		if err != nil {
			t.Errorf("Tempfile Create Error case: %d / %s", cIdx, err)
		}
		w := NewWriter(tempfile)
		w.Endian = c.endian
		for i := range c.data {
			types[i] = c.data[i]
			w.WriteData(c.data[i])
		}
		err = tempfile.Close()
		if err != nil {
			t.Errorf("Tempfile Close Error case: %d / %s", cIdx, err)
		}

		readbuf, err := os.Open(tempfile.Name())
		if err != nil {
			t.Errorf("Tempfile Read Error case: %d / %s", cIdx, err)
		}
		defer readbuf.Close()

		b, err := ioutil.ReadAll(readbuf)
		if !reflect.DeepEqual(b, c.middle) {
			t.Errorf("Not Matched Middle case: %d / %s", cIdx, err)
		}
		readbuf.Seek(0, 0)
		r := NewReader(readbuf)
		r.Endian = c.endian
		for i := range types {
			switch data := types[i].(type) {
			default:
				err = r.ReadData(&data)
				if err != nil {
					t.Errorf("Read Error case: %d / %s", cIdx, err)
				}
				if !reflect.DeepEqual(data, c.data[i]) {
					t.Errorf("Read Data Not Matched case: %d / %s", cIdx, err)
				}
			}
		}

		defer os.Remove(tempfile.Name())
	}
}
