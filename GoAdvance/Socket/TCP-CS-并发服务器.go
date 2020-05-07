package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 创建监听Socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}
	defer listener.Close()

	//监听客户端连接请求
	for {
		fmt.Println("服务器等待客户端连接...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err: ", err)
			return
		}

		// 具体完成服务器和客户端的数据通信
		go handerConnect(conn)
	}
}

func handerConnect(conn net.Conn) {
	defer conn.Close()

	// 获取连接的客户端地址
	addr := conn.RemoteAddr()
	fmt.Println(addr, "客户端成功连接！！！")

	// 循环读取客户端发送的数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if "exit\n" == string(buf[:n]) || "exit\r\n" == string(buf[:n]) {
			fmt.Println("服务器接收到客户端退出请求，服务器关闭。")
			return
		}
		if n == 0 {
			fmt.Println("服务器检测到客户端已关闭，断开连接！！！")
			return
		}
		if err != nil {
			fmt.Println("conn.Read err: ", err)
			return
		}
		fmt.Println("服务器读到数据： ", string(buf[:n]))

		// 小写转大写，回发给客户端
		conn.Write([]byte(strings.ToUpper(string(buf[0:n]))))
	}

}
