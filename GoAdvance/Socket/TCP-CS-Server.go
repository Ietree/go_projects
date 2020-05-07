package main

import (
	"fmt"
	"net"
)

func main() {
	// 1、指定服务器的通信协议、IP地址、port，创建一个用于监听的socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen() err: ", err)
		return
	}
	defer listener.Close()

	// 2、阻塞监听客户端连接请求，成功建立连接，返回用于通信的socket
	fmt.Println("服务器等待客户端建立连接...")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("net.Accept() err: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("服务器与客户端成功建立连接！！！")

	// 3、读取客户端发送的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err: ", err)
		return
	}

	// 4、写出服务器读到的数据
	conn.Write(buf[:n])
	fmt.Println("服务器读到数据： ", string(buf[:n]))
}
