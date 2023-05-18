package main

import "fmt"

func main() {
	fmt.Println("服务器开始运行")
	server := NewServer("127.0.0.1", 8888)
	server.Start()

}