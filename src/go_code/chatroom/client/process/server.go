/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/28 17:22
 * @Project_Name : GoLandProjects
 * @File : server.go
 * @Software :GoLand
 */

package process

import (
	"GoLandProjects/src/go_code/chatroom/client/utils"
	"GoLandProjects/src/go_code/chatroom/common/message"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Printf("Hello, %v(%v)\n", CurUser.UserName, CurUser.UserId)
	fmt.Println("1. 显示在线用户")
	fmt.Println("2. 群发消息")
	fmt.Println("3. 发起聊天")
	fmt.Println("4. 信息列表")
	fmt.Println("5. 退出系统")
	fmt.Printf("请选择(1-5):")
	var key int
	var content string          // 发送的消息
	smsProcess := &SmsProcess{} // 用来群发/私发消息
	var id int                  // 私发消息用户Id
	fmt.Scanln(&key)
	switch key {
	case 1:
		ShowOnlineUser()
	case 2:
		fmt.Println("2. 群发消息")
		for {
			fmt.Printf("请输入要发送的消息:")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan() // use `for scanner.Scan()` to keep reading
			content = scanner.Text()
			if content == "quit" {
				fmt.Println("退出咯~")
				break
			}
			err := smsProcess.SendGroupMes(content)
			if err != nil {
				return
			}
		}
	case 3:
		fmt.Println("3. 发起聊天")
		ShowOnlineUser()
		fmt.Printf("请输入要发起聊天的用户Id:")
		fmt.Scanln(&id)

		for {
			_, ok := OnlineUsers[id]
			if !ok {
				fmt.Printf("\n用户%v不在线\n\n", id)
				break
			}
			fmt.Printf("\n请输入要发送的消息:")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan() // use `for scanner.Scan()` to keep reading
			content = scanner.Text()
			if content == "quit" {
				fmt.Println("退出咯~")
				break
			}
			// 调用相关函数发送消息
			err := smsProcess.SendSingleMes(id, content)
			if err != nil {
				return
			}
		}
	case 4:
		fmt.Println("4. 信息列表")
	case 5:
		Offline(CurUser.UserId, CurUser.UserName)
		fmt.Printf("\n再见!亲爱的%v(%v)用户!\n", CurUser.UserName, CurUser.UserId)
		os.Exit(0)
	default:
		fmt.Println("选择错误")
	}
}

// serverProcessMes 用于和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	transfer := &utils.Transfer{
		Conn: conn,
	}
	for {
		// fmt.Println("客户端正在等待读取服务器端消息:")
		mes, err := transfer.ReadPkg()
		if err != nil {
			return
		}
		switch mes.Type {
		case message.NotifyUserStatusMseType: // 有人上线

			// 将mes.Data反序列化成NotifyUserStatusMes, 取出其中的UserId
			var notifyUserStatusMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				return
			}
			UpdateUserStatus(&notifyUserStatusMes)

			fmt.Printf("\n\n----------------用户%v(%v)上线了----------------\n\n", notifyUserStatusMes.UserName, notifyUserStatusMes.UserId)
			fmt.Printf("Hello, %v(%v)\n", CurUser.UserName, CurUser.UserId)
			fmt.Println("1. 显示在线用户")
			fmt.Println("2. 群发消息")
			fmt.Println("3. 发起聊天")
			fmt.Println("4. 信息列表")
			fmt.Println("5. 退出系统")
			fmt.Printf("请选择(1-5):")
		case message.SmsMesType: // 有群发消息
			fmt.Println()
			OutputGroupMes(mes)
		case message.SingleMesType: // 有私聊消息
			fmt.Println()
			OutputSingleMes(mes)
		case message.OfflineResMesType:
			var offlineResMes message.OfflineResMes
			err := json.Unmarshal([]byte(mes.Data), &offlineResMes)
			if err != nil {
				return
			}
			if offlineResMes.UserId == CurUser.UserId {
				return
			}
			Off(offlineResMes.UserId, offlineResMes.UserName)
		default:
			fmt.Println("暂时无法处理此类型消息")
		}
	}
}

func Off(userId int, userName string) {

	fmt.Printf("\n----------------用户%v(%v)已离线----------------", userName, userId)
	delete(OnlineUsers, userId)
}
