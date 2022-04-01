package main

import (
	"GoLandProjects/src/go_code/chatroom/client/process"
	"fmt"
)

var (
	userId   int
	userPwd  string
	userName string
)

func main() {
	menu()
}

func menu() {
	var key int
	loop := true
	for loop {
		fmt.Println("------------Welcome to chatroom------------")
		fmt.Println("\t\t 1 登陆")
		fmt.Println("\t\t 2 注册")
		fmt.Println("\t\t 3 退出")
		fmt.Printf("请选择(1~3): ")
		fmt.Scanln(&key)
		fmt.Println()

		switch key {
		case 1:
			fmt.Printf("请输入用户ID: ")
			fmt.Scanln(&userId)
			fmt.Printf("请输入用户密码: ")
			fmt.Scanln(&userPwd)
			userProcess := process.UserProcess{}
			err := userProcess.Login(userId, userPwd)
			if err != nil {
				fmt.Println("登陆失败:", err)
			} else {
				fmt.Println("else", err)
			}
		case 2:
			fmt.Println("注册用户")
			fmt.Printf("请设置用户ID: ")
			fmt.Scanln(&userId)
			fmt.Printf("请设置密码: ")
			fmt.Scanln(&userPwd)
			fmt.Printf("请设置用户名称: ")
			fmt.Scanln(&userName)
			userProcess := process.UserProcess{}
			err := userProcess.Register(userId, userPwd, userName)
			if err != nil {
				fmt.Println("注册失败", err)
				return
			}
			// loop = false
		case 3:
			fmt.Printf("欢迎下次继续使用本软件, 祝您生活愉快~\n")
			// os.Exit(0)
			loop = false
		default:
			fmt.Println("输入错误, 请重新选择")
		}
	}
}
