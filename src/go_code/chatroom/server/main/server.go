package main

import (
	"GoLandProjects/src/go_code/chatroom/server/model"
	"fmt"
	"net"
	"time"
)

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool) // 这里的pool是全局变量，在redis.go里面
}

func main() {
	// 注意先初始化Pool，再初始化UserDao
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	listen, err := net.Listen("tcp", "0.0.0.0:6250")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("waiting connect...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
		} else {
			fmt.Printf("%v 连接成功!\n", conn.RemoteAddr().String())
		}
		// 启动协程, 与建立连接的客户端进行交互
		go process(conn)
	}
}

// process 用于处理客户端的请求
func process(conn net.Conn) {
	// 读取客户端发送的消息
	defer conn.Close()

	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		// fmt.Println("协程出错:", err)
		return
	}
}
