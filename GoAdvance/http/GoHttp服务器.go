package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 注册回调函数，该函数在客户端访问服务器时，会被自动调用
	http.HandleFunc("/test", myHandle)

	// 绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func myHandle(writer http.ResponseWriter, request *http.Request) {
	// 写给客户端的数据内容
	writer.Write([]byte("this is a Web server"))

	// 从客户端读到的内容
	fmt.Println("Header: ", request.Header)
	fmt.Println("URL: ", request.URL)
	fmt.Println("Method: ", request.Method)
	fmt.Println("Host: ", request.Host)
	fmt.Println("RemoteAddr: ", request.RemoteAddr)
	fmt.Println("Body: ", request.Body)

}
