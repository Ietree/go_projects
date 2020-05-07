package main

import "fmt"

func main() {
	var a int = 10

	var p *int = &a

	a = 100
	fmt.Println("a = ", a) // 100

	*p = 250
	fmt.Println("a = ", a)   // 100
	fmt.Println("*p = ", *p) // 250

	a = *p
	fmt.Println("a = ", a) // 250
}
