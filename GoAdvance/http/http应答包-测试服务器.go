package main

import "net/http"

func main() {
	// 注册回调函数，该回调函数会在服务器被访问时自动被调用
	http.HandleFunc("/test", handler)

	// 绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello, World"))
}
