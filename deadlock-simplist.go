package main

func main() {
	c := make(chan bool)
	<-c
}


// fatal error: all goroutines are asleep - deadlock!
// goroutine 1 [chan receive]:
// main.main()
//        /Users/gsh/projects/go/learn/deadlock-simplist.go:5 +0x4d
//		exit status 2
