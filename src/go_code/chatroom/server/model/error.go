/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/28 22:55
 * @Project_Name : GoLandProjects
 * @File : error.go
 * @Software :GoLand
 */

package model

import "errors"

// 根据业务逻辑的需要 自定义一些错误
var (
	ERROR_USER_NOTEXISTS = errors.New("用户不存在")
	ERROR_USER_EXISTS    = errors.New("用户已存在")
	ERROR_USER_PWD       = errors.New("密码错误")
)
