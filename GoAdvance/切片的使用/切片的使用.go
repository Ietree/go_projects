package main

import "fmt"

/**
切片的使用：
	切片名称[low:high:max]
	low:起始下标位置
	high:结束下标位置 len = high - low
	cap:容量 cap = max - low

创建切片的常用方式：
	1.自动推导类型创建切片：slice := []int {1, 2, 3, 4}
	2.slice := make([]int, 长度， 容量)
	3.【常用】slice := make([]int, 长度)，创建切片时，没有指定容量，容量=长度
*/
func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}

	s := arr[1:3:5]
	fmt.Println("s = ", s)
	fmt.Println("len = ", len(s))
	fmt.Println("cap = ", cap(s))

	// 创建切片
	s1 := []int {5,4,3,2,1}
	fmt.Println("len = ", len(s1), "cap = ", cap(s1))

	s2 := make([]int, 5, 10)
	fmt.Println("len = ", len(s2), "cap = ", cap(s2))

	s3 := make([]int, 5)
	fmt.Println("len = ", len(s3), "cap = ", cap(s3))
}
