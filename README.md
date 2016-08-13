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

    reader := endibuf.NewReader(YOUR_BUFFER)
    var data TYPE
    err := reader.ReadData(&data)
    if err != nil {
        panic(err)
    }
    fmt.Println(data)
    // define type
    data, err = reader.ReadTYPENAME()
    if err != nil {
        panic(err)
    }
    fmt.Println(data)
    // define type from offset
    data, err = reader.ReadTYPENAMEFromOffset(_offset_)
    if err != nil {
        panic(err)
    }
    fmt.Println(data)

Writer:

    writer := endibuf.NewWriter(YOUR_BUFFER)
    data := YOUR_CODE
    err := writer.WriteData(data)
    if err != nil {
        panic(err)
    }
    // define type
    err = writer.WriteTYPENAME(data)
    if err != nil {
        panic(err)
    }
    // define type from offset
    err = writer.WriteTYPENAMEFromOffset(_offset_, data)
    if err != nil {
        panic(err)
    }

License
------- 
MIT Lisence.

Author
------
Hyeon-Gyu Lee (a.k.a Vazrupe)
