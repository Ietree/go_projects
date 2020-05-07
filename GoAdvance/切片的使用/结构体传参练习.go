package main

import "fmt"

type People struct {
	name      string
	age       int
	flag      bool
	intereset []string
}

func initFunc(p *People) {
	p.name = "Jack"
	p.age = 30
	p.flag = true
	p.intereset = append(p.intereset, "武术")
	p.intereset = append(p.intereset, "电影")
	p.intereset = append(p.intereset, "慈善")
}

func main() {
	var p People
	initFunc(&p)
	fmt.Println(p)
}
