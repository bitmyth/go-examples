package main
import (
	"github.com/euclidr/darts"
	"fmt"
	"sort"
)


func main(){
	builder := darts.DoubleArrayBuilder{}
	keyset := []string{"印度", "印度尼西亚", "印加帝国", "瑞士", "瑞典", "巴基斯坦", "巴勒斯坦", "以色列", "巴比伦", "土耳其"}
	sort.Strings(keyset)

	// Build darts
	builder.Build(keyset)

	// ExactMatchSearch
	key := "印度"
	result, matched := builder.ExactMatchSearch(key)
	if !matched {
		fmt.Errorf("invalid result, not matched: %s", key)
		return
	}
	fmt.Println(keyset[result])

	// CommonPrefixSearch
	values := builder.CommonPrefixSearch("印度尼西亚啊")
	fmt.Printf("%s, %s", keyset[values[0]], keyset[values[1]])
}
