// https://github.com/jemygraw/TechDoc/blob/master/Go%E7%A4%BA%E4%BE%8B%E5%AD%A6/Go%20String%E4%B8%8EByte%E5%88%87%E7%89%87%E4%B9%8B%E9%97%B4%E7%9A%84%E8%BD%AC%E6%8D%A2.markdown
package main

import "fmt"
import "strconv"

func main() {

    s1 := "abcd"
    b1 := []byte(s1)
    fmt.Println(b1) // [97 98 99 100]

    s2 := "中文"
    b2 := []byte(s2)
    fmt.Println(b2)  // [228 184 173 230 150 135], unicode，每个中文字符会由三个byte组成

    r1 := []rune(s1)
    fmt.Println(r1) // [97 98 99 100], 每个字一个数值

    r2 := []rune(s2)
    fmt.Println(r2) // [20013 25991], 每个字一个数值

	//  https://stackoverflow.com/questions/37210379/convert-int-to-a-single-byte-in-go/37210523
	val := "7"
	i, _ := strconv.Atoi(val)
	byteI := byte(i)
	fmt.Printf("%v (%T)", byteI, byteI)

}
