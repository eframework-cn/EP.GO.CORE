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

package xserver

import (
	_ "fmt"
	_ "net/http"
	_ "net/url"
	_ "time"

	"github.com/eframework-cn/EP.GO.CORE/xproto"
	_ "github.com/eframework-cn/EP.GO.CORE/xsession"
	_ "github.com/eframework-cn/EP.GO.UTIL/xfs"
	_ "github.com/eframework-cn/EP.GO.UTIL/xjson"
	_ "github.com/eframework-cn/EP.GO.UTIL/xlog"
	_ "github.com/eframework-cn/EP.GO.UTIL/xrun"
	_ "github.com/eframework-cn/EP.GO.UTIL/xstring"
	_ "github.com/eframework-cn/EP.GO.UTIL/xtime"
)

const (
	LAN_CTRL_ROOT     = "/"         // 运维帮助
	LAN_CTRL_STATUS   = "/status"   // 线路状态
	LAN_CTRL_HEALTH   = "/health"   // 线路检测
	LAN_CTRL_BACKUP   = "/backup"   // 线路备份
	LAN_CTRL_RESUME   = "/resume"   // 线路恢复
	LAN_CTRL_PAUSE    = "/pause"    // 线路暂停
	LAN_CTRL_CLOSE    = "/close"    // 线路关闭
	LAN_CTRL_DUMP     = "/dump"     // 清空内存
	LAN_CTRL_PRINT    = "/print"    // 内存打印
	LAN_CTRL_CONSOLE  = "/console"  // 线路日志
	LAN_CTRL_FLUSHLOG = "/flushlog" // 清空日志
)

// 启动线路
//	lancfg: 线路配置
//	handleMsg: 消息处理函数
func StartLan(lanCfg *LanCfg, handleMsg func(*xproto.MsgReq),
	handleRpc func(*xproto.RpcReq, *xproto.RpcResp),
	handleCgi func(*xproto.CgiReq, *xproto.CgiResp)) {
	return
}

// 线路接收
func RecvLan() {
	return
}

// 暂停线路
func PauseLan() {
	return
}

// 恢复线路
func ResumeLan() {
	return
}

// 关闭线路
func CloseLan() {
	return
}

// 监控线路
func MonitorLan() {
	return
}

// 线路备份
func BackupLan() int {
	return 0
}

// 线路恢复
func RestoreLan() {
	return
}
