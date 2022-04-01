/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/28 17:21
 * @Project_Name : GoLandProjects
 * @File : userProcess.go
 * @Software :GoLand
 */

// Package process 处理和用户相关的业务
package process

import (
	"GoLandProjects/src/go_code/chatroom/client/utils"
	"GoLandProjects/src/go_code/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type UserProcess struct {
}

func (userProcess *UserProcess) Login(userId int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "39.99.119.232:6250")
	if err != nil {
		fmt.Println("Dial err:", err)
		return err
	}
	defer conn.Close()

	// 实例化一个 消息 实体，并声明消息的类型
	var mes message.Message
	mes.Type = message.LoginMesType

	// 实例化一个 登陆消息 实体，并序列化
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 登陆消息序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("loginMes json.Marshal err:", err)
		return err
	}
	// fmt.Println("序列化后的data:", data)
	// 将序列化后的 登陆消息 填充到 消息的Data字段
	mes.Data = string(data)

	// 再将 消息 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("message json.Marshal err:", err)
		return err
	}

	// 先发送data 的长度，再发送data
	// 获取data长度，将长度(int)转换成[]byte后再发送
	var pkgLen uint32
	pkgLen = uint32(len(data))

	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)

	// 发送长度
	n, err := conn.Write(buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.Write:", err)
		return err
	}

	// 发送数据本身
	n, err = conn.Write(data)
	if err != nil {
		fmt.Println("发送数据失败: ", err)
		return err
	}
	// fmt.Println(string(data))

	transfer := &utils.Transfer{
		Conn: conn,
	}
	mes, err = transfer.ReadPkg()

	if err != nil {
		fmt.Println("读取数据失败: ", err)
		return err
	}

	// 反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("反序列化数据失败: ", err)
		return err
	}

	if loginResMes.Code == 200 {
		fmt.Printf("\n------------------登陆成功----------------\n")
		err = errors.New(loginResMes.Error)
		// 对CurUser初始化
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserName = loginResMes.UserName
		CurUser.UserStatus = message.UserOnline

		// fmt.Println("当前在线用户:")
		for i, v := range loginResMes.UserList {
			if i == userId {
				continue
			}
			// fmt.Println("用户Id:", v)
			// 将在线用户同步到 OnlineUsers中
			// OnlineUsers是客户端用来维护在线用户的map[int]*message.User
			user := &message.User{
				UserId:     i,
				UserName:   v,
				UserStatus: message.UserOnline,
			}
			OnlineUsers[i] = user
		}
		fmt.Println()

		// 该协程用于与服务器保持通, 接收服务器端的消息推送
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}
	} else {
		err = errors.New(loginResMes.Error)
	}
	return err
}

func (userProcess *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	conn, err := net.Dial("tcp", "39.99.119.232:6250")
	if err != nil {
		fmt.Println("Dial err:", err)
		return err
	}
	defer conn.Close()

	// 实例化一个 消息 实体，并声明消息的类型
	var mes message.Message
	mes.Type = message.RegisterMesType

	// 实例化一个 注册消息 实体，并序列化
	var RegisterMes message.RegisterMes
	RegisterMes.User.UserId = userId
	RegisterMes.User.UserPwd = userPwd
	RegisterMes.User.UserName = userName

	// 注册消息序列化
	data, err := json.Marshal(RegisterMes)
	if err != nil {
		fmt.Println("RegisterMes json.Marshal err:", err)
		return err
	}
	// fmt.Println("序列化后的data:", data)
	// 将序列化后的 注册消息 填充到 消息的Data字段
	mes.Data = string(data)

	// 再将 消息 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("message json.Marshal err:", err)
		return err
	}

	transfer := &utils.Transfer{
		Conn: conn,
	}

	// 发送注册信息给服务器
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("发送数据失败: ", err)
		return err
	}

	// 接收从服务器传回来的消息
	mes, err = transfer.ReadPkg()
	if err != nil {
		fmt.Println("读取数据失败: ", err)
		return err
	}

	// 反序列化
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("反序列化数据失败: ", err)
		return err
	}

	// 根据返回消息的Code字段处理
	if registerResMes.Code == 200 {
		fmt.Printf("\n\n------------注册成功, 请重新登录------------\n\n\n")
	} else if registerResMes.Code == 400 {
		// fmt.Println("注册失败")
		fmt.Println(registerResMes.Code, registerResMes.Error)
	}
	return
}
