package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
)

func main() {
    FileName := "assets/file.txt"
    file, err := os.Open(FileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() { 
        fmt.Println(scanner.Text()) 

    }
}
