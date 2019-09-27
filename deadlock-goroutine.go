package main
import "fmt"
import "time"

func main() {
	c := make(chan bool)
	for i := 0; i < 5; i++ {
		c <- true
		go do(c)
	}

	time.Sleep(time.Second)

	c <- true
}

func do(c <-chan bool) {
	<-c
	fmt.Println("hello")
}
