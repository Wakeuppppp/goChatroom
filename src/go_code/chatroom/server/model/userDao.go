/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/28 22:55
 * @Project_Name : GoLandProjects
 * @File : userDao.go
 * @Software :GoLand
 */

package model

import (
	"GoLandProjects/src/go_code/chatroom/common/message"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 定义一个结构体，实现对User 结构体的各种操作

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

// NewUserDao 使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (userDao *UserDao) getUserById(conn redis.Conn, id int) (user *message.User, err error) {
	res, err := redis.String(conn.Do("hGet", "users", id))
	if err != nil {
		if err == redis.ErrNil { // 表示在users哈希中，没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	// 将res反序列化成一个对象实例
	// user = &User{}
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		return
	}
	return
}

// Login 完成对用户的验证
func (userDao *UserDao) Login(userId int, userPwd string) (user *message.User, err error) {
	conn := userDao.pool.Get()
	defer conn.Close()

	user, err = userDao.getUserById(conn, userId) // 验证用户ID是否存在
	if err != nil {
		// fmt.Println("userDao.getUserById err:", err)
		return
	}
	// 表示用户存在，接着验证密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (userDao *UserDao) Register(user *message.User) (err error) {
	conn := userDao.pool.Get()
	defer conn.Close()

	_, err = userDao.getUserById(conn, user.UserId) // 验证用户ID是否存在
	if err == nil {                                 // 没有错误, 说明取到了UserId, 说明用户是存在的, 不能注册
		err = ERROR_USER_EXISTS
		return
	}

	// 说明redis中没有UserId, 可以注册, 将user序列化以存储到数据库中
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("register json.marshal err:", err)
		return
	}
	// fmt.Println("可以注册咯~")

	_, err = conn.Do("hSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("reserve to redis err:", err)
		return
	}
	return
}

func (userDao *UserDao) name() {

}
