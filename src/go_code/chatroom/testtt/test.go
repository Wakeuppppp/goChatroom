/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/26 17:13
 * @Project_Name : GoLandProjects
 * @File : test.go.go
 * @Software :GoLand
 */

package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

var pool *redis.Pool

func init1() {
	pool = &redis.Pool{
		MaxIdle:     8,   // 最大空闲连接数
		MaxActive:   0,   // 和数据库的最大连接数，0表示不限制
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化连接的代码
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func add() {
	conn := pool.Get()
	defer conn.Close()
	for i := 1; i <= 5; i++ {
		str := strconv.Itoa(i)
		_, err := conn.Do("set", "0000"+str, "user"+str)
		if err != nil {
			fmt.Printf("第%d个set err: %v\n", i, err)
		}
		fmt.Printf("第%d添加用户成功\n", i)
	}
	// _, err := conn.Do("set", "708966958", "symxmyz2022.")
	// if err != nil {
	// 	fmt.Println("set err:", err)
	// 	return
	// }
	// fmt.Println("添加用户成功")
	//
	// _, err = conn.Do("set", "20220222", "pass2022")
	// if err != nil {
	// 	fmt.Println("set err:", err)
	// 	return
	// }
	// fmt.Println("添加用户成功")
}

func main1() {
	add()
}
