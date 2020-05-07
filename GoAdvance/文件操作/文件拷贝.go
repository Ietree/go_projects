package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	source_file, err := os.Open("/Users/yuanjin/Desktop/test.xyz")
	if err != nil {
		fmt.Println("open file fail!!!")
		return
	}
	defer source_file.Close()

	des_file, err := os.Create("/Users/yuanjin/Desktop/copy.txt")
	if err != nil {
		fmt.Println("创建文件失败")
		return
	}
	defer des_file.Close()

	// 重文件中读取数据，放到缓冲区中
	// 创建缓冲区
	buf := make([]byte, 4096)
	for {
		n, err := source_file.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Printf("读完：n = %d\n", n)
			return
		}
		des_file.Write(buf[:n])
	}
}
