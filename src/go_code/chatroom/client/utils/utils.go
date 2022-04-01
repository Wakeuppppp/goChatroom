/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/28 15:44
 * @Project_Name : GoLandProjects
 * @File : utils.go
 * @Software :GoLand
 */

// Package utils 将消息反序列化
package utils

import (
	"GoLandProjects/src/go_code/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// Transfer 将这些方法关联到结构体中 ※※※
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte // 传输时使用的缓冲
}

// ReadPkg 服务器接收客户端的数据包
func (transfer *Transfer) ReadPkg() (mes message.Message, err error) {
	// buf := make([]byte, 8096)
	// fmt.Println("读取客户端发送的数据...")
	_, err = transfer.Conn.Read(transfer.Buf[:4])
	if err != nil {
		// fmt.Println("读取消息长度失败", err)
		return
	}
	// 根据buf[:4] 转成一个 uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(transfer.Buf[:4])
	// fmt.Println("要读取的长度为:", pkgLen)

	// 根据接收到的消息长度来读取buf中的内容, 防止多读或少读
	n, err := transfer.Conn.Read(transfer.Buf[:pkgLen])
	if err != nil || n != int(pkgLen) {
		fmt.Println("读取消息内容失败:", err)
		return
	}

	// 将pkgLen反序列化
	err = json.Unmarshal(transfer.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return
	}
	// fmt.Println("反序列化成功")
	return
}

// WritePkg 用于发送数据包
func (transfer *Transfer) WritePkg(data []byte) (err error) {
	// 先发送消息的长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))

	// var buf [4]byte
	binary.BigEndian.PutUint32(transfer.Buf[:4], pkgLen)

	// 发送长度
	n, err := transfer.Conn.Write(transfer.Buf[:4])
	if err != nil || n != 4 {
		fmt.Println("writePkg err:", err)
		return err
	}

	// 发送消息本身
	n, err = transfer.Conn.Write(data)
	if err != nil || n != int(pkgLen) {
		fmt.Println("writePkg err:", err)
		return err
	}
	return err
}
