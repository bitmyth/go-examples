package main

import (
    "crypto/sha256"
    "fmt"
)

func main(){
    s := "sha256 芳华"

    h := sha256.New()
    h.Write([]byte(s))
    bs := h.Sum(nil)

    fmt.Printf("origin: %s, sha256 hash: %x\n", s, bs)
}
