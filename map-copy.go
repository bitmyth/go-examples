// https://riptutorial.com/go/example/9834/copy-a-map

package main
import(
	"fmt"
)
// Like slices, maps hold references to an underlying data structure.
// So by assigning its value to another variable, only the reference will be passed.
// To copy the map, it is necessary to create another map and copy each value:

func main(){
    // Create the original map
	originalMap := make(map[string]int)
	originalMap["one"] = 1
	originalMap["two"] = 2

	// Create the target map
	targetMap := make(map[string]int)

	// Copy from the original map to the target map
	for key, value := range originalMap {
		targetMap[key] = value
	}

	fmt.Printf("originalMap %v",originalMap)
	fmt.Printf("targetMap %v",targetMap)
}
