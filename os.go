// [Timers](timers) are for when you want to do
// something once in the future - _tickers_ are for when
// you want to do something repeatedly at regular
// intervals. Here's an example of a ticker that ticks
// periodically until we stop it.

// https://gobyexample.com/tickers

package main

import "os"
import "fmt"

func main() {
	err :=os.MkdirAll("dir1/dir2",os.ModePerm)
	if err!=nil{
		fmt.Println(err)
	}
}

