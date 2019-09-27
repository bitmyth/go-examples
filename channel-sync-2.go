package main
import "fmt"
import "time"

func main() {
	c := make(chan bool)
	done := make(chan bool)
	for i := 0; i < 5; i++ {
		go do(c,done)
	}

	time.Sleep(time.Second)

	close(c)

	for i := 0; i < 5; i++ {
		<-done
		fmt.Println("done",i)
	}
}

func do(c <-chan bool, done chan<- bool) {
	<-c
	fmt.Println("hello")
	done<- true
}

