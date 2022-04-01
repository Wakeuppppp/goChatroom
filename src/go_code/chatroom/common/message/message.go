/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/26 15:56
 * @Project_Name : GoLandProjects
 * @File : message.go
 * @Software :GoLand
 */

package message

const (
	LoginMesType            = "LoginMes"
	LoginReMesType          = "LoginReMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMseType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
	SingleMesType           = "SingleMes"
	OfflineMesType          = "OfflineMes"
	OfflineResMesType       = "OfflineResMes"
)

// 定义用户状态
const (
	UserOnline     = iota // 用户在线
	UserOffline           // 用户下线
	UserBusyStatus        // 用户忙碌
)

// Message 消息的统一接口
type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息数据
}

// LoginMes 登陆消息
type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

// LoginResMes 登陆返回消息
type LoginResMes struct {
	Code     int    `json:"code"`     // 状态码 500 表示用户未注册, 200 表示登陆成功, 403表示密码错误
	Error    string `json:"error"`    // 返回错误信息
	UserName string `json:"userName"` // 用户的昵称
	// UsersId  []int          `json:"usersId"`  // 保存在线用户id的切片
	UserList map[int]string `json:"userList"` // 保存在线用户列表
}

// RegisterMes 注册消息
type RegisterMes struct {
	User User `json:"user"`
}

// RegisterResMes 注册返回消息
type RegisterResMes struct {
	Code  int    `json:"code"` // 返回状态码 200表示注册成功, 400 表示用户已存在
	Error string `json:"error"`
}

// NotifyUserStatusMes 用于服务器推送用户状态变换的信息
type NotifyUserStatusMes struct {
	UserId     int    `json:"userId"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"`
}

// SmsMes 群发消息
type SmsMes struct {
	Content string `json:"content"`
	User
}

// SingleMes 单独消息
type SingleMes struct {
	Content       string `json:"content"`
	ReceiveUserId int    `json:"receiveUserId"`
	User
}

// OfflineMes 用户下线消息
type OfflineMes struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
}

type OfflineResMes struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	Code     int    `json:"code"`
	Err      string `json:"err"`
}
