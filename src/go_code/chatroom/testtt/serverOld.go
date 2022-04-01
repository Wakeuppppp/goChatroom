/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/24 22:45
 * @Project_Name : GoLandProjects
 * @File : serverOld.go
 * @Software :GoLand
 */

package main

import (
	"fmt"
	"io"
	"net"
)

func sProcess1(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("等待输入%v...\n", conn.RemoteAddr().String())
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Printf("客户端%v断开连接\n", conn.RemoteAddr().String())
			return
		}
		// 显示客户端发送的数据
		fmt.Printf("来自%v的消息: %v", conn.RemoteAddr().String(), string(buf[:n]))
	}
}

func mains() {
	listen, err := net.Listen("tcp", "0.0.0.0:6250")
	if err != nil {
		fmt.Println("err:", err)
	}
	defer listen.Close()

	for {
		fmt.Println("waiting connect...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
		} else {
			fmt.Printf("Accept succ! 客户端IP: %v\n", conn.RemoteAddr().String())
		}
		go sProcess1(conn)
		// 启动协程, 与建立连接的客户端进行交互
	}
}
