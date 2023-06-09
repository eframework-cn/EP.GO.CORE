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
	"github.com/eframework-cn/EP.GO.CORE/xproto"
	"github.com/eframework-cn/EP.GO.UTIL/xevt"
)

var (
	GMsg = xevt.NewEvtMgr(true)  // Msg消息中心
	GRpc = xevt.NewEvtMgr(false) // Rpc消息中心
	GCgi = xevt.NewEvtMgr(false) // Cgi消息中心
	GEvt = xevt.NewEvtMgr(true)  // Evt消息中心
)

// Msg函数类型
type MsgFunc func(*xproto.MsgReq)

// 处理回调
//	reply: 响应对象
//	param1: 参数1
//	param2: 参数2
func (this MsgFunc) Handle(reply *xevt.EvtReply, param1 interface{}, param2 interface{}) {
	return
}

// 注册Msg消息（用于客户端和服务器之间交互）（全局）
//	id: 消息ID
//	fun: 消息回调
func RegMsg(id int, fun func(*xproto.MsgReq)) int {
	return 0
}

// 注销Msg消息（用于客户端和服务器之间交互）（全局）
//	id: 消息ID
//	hid: 句柄ID
func UnregMsg(id int, hid int) bool {
	return false
}

// 广播Msg消息（用于客户端和服务器之间交互）（全局）
//	id: 消息ID
//	mreq: 消息对象
func NotifyMsg(id int, mreq *xproto.MsgReq) bool {
	return false
}

// Rpc函数类型
type RpcFunc func(rreq *xproto.RpcReq, rresp *xproto.RpcResp)

// 处理回调
//	reply: 响应对象
//	param1: 参数1
//	param2: 参数2
func (this RpcFunc) Handle(reply *xevt.EvtReply, param1 interface{}, param2 interface{}) {
	return
}

// 注册Rpc消息（用于服务器之间交互）（全局）
//	id: 消息ID
//	fun: 消息回调
func RegRpc(id int, fun func(rreq *xproto.RpcReq, rresp *xproto.RpcResp)) int {
	return 0
}

// 注销Rpc消息（用于服务器之间交互）（全局）
//	id: 消息ID
//	hid: 句柄ID
func UnregRpc(id int, hid int) bool {
	return false
}

// 广播Rpc消息（用于服务器之间交互）（全局）
//	id: 消息ID
//	rreq: 消息请求
//	rresp: 消息响应
func NotifyRpc(id int, rreq *xproto.RpcReq, rresp *xproto.RpcResp) bool {
	return false
}

// Cgi函数类型
type CgiFunc func(req *xproto.CgiReq, resp *xproto.CgiResp)

// 处理回调
//	reply: 响应对象
//	param1: 参数1
//	param2: 参数2
func (this CgiFunc) Handle(reply *xevt.EvtReply, req interface{}, resp interface{}) {
	return
}

// 注册Cgi消息（用于客户端和服务器之间交互）（全局）
//	id: 消息ID
//	fun: 消息回调
func RegCgi(id int, fun func(creq *xproto.CgiReq, cresp *xproto.CgiResp)) int {
	return 0
}

// 注销Cgi消息（用于客户端和服务器之间交互）（全局）
//	id: 消息ID
//	hid: 句柄ID
func UnregCgi(id int, hid int) bool {
	return false
}

// 广播Cgi消息（用于客户端和服务器之间交互）（全局）
//	id: 消息ID
//	creq: 消息请求
//	cresp: 消息响应
func NotifyCgi(id int, creq *xproto.CgiReq, cresp *xproto.CgiResp) bool {
	return false
}

// Evt函数类型
type EvtFunc func(interface{})

// 处理回调
//	reply: 响应对象
//	param1: 参数1
//	param2: 参数2
func (this EvtFunc) Handle(reply *xevt.EvtReply, param1 interface{}, param2 interface{}) {
	return
}

// 注册Evt消息（用于服务器内部）（全局）
//	id: 消息ID
//	fun: 消息回调
func RegEvt(id int, fun func(interface{})) int {
	return 0
}

// 注销Evt消息（用于服务器内部）（全局）
//	id: 消息ID
//	hid: 句柄ID
func UnregEvt(id int, hid int) bool {
	return false
}

// 广播Evt消息（用于服务器内部）（全局）
//	id: 消息ID
//	param: 透传参数
func NotifyEvt(id int, param interface{}) bool {
	return false
}
