package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func readFile(s string) int {
	f, err := os.OpenFile(s, os.O_RDONLY, 6)
	if err != nil {
		fmt.Println("Open err：", err)
		return -1
	}
	defer f.Close()

	row := bufio.NewReader(f)
	total := 0

	for {
		buf, err := row.ReadBytes('\n')
		if err != nil && err == io.EOF {
			total += wordCount(string(buf[:]))
			break
		}
		total += wordCount(string(buf[:]))
	}
	return total
}

// 从一行字符串中获取love出现的次数，将该次数返回
func wordCount(s string) int {
	arr := strings.Fields(s)
	m := make(map[string]int)

	// 对arr中的每个单词进行循环，存入Map中
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}

	// 统计Map中一共有多少个"Love"
	for k, v := range m {
		if k == "love" {
			fmt.Printf("%s : %d\n", k, v)
			return m[k]
		}
	}
	return 0
}

func main() {
	fmt.Println("请输入要找寻的目录：")
	var path string
	fmt.Scan(&path)

	// 只读模式打开指定目录
	dir, _ := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	defer dir.Close()
	// 读取当前目录下所有文件名和目录名，存入names切片
	names, _ := dir.Readdir(-1)
	// 遍历切片，获取文件名/目录名
	allLove := 0
	for _, name := range names {
		if !name.IsDir() {
			// 文件名不包括路径
			s := name.Name()
			if strings.HasSuffix(s, ".txt") {
				// 拼接路径和文件名
				allLove += readFile(path + s)
			}
		}
	}
	fmt.Printf("目录所有文件中共有%d个Love\n", allLove)
}
