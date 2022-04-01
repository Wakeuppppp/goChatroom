/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/28 17:21
 * @Project_Name : GoLandProjects
 * @File : smsProcess.go
 * @Software :GoLand
 */

package process

import (
	"GoLandProjects/src/go_code/chatroom/client/utils"
	"GoLandProjects/src/go_code/chatroom/common/message"
	"encoding/json"
)

type SmsProcess struct {
}

// SendGroupMes 发送群发消息
func (smsProcess *SmsProcess) SendGroupMes(content string) (err error) {
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserName = CurUser.UserName
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		return
	}

	var mes message.Message
	mes.Type = message.SmsMesType
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
	return
}

// SendSingleMes 发送私发消息
func (smsProcess *SmsProcess) SendSingleMes(userId int, content string) (err error) {
	var singleMes message.SingleMes
	singleMes.Content = content
	singleMes.UserId = CurUser.UserId
	singleMes.UserName = CurUser.UserName
	singleMes.ReceiveUserId = userId

	data, err := json.Marshal(singleMes)
	if err != nil {
		return
	}

	var mes message.Message
	mes.Type = message.SingleMesType
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

	return
}
