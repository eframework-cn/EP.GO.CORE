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

// 提供了服务注册发现、服务互联互通、线路负载均衡、业务逻辑承载等功能.
package xserver

import (
	_ "fmt"
	_ "sort"
	_ "time"

	_ "github.com/eframework-cn/EP.GO.CORE/xorm"
	"github.com/eframework-cn/EP.GO.CORE/xproto"
	_ "github.com/eframework-cn/EP.GO.CORE/xsession"
	_ "github.com/eframework-cn/EP.GO.UTIL/xfs"
	_ "github.com/eframework-cn/EP.GO.UTIL/xlog"
	"github.com/eframework-cn/EP.GO.UTIL/xobj"
	_ "github.com/eframework-cn/EP.GO.UTIL/xos"
	_ "github.com/eframework-cn/EP.GO.UTIL/xrun"
	_ "github.com/eframework-cn/EP.GO.UTIL/xtime"
)

const (
	EVT_SERVER_STARTED = -1 // 服务就绪（配置就绪 & 日志就绪 & DB就绪 & Redis就绪 & Lan就绪）
	EVT_SERVER_CHANGED = -2 // 服务变更（参数类型：[]interface{}{added map[string][]string, removed map[string][]string}）
	EVT_SERVER_PREQUIT = -3 // 服务即将退出
)

// 服务接口
type IServer interface {
	Init() bool                                         // 初始化
	Start()                                             // 服务启动
	Update(delta float32)                               // 服务循环
	Destroy()                                           // 服务结束
	PreQuit()                                           // 服务即将退出
	Name() string                                       // 服务名称
	InitConfig() bool                                   // 读取配置
	GetConfig() *SvrCfg                                 // 获取配置
	GetFPS() int                                        // 获取帧率
	UpdateTitle() string                                // 更新标题
	GetTitle() string                                   // 获取标题
	RecvMsg(mreq *xproto.MsgReq)                        // 接收Msg消息
	RecvRpc(rreq *xproto.RpcReq, rresp *xproto.RpcResp) // 接收Rpc消息
	RecvCgi(creq *xproto.CgiReq, cresp *xproto.CgiResp) // 接收Cgi消息
}

// 服务对象
type Server struct {
	xobj.OBJECT
	REAL   IServer
	Config *SvrCfg // 配置信息
	FPS    int     // 应用帧率
	Title  string  // 应用标题
}

// 构造函数
func (this *Server) CTOR(CHILD interface{}) {
	return
}

// 初始化
func (_this *Server) Init() bool {
	return false
}

// 服务启动
func (this *Server) Start() { return }
func (this *Server) Update(delta float32) {
	return
}

// 服务结束
func (this *Server) Destroy() { return }
func (_this *Server) PreQuit() {
	return
}

// 服务名称
func (this *Server) Name() string {
	return ""
}

// 读取配置
func (this *Server) InitConfig() bool {
	return false
}

// 获取配置
func (this *Server) GetConfig() *SvrCfg {
	return nil
}

// 获取帧率
func (this *Server) GetFPS() int { return 0 }
func (this *Server) UpdateTitle() string {
	return ""
}

// 获取标题
func (this *Server) GetTitle() string { return "" }
func (this *Server) RecvMsg(mreq *xproto.MsgReq) {
	return
}

// 接收Rpc消息
func (this *Server) RecvRpc(rreq *xproto.RpcReq, rresp *xproto.RpcResp) {
	return
}

// 接收Cgi消息
func (this *Server) RecvCgi(rreq *xproto.CgiReq, rresp *xproto.CgiResp) {
	return
}
