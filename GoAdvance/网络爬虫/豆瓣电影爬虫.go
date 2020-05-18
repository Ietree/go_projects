package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func doubanHttpGet(url string) (result string, err error) {
	respose, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		fmt.Println("http.Get err: ", err1)
		return
	}
	defer respose.Body.Close()

	buf := make([]byte, 4096)
	// 循环爬取整页数据
	for {
		n, err2 := respose.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			fmt.Println("respose.Body.Read err: ", err)
			return
		}
		result += string(buf[:n])
	}
	return
}

func save2File(idx int, filmName, filmScore, peopleNum [][]string) {
	path := "第" + strconv.Itoa(idx) + "页.txt"
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer f.Close()

	n := len(filmName)
	f.WriteString("电影名称" + "\t\t\t" + "评分" + "\t\t" + "评分人数" + "\n")
	for i := 0; i < n; i++ {
		f.WriteString(filmName[i][1] + "\t\t\t" + filmScore[i][1] + "\t\t" + peopleNum[i][1] + "\n")
	}
}

func spiderDoubanPage(idx int, page chan int) {
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((idx-1)*25) + "&filter="
	result, err := doubanHttpGet(url)
	if err != nil {
		fmt.Println("doubanHttpGet err: ", err)
		return
	}

	// 获取电影名称
	pattern1 := `<img width="100" alt="(?s:(.*?))"`
	ret1 := regexp.MustCompile(pattern1)
	filmName := ret1.FindAllStringSubmatch(result, -1)

	// 获取电影分数
	pattern2 := `<span class="rating_num" property="v:average">(?s:(.*?))</span>`
	ret2 := regexp.MustCompile(pattern2)
	filmScore := ret2.FindAllStringSubmatch(result, -1)

	// 获取电影评论人数
	pattern3 := `<span>(?s:(.*?))人评价</span>`
	ret3 := regexp.MustCompile(pattern3)
	peopleNum := ret3.FindAllStringSubmatch(result, -1)

	// 存入文件中
	save2File(idx, filmName, filmScore, peopleNum)

	// 与主go程配合，完成同步
	page <- idx
}

func toWork(start int, end int) {
	fmt.Printf("正在爬取豆瓣电影 %d 页到 %d 页的电影数据", start, end)

	// 防止主go 程提前结束
	page := make(chan int)
	for i := start; i < end; i++ {
		go spiderDoubanPage(i, page)
	}

	for i := start; i < end; i++ {
		fmt.Printf("第 %d 页爬取完毕\n", <-page)
	}
}

func main() {
	var start, end int
	fmt.Print("请输入开始查询页：")
	fmt.Scan(&start)
	fmt.Print("请输入查询结束页：")
	fmt.Scan(&end)

	toWork(start, end)
}
