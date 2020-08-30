package main

import (
    . "fmt"
)

func main() {
    ch := make(chan string)
    for i := 0; i < 10; i++ {
        go printHelloWorld(i, ch) // go starts goroutine "go" 开启并发
    }

    for {
        msg := <-ch
        Println(msg)
    }
}

func printHelloWorld(i int, ch chan string) {
    // chan
    for {
        ch <- Sprintf("go run helloWorld %d \n", i)
    }
}
