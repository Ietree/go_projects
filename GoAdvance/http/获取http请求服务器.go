package main

import (
	"fmt"
	"net"
	"os"
)

func errFunc(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		// 将当前进程结束
		os.Exit(1)
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	errFunc(err, "net.Listen err: ")
	defer listener.Close()

	conn, err := listener.Accept()
	errFunc(err, "Accept err:")
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if n == 0 {
		return
	}
	errFunc(err, "conn.Read err: ")

	fmt.Println(string(buf[:n]))
}
