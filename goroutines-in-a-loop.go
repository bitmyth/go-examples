// https://riptutorial.com/go/example/3897/using-closures-with-goroutines-in-a-loop

package main

import(
	"fmt"
	"os"
	"syscall"
	"os/signal"
)

func main(){
	values:=[]int{1,2,3,4,5}

	// When in a loop, the loop variable (val) in the following example is a single variable that changes value as it goes over the loop.
	// Therefore one must do the following to actually pass each val of values to the goroutine:
	//for val := range values {
	//	go func(val interface{}) {
	//		fmt.Println(val)
	//	}(val)
	//}

	// If you were to do just do go func(val interface{}) { ... }() without passing val,
	// then the value of val will be whatever val is when the goroutines actually runs.
	//for val := range values {
    //	go func() {
	//		fmt.Println(val)
	//	}()
	//}

	//// Another way to get the same effect is:
	for val := range values {
		val :=val
		go func() {
			fmt.Println(val)
		}()
	}

	sigs :=make(chan os.Signal,1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
