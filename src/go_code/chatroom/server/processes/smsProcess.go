/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/28 15:43
 * @Project_Name : GoLandProjects
 * @File : smsProcess.go
 * @Software :GoLand
 */

package processes

import (
	"GoLandProjects/src/go_code/chatroom/common/message"
	"GoLandProjects/src/go_code/chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

// SendGroupMes1 服务器转发消息给所有的在线用户
func (smsProcess *SmsProcess) SendGroupMes1(mes *message.Message) (err error) {
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		return
	}
	transfer := &utils.Transfer{}
	for i, v := range userMgr.OnlineUsers {
		if i == smsMes.UserId {
			continue
		}
		transfer.Conn = v.Conn
		data, err := json.Marshal(mes)
		if err != nil {
			fmt.Println("done!!!!!!!!!!")
			return err
		}
		err = transfer.WritePkg(data)
		if err != nil {
			return err
		}
	}
	fmt.Println(smsMes.Content)
	return
}

// SendGroupMes 服务器转发消息给所有的在线用户
func (smsProcess *SmsProcess) SendGroupMes(mes *message.Message) (err error) {
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		return
	}

	for i, v := range userMgr.OnlineUsers {
		if i == smsMes.UserId {
			continue
		}
		smsProcess.SendMesToEachOnlineUser(data, v.Conn)
	}

	return
}

// SendMesToEachOnlineUser 给单个用户发送消息
func (smsProcess *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {

	transfer := &utils.Transfer{Conn: conn}
	err := transfer.WritePkg(data)
	if err != nil {
		return
	}
}

// SendSingleMes 给指定用户发消息
func (smsProcess *SmsProcess) SendSingleMes(mes *message.Message) (err error) {
	var singleMes message.SingleMes
	err = json.Unmarshal([]byte(mes.Data), &singleMes)
	if err != nil {
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		return
	}
	smsProcess.SendMesToEachOnlineUser(data, userMgr.OnlineUsers[singleMes.ReceiveUserId].Conn)
	return
}
