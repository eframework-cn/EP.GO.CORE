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
	_ "math/rand"
	"time"

	_ "github.com/eframework-cn/EP.GO.UTIL/xlog"
)

const (
	SERVER_SLEEP time.Duration = 10 * time.Millisecond // 帧刷新间隔
)

var (
	GWrap   *Wrap   // 服务封装器
	GServer IServer // 全局服务
)

// 服务封装器
type Wrap struct {
	Svr    IServer
	ChQuit chan bool // 阻塞chan
}

// 新建服务封装器
//	server: 服务对象
func NewWrap(server IServer) *Wrap {
	return nil
}

// 初始化
func (this *Wrap) Init() bool {
	return false
}

// 运行
func (this *Wrap) Run() {
	return
}

// 销毁
func (this *Wrap) Destroy() {
	return
}

// 停止
func (this *Wrap) Stop() {
	return
}

// 启动
//	server: 服务对象
func Start(server IServer) {
	return
}

// 停止
func Stop() {
	return
}
