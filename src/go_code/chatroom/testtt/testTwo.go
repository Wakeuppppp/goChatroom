/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/28 23:18
 * @Project_Name : GoLandProjects
 * @File : testTwo.go
 * @Software :GoLand
 */

package main

import "fmt"

type p struct {
	userId   int
	userPwd  string
	userName string
}

func tt() (p1 p) {
	p1.userName = "jack"
	p1.userId = 120
	p1.userPwd = "pass"
	fmt.Println(p1)
	return
}

func main() {
	fmt.Println("ssss")
	p2 := tt()
	fmt.Println(p2)
}
