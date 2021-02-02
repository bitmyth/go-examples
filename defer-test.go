package main

import(
	"fmt"
	"time"
)

func main() {
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
	}()

	for i:=0;i<10;i++{
		process()
		time.Sleep(time.Second)
	}
}

func process(){
	return
	defer fmt.Println(time.Now().Unix())
}

