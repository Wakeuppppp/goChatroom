/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/28 15:42
 * @Project_Name : GoLandProjects
 * @File : userProcess.go
 * @Software :GoLand
 */

// Package processes 处理和用户相关的业务
package processes

import (
	"GoLandProjects/src/go_code/chatroom/common/message"
	"GoLandProjects/src/go_code/chatroom/server/model"
	"GoLandProjects/src/go_code/chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn     net.Conn
	UserId   int
	UserName string
}

// ServerProcessLogin 处理登陆
func (userProcess *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("serverProcessLogin unmarshal err:", err)
		return err
	}

	var resMes message.Message // 返回给客户端的消息实体
	resMes.Type = message.LoginReMesType

	var loginResMes message.LoginResMes

	// 数据库验证
	// 使用UserDao去redis验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
	} else {
		loginResMes.Code = 200
		loginResMes.Error = "登陆成功"
		loginResMes.UserName = user.UserName
		loginResMes.UserList = make(map[int]string, 10)
		userProcess.UserId = user.UserId
		userProcess.UserName = user.UserName

		// 将用户添加到在线用户列表中
		userMgr.AddOnlineUser(userProcess)

		// 通知其他在线用户有新用户上线
		err = userProcess.NotifyOthersOnlineUser(userProcess.UserId, userProcess.UserName)
		if err != nil {
			fmt.Println("NotifyOthersOnlineUser err:", err)
			return err
		}
		// 遍历userMgr.OnlineUsers切片, 将在线用户的Id放入到loginResMes.UsersId切片中去
		fmt.Println("在线用户:")

		for id, v := range userMgr.OnlineUsers {
			// loginResMes.UsersId = append(loginResMes.UsersId, id)
			loginResMes.UserList[id] = v.UserName
			fmt.Printf("%d %v\n", id, v.UserName)
		}
		// for i, v := range userMgr.OnlineUsers {
		// 	fmt.Printf("%d %v\n", i, v.UserName)
		// }
		fmt.Println()
	}

	// 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("loginResMes marshal err:", err)
		return err
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("loginMes marshal err:", err)
		return err
	}

	// 调用专门的函数来发送数据包
	transfer := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = transfer.WritePkg(data)
	return err
}

// ServerProcessRegister 处理注册
func (userProcess *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("ServerProcessRegister unmarshal err:", err)
		return err
	}
	fmt.Println(registerMes)
	var resMes message.Message // 返回给客户端的消息实体
	resMes.Type = message.RegisterResMesType

	var registerResMes message.RegisterResMes

	// 数据库验证
	// 使用UserDao去redis验证用户是否存在, 若不存在才允许注册
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			// fmt.Println("用户已被注册")
			registerResMes.Code = 400
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
			// fmt.Println(registerResMes.Code, registerResMes.Error)
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "未知错误"
			fmt.Println(registerResMes.Code, registerResMes.Error)
		}
	} else {
		// fmt.Println("注册成功~~~")
		registerResMes.Code = 200
		registerResMes.Error = "注册成功"
		// fmt.Println(registerResMes.Code, registerResMes.Error)
	}

	// 序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("registerResMes marshal err:", err)
		return err
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("registerMes marshal err:", err)
		return err
	}

	// 调用专门的函数来发送数据包
	transfer := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = transfer.WritePkg(data)
	return
}

// NotifyOthersOnlineUser 用于通知其他用户新用户(userId)上线
func (userProcess *UserProcess) NotifyOthersOnlineUser(userId int, userName string) (err error) {
	// 序列化通知消息
	data := MarshalNotifyData(userId, userName)

	// 遍历OnlineUsers, 然后逐个发送
	for i, up := range userMgr.OnlineUsers {
		if i == userId {
			continue
		}
		// 这里的up是其他在线用户的连接, 要通知其他用户就需要用到其他用户的连接
		// 传入的userId是此刻上线的用户
		up.NotifyMeOnline(data)
	}
	return err
}

// NotifyMeOnline 将已经序列化的消息 发送给在线用户
func (userProcess *UserProcess) NotifyMeOnline(data []byte) {
	transfer := &utils.Transfer{Conn: userProcess.Conn}
	err := transfer.WritePkg(data)
	if err != nil {
		return
	}
}

// MarshalNotifyData 用于将通知信息序列化
func MarshalNotifyData(userId int, userName string) (data []byte) {
	// 组装NotifyUserStatusMes
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.UserName = userName
	notifyUserStatusMes.UserStatus = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		return
	}

	var mes message.Message
	mes.Data = string(data)
	mes.Type = message.NotifyUserStatusMseType

	data, err = json.Marshal(mes)
	if err != nil {
		return
	}
	return data
}

// PushOfflineUser 用户下线时通知其他在线用户
func (userProcess *UserProcess) PushOfflineUser(mes *message.Message) {

	var offlineMes message.OfflineMes

	err := json.Unmarshal([]byte(mes.Data), &offlineMes)
	if err != nil {
		return
	}
	// 将用户从在线列表中删除

	userMgr.DelOnlineUser(offlineMes.UserId)

	var offlineResMes message.OfflineResMes
	offlineResMes.UserId = offlineMes.UserId
	offlineResMes.UserName = offlineMes.UserName
	offlineResMes.Code = 200
	offlineResMes.Err = "用户已下线"
	data, err := json.Marshal(offlineResMes)
	if err != nil {
		return
	}

	mes.Type = message.OfflineResMesType
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		return
	}

	// 通知其他在线用户
	for _, conn := range userMgr.OnlineUsers {
		conn.NotifyMeOnline(data)
	}
}
