// https://golang.org/ref/spec#Receive_operator
package main
import(
	"fmt"
)
func main(){
	jobs:=make(chan struct{})
	close(jobs)
	// a receive operation on a closed channel can always proceed immediately, yielding the element type's zero value.
	fmt.Printf("%v",<-jobs)

	intChan:=make(chan int)
	close(intChan)
	fmt.Printf("%v",<-intChan)
}
