/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/28 15:44
 * @Project_Name : GoLandProjects
 * @File : processor.go.go
 * @Software :GoLand
 */

package main

import (
	"GoLandProjects/src/go_code/chatroom/common/message"
	"GoLandProjects/src/go_code/chatroom/server/processes"
	"GoLandProjects/src/go_code/chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (processor *Processor) process2() (err error) {
	for {
		transfer := &utils.Transfer{
			Conn: processor.Conn,
		}
		mes, err := transfer.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%v已退出!\n", processor.Conn.RemoteAddr().String())
				return err
			} else {
				fmt.Printf("%v断开连接\n", processor.Conn.RemoteAddr().String())
				fmt.Println()
				return err
			}
		}
		err = processor.ServerProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessMes err", err)
			return err
		}
	}
}

// ServerProcessMes 根据消息类型调用不同的函数
func (processor *Processor) ServerProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType: // 用户登陆请求
		userProcess := &processes.UserProcess{
			Conn: processor.Conn,
		}
		err = userProcess.ServerProcessLogin(mes)
		if err != nil {
			return
		}

	case message.RegisterMesType: // 用户请求注册
		userProcess := &processes.UserProcess{
			Conn: processor.Conn,
		}
		err = userProcess.ServerProcessRegister(mes)
		if err != nil {
			return
		}

	case message.SmsMesType: // 用户群发消息
		smsProcess := &processes.SmsProcess{}
		err = smsProcess.SendGroupMes(mes)
		if err != nil {
			return
		}
	case message.SingleMesType:
		// 私发消息
		smsProcess := &processes.SmsProcess{}
		err = smsProcess.SendSingleMes(mes)
		if err != nil {
			return
		}
	case message.OfflineMesType:
		userProcess := &processes.UserProcess{
			Conn: processor.Conn,
		}
		userProcess.PushOfflineUser(mes)
	default:
		fmt.Println("消息种类不存在, 无法处理")
	}
	return
}
