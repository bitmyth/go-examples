package main

import "fmt"

func main() {
        var a int
        var b, c = &a, &a
        fmt.Println(b, c)   // 0x1040a124 0x1040a124
        fmt.Println(&b, &c) // 0x1040c108 0x1040c110
		var d int
		b=&d
        fmt.Println(b, c)   // 0x1040a124 0x1040a124
}
