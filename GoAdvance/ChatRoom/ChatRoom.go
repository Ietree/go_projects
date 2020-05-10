package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// 创建用户结构体类型
type Client struct {
	c    chan string
	name string
	addr string
}

// 创建全局map，存储在线用户
var onlineMap map[string]Client

// 创建全局channel传递用户消息
var message = make(chan string)

func handlerConnect(conn net.Conn) {
	defer conn.Close()

	// 创建channel判断用户是否活跃
	hasData := make(chan bool)
	// 获取用户网络地址 IP + Port
	addr := conn.RemoteAddr().String()
	// 创建连接新用户的结构体，默认用户是 IP + Port
	clnt := Client{make(chan string), addr, addr}
	// 将新连接用户添加到在线用户map中，key：IP + Port value：client
	onlineMap[addr] = clnt

	// 创建专门用来给当前用户发送消息的go程
	go writeMsgToClient(clnt, conn)

	// 发送用户上线消息到全局channel中
	//message <- "[" + addr + "]" + clnt.name + "login"
	message <- makeMsg(clnt, "login")

	// 创建一个channel，用来判断用户退出状态
	isQuit := make(chan bool)

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Printf("检测到客户端：%s退出\n", clnt.name)
				return
			}
			if err != nil {
				fmt.Println("conn.Read err", err)
				return
			}
			msg := string(buf[:n-1])

			// 判断输入命令
			// 提取在线用户列表
			if msg == "who" && len(msg) == 3 {
				conn.Write([]byte("online user list:\n"))
				// 遍历当前map，获取在线用户
				for _, user := range onlineMap {
					userInfo := user.addr + ":" + user.name + "\n"
					conn.Write([]byte(userInfo))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				newName := strings.Split(msg, "|")[1]
				clnt.name = newName
				onlineMap[addr] = clnt
				conn.Write([]byte("rename successful\n"))
			} else {
				// 将读到的用户消息，写入到message中
				message <- makeMsg(clnt, msg)
			}
			hasData <- true
		}
	}()

	// 保证不退出
	for {
		// 监听channel上的数据流动
		select {
		case <-isQuit:
			// 删除退出的用户
			delete(onlineMap, clnt.addr)
			// 写入用户退出消息到全局channel
			message <- makeMsg(clnt, "logout")
			return
		case <-hasData:
			// 目的是为了重置下面case的计时器
		case <-time.After(time.Second * 10):
			// 删除退出的用户
			delete(onlineMap, clnt.addr)
			// 写入用户退出消息到全局channel
			message <- makeMsg(clnt, "logout")
			return
		}
	}
}

func makeMsg(clnt Client, msg string) (buf string) {
	buf = "[" + clnt.addr + "]" + clnt.name + msg
	return
}

func writeMsgToClient(clnt Client, conn net.Conn) {
	// 监听用户自带Channel上是否有消息
	for msg := range clnt.c {
		conn.Write([]byte(msg + "\n"))
	}
}

func manager() {
	// 初始化onlineMap
	onlineMap = make(map[string]Client)

	// 监听全局channel中是否有数据，有数据存储至msg，无数据阻塞
	for {
		msg := <-message

		// 循环发送消息给所有在线用户
		for _, clnt := range onlineMap {
			clnt.c <- msg
		}
	}
}

func main() {
	// 创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Listen err", err)
		return
	}
	defer listener.Close()

	// 创建管理者go程，管理map和全局channel
	go manager()

	// 循环监听客户端连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err", err)
			return
		}
		// 启动go程处理客户端数据请求
		go handlerConnect(conn)
	}
}
