// https://golangbyexample.com/read-large-file-line-by-line-go/
// Read a large file Line by Line in Go (Golang)
// Posted on October 19, 2019
// When it comes to reading large files,
// obviously we don’t want to load the entire file in memory.
// bufio package in golang comes to the rescue when reading large files.

package main
import (
    "bufio"
    "fmt"
    "log"
    "os"
)
func main(){
    LinebyLineScan()
}
func LinebyLineScan() {
    file, err := os.Open("./sample.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
