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

// 提供了数据模型（xorm）、内存沙箱（Cache > Redis > DB）、日志输出（GLog...）等功能.
package xsession

import (
	_ "fmt"
	_ "reflect"
	_ "strings"
	_ "sync"
	_ "sync/atomic"

	_ "github.com/eframework-cn/EP.GO.UTIL/xfs"
	_ "github.com/eframework-cn/EP.GO.UTIL/xos"
	_ "github.com/eframework-cn/EP.GO.UTIL/xrun"
	_ "github.com/eframework-cn/EP.GO.UTIL/xtime"
)

// 开始会话
//	tag: 会话标签
//	log: 日志层级（参考xlog的LogLevel）
//	rw: 是否读写
//	return: 返回会话id，使用该id结束会话
func GStart(tag interface{}, log int, rw bool) int64 {
	return 0
}

// 完成会话（写入/更新/删除内存，性能分析等）
//	id: 会话ID
func GFinish(id int64) {
	return
}

// Dump内存
//	reset: 是否清除
func GDump(reset bool) string {
	return ""
}
