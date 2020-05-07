package main

import (
	"fmt"
	"net"
)

func main() {
	// 1、指定 服务器 IP + Port创建Socket
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dail err: ", err)
		return
	}
	defer conn.Close()

	// 2、主动写数据给服务器
	conn.Write([]byte("Are you ready?"))

	// 3、接收服务器回发的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err: ", err)
		return
	}
	fmt.Println("服务器回发: ", string(buf[:n]))
}
