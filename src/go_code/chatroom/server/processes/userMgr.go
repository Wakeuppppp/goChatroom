/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/30 10:38
 * @Project_Name : GoLandProjects
 * @File : userMgr.go
 * @Software :GoLand
 */

package processes

import "fmt"

// UserMgr实例在服务器端有且只有一个
var (
	userMgr *UserMgr
)

type UserMgr struct {
	OnlineUsers map[int]*UserProcess
}

// init 用于初始化
func init() {
	userMgr = &UserMgr{OnlineUsers: make(map[int]*UserProcess, 1024)}

}

// AddOnlineUser 用户上线
func (userMgr *UserMgr) AddOnlineUser(userProcess *UserProcess) {
	// fmt.Println("-----------", userProcess.UserId)
	userMgr.OnlineUsers[userProcess.UserId] = userProcess
}

// DelOnlineUser 用户下线
func (userMgr *UserMgr) DelOnlineUser(userId int) {
	delete(userMgr.OnlineUsers, userId)
}

// GetAllOnlineUser 返回在线用户
func (userMgr *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return userMgr.OnlineUsers
}

// GetOnlineUserById 根据Id返回对应的值
func (userMgr *UserMgr) GetOnlineUserById(userId int) (useProcess *UserProcess, err error) {
	useProcess, ok := userMgr.OnlineUsers[userId]
	if !ok { // 说明要查找的用户当前不在线
		err = fmt.Errorf("用户%d不在线", userId)
		return
	}
	return
}
