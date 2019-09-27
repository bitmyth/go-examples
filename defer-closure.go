// https://www.calhoun.io/gotchas-and-common-mistakes-with-closures-in-go/
// It turns out both defer and go take a function call as an argument, not a function.

package main

import "fmt"

func main() {
  defer setup()()
}

func setup() func() {
  fmt.Println("pretend to set things up")

  return func() {
    fmt.Println("pretend to tear things down")
  }
}
