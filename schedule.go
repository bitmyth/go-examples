package main

import (
	"time"
)

func main() {
	done := make(chan bool)
	Call(500, func(t time.Time) {
		println("Tick at", t.String())
	})
	<-done
}

func Call(interval int, task func(time time.Time)) {
	ticker := time.NewTicker(time.Duration(interval) * time.Millisecond)

	go func() {
		for t := range ticker.C {
			task(t)
		}
	}()

}
