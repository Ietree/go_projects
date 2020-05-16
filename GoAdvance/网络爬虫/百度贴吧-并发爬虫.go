package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func httpGet1(url string) (result string, err error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get err: ", err)
		return
	}
	defer response.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, err2 := response.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取网页数据完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}

func spider(i int, page chan int) {
	url := "https://tieba.baidu.com/f?kw=%E6%97%A0%E4%BA%BA%E6%9C%BA&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	result, err := httpGet1(url)
	if err != nil {
		fmt.Println("http.Get err: ", err)
		return
	}

	f, err := os.Create("爬取到" + strconv.Itoa(i) + "页数据.html")
	if err != nil {
		fmt.Println("os.Create err: ", err)
		return
	}

	f.WriteString(result)
	f.Close()

	page <- i
}

func working(start int, end int) {
	fmt.Printf("正在爬取第%d页到%d页...\n", start, end)

	page := make(chan int)
	for i := start; i < end; i++ {
		go spider(i, page)
	}

	for i := start; i < end; i++ {
		fmt.Printf("第 %d 页数据爬取完成", <-page)
	}
}

func main() {
	var start, end int
	fmt.Print("请输入查询起始页：")
	fmt.Scan(&start)
	fmt.Print("请输入查询结束页：")
	fmt.Scan(&end)
	// 爬取页面信息
	working(start, end)
}
