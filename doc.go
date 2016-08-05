package endibuf

/*
endibuf
==========

Description
-----------
Big/Little endian based binary reader/writer module of Go

Usage
-----
Installation:

    go get github.com/vazrupe/endibuf

Import:

    import (
        ...
        "github.com/vazrupe/endibuf"
        ...
    )

Reader:

    reader := endibuf.NewReader(_buffer_)
    var data _type_
    err := reader.ReadData(&data)
    if err != nil {
        panic(err)
    }
    fmt.Println(data)
    // define type
    data, err = reader.Read_TYPENAME_()
    if err != nil {
        panic(err)
    }
    fmt.Println(data)
    // define type from offset
    data, err = reader.Read_TYPENAME_FromOffset(_offset_)
    if err != nil {
        panic(err)
    }
    fmt.Println(data)

Writer:

    writer := endibuf.NewWriter(_buffer_)
    data := _your code_
    err := writer.WriteData(data)
    if err != nil {
        panic(err)
    }
    // define type
    err = writer.Write_TYPENAME_(data)
    if err != nil {
        panic(err)
    }
    // define type from offset
    err = writer.Write_TYPENAME_FromOffset(_offset_, data)
    if err != nil {
        panic(err)
    }

License
-------
MIT Lisence.

Author
------
Hyeon-Gyu Lee (a.k.a Vazrupe)
*/
