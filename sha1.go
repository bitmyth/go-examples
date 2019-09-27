package main

import(
	"crypto/sha1"
	"math/rand"
	"fmt"
)

func main(){
	v :=rand.Int63()
	println(v)
	fmt.Printf("%v",sha1.Sum([]byte("abc")))
}
