package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//func main() {
//	f, err := os.Create("/Users/yuanjin/Desktop/test.xyz")
//	if err != nil {
//		fmt.Println("create err", err)
//		return
//	}
//	defer f.Close()
//	fmt.Println("Successful")
//}

func main() {
	//f, err := os.Open("/Users/yuanjin/Desktop/test.xyz")
	////打开文件
	//if err != nil {
	//	fmt.Println("open fail:", err)
	//	return
	//}
	//defer f.Close()
	//
	//// 写入数据
	//_, err = f.WriteString("################")
	//if err != nil {
	//	fmt.Println("write fail:", err)
	//	return
	//}

	//// 向文件中写入字符串
	//f, err := os.Create("/Users/yuanjin/Desktop/test.xyz")
	//f, err = os.OpenFile("/Users/yuanjin/Desktop/test.xyz", os.O_RDWR, 6)
	//if err != nil {
	//	fmt.Println("open file fail:", err)
	//	return
	//}
	//defer f.Close()
	//
	//_, err = f.WriteString("###############???????????????@@@@@@@@@@@@@@@@@@@@")
	//if err != nil {
	//	fmt.Println("write file fail:", err)
	//	return
	//}
	//
	//fmt.Println("Successful")

	//// 向文件中写入字节
	//f, err := os.OpenFile("/Users/yuanjin/Desktop/test.xyz", os.O_RDWR, 6)
	//if err != nil {
	//	fmt.Println("open file fail:", err)
	//	return
	//}
	//defer f.Close()
	//
	//off, _ := f.Seek(-5, io.SeekEnd)
	//fmt.Println("off:", off)
	//
	//n, _ := f.WriteAt([]byte("11111111111"), off)
	//fmt.Println("WriteAt", n)

	// 向文件中读取字符串
	f, err := os.OpenFile("/Users/yuanjin/Desktop/test.xyz", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("open file fail:", err)
		return
	}
	defer f.Close()

	// 创建一个带有缓冲区的reader
	reader := bufio.NewReader(f)

	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil && err == io.EOF {
			fmt.Println("文件读取完毕")
			return
		} else if err != nil {
			fmt.Println("ReadeBytes err:", err)
		}
		fmt.Print(string(buf))
	}
}
