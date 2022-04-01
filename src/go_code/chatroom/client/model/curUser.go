/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/31 16:04
 * @Project_Name : GoLandProjects
 * @File : curUser.go
 * @Software :GoLand
 */

package model

import (
	"GoLandProjects/src/go_code/chatroom/common/message"
	"net"
)

// CurUser 在客户端很多地方可能会使用, 做成全局
type CurUser struct {
	Conn net.Conn
	message.User
}
