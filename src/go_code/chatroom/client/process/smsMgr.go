/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/31 23:40
 * @Project_Name : GoLandProjects
 * @File : smsMgr.go
 * @Software :GoLand
 */

package process

import (
	"GoLandProjects/src/go_code/chatroom/common/message"
	"encoding/json"
	"fmt"
)

func OutputGroupMes(mes message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		return
	}
	fmt.Printf("\n来自%v(%v)的群发消息:%v\n", smsMes.UserName, smsMes.UserId, smsMes.Content)
}

func OutputSingleMes(mes message.Message) {
	var singleMes message.SingleMes
	err := json.Unmarshal([]byte(mes.Data), &singleMes)
	if err != nil {
		return
	}
	fmt.Printf("\n来自%v(%v)的私聊:%v\n", singleMes.UserName, singleMes.UserId, singleMes.Content)
}
