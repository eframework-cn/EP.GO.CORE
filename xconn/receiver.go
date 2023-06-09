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
	_ "io"
	"net"

	_ "github.com/eframework-cn/EP.GO.CORE/xproto"
	_ "github.com/eframework-cn/EP.GO.UTIL/xmath"
)

const (
	RECEIVER_CHAN_LENGTH = 512 // 接收通道长度
)

var (
	ERR_RECEIVER_CHAN_FULL          = errors.New("receiver chan is full")      // 错误：接收通道已满
	ERR_RECEIVER_INVALID_HEAD       = errors.New("receiver head is invalid")   // 错误：消息头部不合法
	ERR_RECEIVER_LARGE_BODY_SIZE    = errors.New("receiver body is too large") // 错误：消息体过长
	ERR_RECEIVER_NEGATIVE_BODY_SIZE = errors.New("receiver body is negative")  // 错误：消息体不合法
)

// 消息接收器
type Receiver struct {
	ChMsg   chan []byte // 消息通道
	MaxBody int         // 最大消息体
}

// 新建消息接收器
//	maxBoxy: 最大消息体
func NewReceiver(maxBody int) *Receiver {
	return nil
}

// 接收消息
func (this *Receiver) Recv(c net.Conn) (int64, error) {
	return 0, nil
}

// 清空对象
func (this *Receiver) Clear() {
	return
}
