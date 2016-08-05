package bytereader

import (
	"bytes"
	"fmt"

	"github.com/vazrupe/endibuf"
)

func main() {
	b := []byte("Test Bytes Buffer\x00\x75\x32\x34\x66")
	buf := bytes.NewReader(b)
	reader := endibuf.NewReader(buf)

	line, err := reader.ReadCString()
	if err != nil {
		panic(err)
	}
	fmt.Println(line)
	numOffset := reader.GetOffset()
	fmt.Println(numOffset)
	num, err := reader.ReadUint32()
	if err != nil {
		panic(err)
	}
	fmt.Println(num)

	reader.Seek(0, 0)

	line, err = reader.ReadCStringFromOffset(0)
	if err != nil {
		panic(err)
	}
	fmt.Println(line)
	num, err = reader.ReadUint32FromOffset(numOffset)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)

	var line2 string
	var num2 uint32
	err = reader.ReadData(&line2)
	if err != nil {
		panic(err)
	}
	err = reader.ReadData(&num2)
	if err != nil {
		panic(err)
	}
	fmt.Println(line2)
	fmt.Println(num2)
}
