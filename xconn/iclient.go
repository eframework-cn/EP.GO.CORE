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

package xconn

import (
	"net"
)

const (
	LOGON_WAIT_TIME     = 10   // 等待第一条消息的时间
	READ_WRITE_DEADLINE = 1e10 // 读取/写入的超时时间
)

// 网络连接
type IClient interface {
	// 获取所属服务
	GetServer() *Server
	// 设置所属服务
	SetServer(server *Server)
	// 获取负载服务
	GetConn() net.Conn
	// 设置负载服务
	SetConn(conn net.Conn)
	// 获取连接ID
	GetID() int
	// 设置连接ID
	SetID(id int)
	// 获取超时时间
	GetExpired() int
	// 设置超时时间
	SetExpired(expired int)
	// 获取心跳时间
	GetTick() int
	// 设置心跳时间
	SetTick(tick int)
	// 获取负载数量
	GetLoad() int
	// 设置负载数量
	SetLoad(load int)
	// 获取关闭标识
	GetChClose() chan bool
	// 设置关闭标识
	SetChClose(ch chan bool)
	// 获取消息发送对象
	GetSender() *Sender
	// 设置消息发送对象
	SetSender(sender *Sender)
	// 获取消息接收对象
	GetReceiver() *Receiver
	// 设置消息接收对象
	SetReceiver(receiver *Receiver)
	// 重置
	Reset(id int, conn net.Conn)
	// 清除
	Clear()
	// 写入数据
	Write(bytes []byte) error
	// 踢出
	Kick()
	// 帧刷新
	Update() bool
}
