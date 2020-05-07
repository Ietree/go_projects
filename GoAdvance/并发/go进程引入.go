package main

import (
	"fmt"
	"time"
)

func sing() {
	for i := 0; i < 6; i++ {
		fmt.Println("-----------唱歌-------------")
		time.Sleep(4 * time.Millisecond)
	}

}

func dance() {
	for i := 0; i < 6; i++ {
		fmt.Println("-----------跳舞-------------")
		time.Sleep(4 * time.Millisecond)
	}
}

func main() {
	go sing()
	go dance()
	for {

	}
}
