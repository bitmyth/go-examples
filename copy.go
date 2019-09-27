package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3}}
	// tmp := make([][]int, len(arr))
	tmp := [][]int{}

	// copy(tmp, arr)
	tmp= arr
	fmt.Println(len(tmp))
	fmt.Println(tmp)
	fmt.Println(arr)
	tmp = [][]int{}
	fmt.Println(arr)
	fmt.Println(tmp)
}

