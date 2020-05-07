package main

import "fmt"

// 单线程自己会死锁，channel应该至少2个以上的go程中通信，否则死锁
//func main() {
////	ch := make(chan int)
////
////	ch <- 789
////	num := <-ch
////	fmt.Println("num = ", num)
////}

// go程间channel访问顺序导致死锁，使用channel一端读，要保证另一端写，同时有机会执行
func main() {
	ch := make(chan int)
	num := <-ch
	fmt.Println("num = ", num)
	go func() {
		ch <- 789
	}()
}

// 多go程，多channel交叉死锁
