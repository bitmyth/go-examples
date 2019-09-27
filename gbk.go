// http://mengqi.info/html/2015/201507071345-using-golang-to-convert-text-between-gbk-and-utf-8.html
package main

import (
    "bytes"
    "golang.org/x/text/encoding/simplifiedchinese"
    "golang.org/x/text/transform"
    "io/ioutil"
    "fmt"
)

func GbkToUtf8(s []byte) ([]byte, error) {
    reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
    d, e := ioutil.ReadAll(reader)
    if e != nil {
        return nil, e
    }
    return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
    reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
    d, e := ioutil.ReadAll(reader)
    if e != nil {
        return nil, e
    }
    return d, nil
}

func main() {

    s := "GBK 与 UTF-8 编码转换测试"
    gbk, err := Utf8ToGbk([]byte(s))
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(string(gbk))
    }

    utf8, err := GbkToUtf8(gbk)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(string(utf8))
    }
}
