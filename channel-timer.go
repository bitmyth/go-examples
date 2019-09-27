package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 2)

	go func() {
		time.Sleep(10 * time.Second)
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			}
		}
	}()

	ticker := time.NewTicker(1 * time.Second)

	j := 0
	for t := range ticker.C {
		fmt.Println("Tick at", t)
		j++
		jobs <- j
		fmt.Println("sent job", j)
	}

	fmt.Println("sent all jobs")
}
