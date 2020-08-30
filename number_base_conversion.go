// https://my.oschina.net/u/1766862/blog/1820080
package main

import(
	"fmt"
	"strconv"
)

func main() {

	var v int64 = 425217101 //默认10进制
	s2 := strconv.FormatInt(v, 2) //10 yo 16
	fmt.Printf("2  base %v\n", s2)

	s8 := strconv.FormatInt(v, 8)
	fmt.Printf("8  base %v\n", s8)

	s10 := strconv.FormatInt(v, 10)
	fmt.Printf("10 base %v\n", s10)

	s16 := strconv.FormatInt(v, 16) //10 yo 16
	fmt.Printf("16 base %v\n", s16)

	var sv = "19584c4d"; // 16 to 10
	fmt.Println(strconv.ParseInt(sv, 16, 32))
}
