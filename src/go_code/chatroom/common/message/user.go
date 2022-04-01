/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/28 22:55
 * @Project_Name : GoLandProjects
 * @File : user.go
 * @Software :GoLand
 */

package message

type User struct {
	UserId     int    `json:"userId"`     // 用户ID
	UserPwd    string `json:"userPwd"`    // 用户密码
	UserName   string `json:"userName"`   // 用户昵称
	UserStatus int    `json:"userStatus"` // 用户状态
}
