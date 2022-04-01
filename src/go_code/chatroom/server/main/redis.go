/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/3/28 23:45
 * @Project_Name : GoLandProjects
 * @File : redis.go
 * @Software :GoLand
 */

package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func initPool(address string, maxIdle int, maxActive int, idleTimeOut time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,     // 最大空闲连接数
		MaxActive:   maxActive,   // 和数据库的最大连接数，0表示不限制
		IdleTimeout: idleTimeOut, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化连接的代码
			return redis.Dial("tcp", address)
		},
	}
}
