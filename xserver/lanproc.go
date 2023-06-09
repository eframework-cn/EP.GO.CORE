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
	"sync"
	_ "sync/atomic"
	_ "time"

	"github.com/eframework-cn/EP.GO.CORE/xproto"
	_ "github.com/eframework-cn/EP.GO.UTIL/xlog"
	_ "github.com/eframework-cn/EP.GO.UTIL/xrun"
)

const (
	LAN_CIN_MAX_FRAME  = 50000 // 最大输入网络帧数
	LAN_COUT_MAX_FRAME = 50000 // 最大输出网络帧数
)

var (
	GLan  *LanSvr // 全局线路服务
	GProc []*Proc // 全局业务处理器
)

// 业务处理器
type Proc struct {
	TID   int64              // 线路的GoID
	Num   int                // 线路线程总数
	CIN   chan xproto.IFrame // 输入队列
	COUT  chan xproto.IFrame // 输出队列
	Loop  bool               // 循环标识
	Pause bool               // 暂停标识
	Resp  sync.Map           // map[int64]chan *xproto.RpcReq/*xproto.CgiFrame
	// 自增ID
}

// 新建业务处理器
func NewProc() *Proc {
	return nil
}

// 弹出第一个输入网络帧
func (this *Proc) PopCIN() (xproto.IFrame, bool) {
	return nil, false
}

// 压入一个输入网络帧
//	frame: 网络帧
func (this *Proc) PushCIN(frame xproto.IFrame) bool {
	return false
}

// 自增ID
func (this *Proc) MaxID() int64 {
	return 0
}
func procFrame(goNum int, handleMsg func(*xproto.MsgReq),
	handleRpc func(*xproto.RpcReq, *xproto.RpcResp),
	handleCgi func(*xproto.CgiReq, *xproto.CgiResp)) {
	return
}
