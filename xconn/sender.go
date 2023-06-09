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
	"errors"
	"net"
)

const (
	SENDER_CHAN_LENGTH = 512 // 发送通道长度
)

var (
	ERR_SENDER_FULL = errors.New("sender chan is full") // 错误：发送通道已满
)

// 消息发送器
type Sender struct {
	ChMsg  chan []byte // 消息通道
	ChSend chan bool   // 发送标识
}

// 新建消息发送器
func NewSender() *Sender {
	return nil
}

// 发送消息
func (this *Sender) Send(conn net.Conn) error {
	return nil
}

// 写入数据
func (this *Sender) Write(data []byte) error {
	return nil
}

// 清空对象
func (this *Sender) Clear() {
	return
}
