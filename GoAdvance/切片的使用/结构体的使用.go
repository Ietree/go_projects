package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func main() {
	// 方式一
	var man1 Person = Person{"Jack", 1, 18}
	fmt.Println(man1)

	// 方式二
	man2 := Person{
		name: "Rose",
		sex:  0,
		age:  21,
	}
	fmt.Println(man2)

	// 方式三
	var man3 Person
	man3.name = "Jay"
	man3.sex = 1
	man3.age = 40
	fmt.Println(man3)
}
