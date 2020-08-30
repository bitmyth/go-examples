package main

import (
	"fmt"
	"time"
)

func sleep(){
	fmt.Println("sleep 30s")
	time.Sleep(30000 * time.Millisecond)
	fmt.Println("end sleep 30 second")
}

func say(s string) {
	if(s=="w"){
		fmt.Println("sleep 3s")
		time.Sleep(3000 * time.Millisecond)
		fmt.Println("sleep 3s")
	}
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go sleep()
	go say("world")
	go say("w")
	say("hello")
}

