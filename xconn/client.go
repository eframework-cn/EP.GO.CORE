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
	_ "fmt"
	"net"
	_ "sync/atomic"
	_ "time"

	_ "github.com/eframework-cn/EP.GO.UTIL/xlog"
	"github.com/eframework-cn/EP.GO.UTIL/xobj"
	_ "github.com/eframework-cn/EP.GO.UTIL/xrun"
	_ "github.com/eframework-cn/EP.GO.UTIL/xtime"
)

// 网络连接
type Client struct {
	xobj.OBJECT
	Server   *Server   // 所属服务
	Conn     net.Conn  // 负载服务
	ID       int       // 连接ID
	Expired  int       // 超时时间
	Tick     int       // 心跳时间
	Load     int       // 负载数量
	Sender   *Sender   // 消息发送对象
	Receiver *Receiver // 消息接收对象
	ChClose  chan bool // 关闭标识
	LkClose  int32     // 锁定标识
}

// 新建网络连接
//	server: 所属服务
func NewClient(server *Server) *Client {
	return nil
}

// 构造函数
func (this *Client) CTOR(CHILD interface{}) {
	return
}

// 获取所属服务
func (this *Client) GetServer() *Server {
	return nil
}

// 设置所属服务
func (this *Client) SetServer(server *Server) {
	return
}

// 获取负载服务
func (this *Client) GetConn() net.Conn {
	return nil
}

// 设置负载服务
func (this *Client) SetConn(conn net.Conn) {
	return
}

// 获取连接ID
func (this *Client) GetID() int {
	return 0
}

// 设置连接ID
func (this *Client) SetID(id int) {
	return
}

// 获取超时时间
func (this *Client) GetExpired() int {
	return 0
}

// 设置超时时间
func (this *Client) SetExpired(expired int) {
	return
}

// 获取心跳时间
func (this *Client) GetTick() int {
	return 0
}

// 设置心跳时间
func (this *Client) SetTick(tick int) {
	return
}

// 获取负载数量
func (this *Client) GetLoad() int {
	return 0
}

// 设置负载数量
func (this *Client) SetLoad(load int) {
	return
}

// 获取关闭标识
func (this *Client) GetChClose() chan bool {
	return nil
}

// 设置关闭标识
func (this *Client) SetChClose(ch chan bool) {
	return
}

// 获取消息发送对象
func (this *Client) GetSender() *Sender {
	return nil
}

// 设置消息发送对象
func (this *Client) SetSender(sender *Sender) {
	return
}

// 获取消息接收对象
func (this *Client) GetReceiver() *Receiver {
	return nil
}

// 设置消息接收对象
func (this *Client) SetReceiver(receiver *Receiver) {
	return
}

// 重置
func (this *Client) Reset(id int, conn net.Conn) {
	return
}

// 清除
func (this *Client) Clear() {
	return
}

// 写入数据
func (this *Client) Write(bytes []byte) error {
	return nil
}

// 踢出
func (this *Client) Kick() {
	return
}

// 帧刷新
func (this *Client) Update() bool {
	return false
}
