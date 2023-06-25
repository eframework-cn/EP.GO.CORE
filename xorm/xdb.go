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

	_ "github.com/eframework-cn/EP.GO.UTIL/xlog"
	_ "github.com/eframework-cn/EP.GO.UTIL/xtime"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Configs map[string]*MysqlCfg = make(map[string]*MysqlCfg) // 数据库配置
	IsValid bool                 = false                      // 是否有效
	// ORM对象池
)

// Mysql配置
type MysqlCfg struct {
	Database string `json:"database"` // 数据库名称
	Address  string `json:"address"`  // 数据库地址
	Username string `json:"username"` // 数据库用户名
	Password string `json:"password"` // 数据库密码
	MaxIdle  int    `json:"maxIdle"`  // 最大空闲数
	MaxConn  int    `json:"maxConn"`  // 最大连接数
	CheckSec int    `json:"checkSec"` // PING检测间隔
	CheckAt  int    `json:"-"`        // 上一次检测时间
}

// DB连接线路
func (this *MysqlCfg) Dsn() string {
	return ""
}

// 初始化Mysql
//	cfgs: 配置
func InitMysql(cfgs map[string]*MysqlCfg) error {
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
