package main
import "fmt"

var complete chan int = make(chan int)

func loop(start int) {
	for i := start; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	complete <- 0 // 执行完毕了，发个消息
}


func main() {
	//loop(0)
	go loop(1)
	<- complete // 直到线程跑完, 取到消息. main在此阻塞住
}
