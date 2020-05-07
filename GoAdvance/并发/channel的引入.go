package main

import (
	"fmt"
	"time"
)

// 全局定义channel，用来完成数据同步
var channel = make(chan int)

func printer(s string) {
	for _, ch := range s {
		fmt.Printf( "%c", ch)
		time.Sleep(1000 * time.Millisecond)
	}
}

func person1() {
	printer("Hello, How are you")
	channel <- 8
}

func person2() {
	<-channel
	printer("I'm fine, think you")
}

func main() {
	go person1()
	go person2()
	select {}
}
