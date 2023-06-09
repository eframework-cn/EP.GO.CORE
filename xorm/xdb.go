//go:binary-only-package

//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

package xorm

import (
	_ "fmt"
	_ "sync"

	_ "github.com/eframework-cn/EP.GO.UTIL/xconfig"
	_ "github.com/eframework-cn/EP.GO.UTIL/xfs"
	_ "github.com/eframework-cn/EP.GO.UTIL/xlog"
	_ "github.com/eframework-cn/EP.GO.UTIL/xtime"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Configs map[string]*MysqlConfig = make(map[string]*MysqlConfig) // 数据库配置
	IsValid bool                    = false                         // 是否有效
	// ORM对象池
)

// Mysql配置
type MysqlConfig struct {
	Database string `json:"database"` // 数据库名称
	Address  string `json:"address"`  // 数据库地址
	Username string `json:"username"` // 数据库用户名
	Password string `json:"password"` // 数据库密码
	MaxIdle  int    `json:"maxIdle"`  // 最大空闲数
	MaxConn  int    `json:"maxConn"`  // 最大连接数
	CheckSec int    `json:"checkSec"` // PING检测间隔
	CheckAt  int    // 上一次检测时间
}

// DB连接线路
func (this *MysqlConfig) Dsn() string {
	return ""
}

// 初始化Mysql
func InitMysql() error {
	return nil
}

// DB健康检测（ping）
func HeathCheck() error {
	return nil
}

// 从对象池获取Ormer
func GetOrmer() Ormer {
	return nil
}

// 将Ormer返给对象池
func PutOrmer(ormer Ormer) {
	return
}

// 执行SQL语句
//	dbName: 数据库名称
//	sql: 执行语句
//	args: 携带参数
func ExecSQL(dbName string, sql string, args ...interface{}) RawSeter {
	return nil
}
