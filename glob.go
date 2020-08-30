package main
import (
	"path/filepath"
	"fmt"
)

func main() {
	m,_ := filepath.Glob("/usr/*")
	fmt.Println(m)
	for _,path :=range(m){
		fmt.Println(path)
	}
}
