package main
import (
	"sort"
	"fmt"
)
//https://blog.golang.org/go-maps-in-action

func main(){
	var m map[int]string
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}
