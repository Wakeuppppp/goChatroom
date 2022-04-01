/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/31 10:50
 * @Project_Name : GoLandProjects
 * @File : userMgr.go
 * @Software :GoLand
 */

package process

import (
	"GoLandProjects/src/go_code/chatroom/client/model"
	"GoLandProjects/src/go_code/chatroom/client/utils"
	"GoLandProjects/src/go_code/chatroom/common/message"
	"encoding/json"
	"fmt"
)

// OnlineUsers 客户端维护的在线用户的map
var OnlineUsers = make(map[int]*message.User, 10)
var CurUser model.CurUser

// ShowOnlineUser 在客户端显示当前在线用户
func ShowOnlineUser() {
	fmt.Println()
	fmt.Println("当前在线用户列表:")
	fmt.Println("---------------------")
	fmt.Println()
	for i, v := range OnlineUsers {
		fmt.Printf("%v(%v)\n", v.UserName, i)
		fmt.Println()
	}
	fmt.Println("---------------------")
	fmt.Println()
}

// UpdateUserStatus 处理返回的NotifyUserStatusMes
func UpdateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	// 保存上线用户信息到map中
	fmt.Println(notifyUserStatusMes)
	user, ok := OnlineUsers[notifyUserStatusMes.UserId]
	if !ok { // 说明之前没有
		user = &message.User{
			UserId:   notifyUserStatusMes.UserId,
			UserName: notifyUserStatusMes.UserName,
		}
	}
	// 更新状态
	OnlineUsers[notifyUserStatusMes.UserId] = user
	// 更新map
	user.UserStatus = notifyUserStatusMes.UserStatus
}

// Offline 用户下线通知
func Offline(userId int, userName string) {
	var offlineMes message.OfflineMes
	// offlineMes.Conn = CurUser.Conn
	offlineMes.UserId = userId
	offlineMes.UserName = userName
	data, err := json.Marshal(offlineMes)
	if err != nil {
		return
	}

	var mes message.Message
	mes.Type = message.OfflineMesType
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		return
	}

	transfer := &utils.Transfer{Conn: CurUser.Conn}
	err = transfer.WritePkg(data)
	if err != nil {
		return
	}

}
