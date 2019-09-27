// go run md5file.go -path=$HOME"/Downloads/3pp.jpg"
package main

import (
    "crypto/md5"
    "flag"
    "fmt"
    "io"
    "os"
)

var path = flag.String("path", "", "")

func main(){
	flag.Parse()
	fmt.Println( *path)
	f, err := os.Open(*path)
    if err != nil {
        fmt.Println( err.Error())
        return
    }

    defer f.Close()

    md5hash := md5.New()
    if _, err := io.Copy(md5hash, f); err != nil {
        fmt.Println("Copy", err)
        return
    }

	fmt.Printf("%x\n", md5hash.Sum(nil))
}
