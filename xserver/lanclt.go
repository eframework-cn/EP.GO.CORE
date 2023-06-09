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
	_ "errors"
	_ "fmt"

	_ "github.com/eframework-cn/EP.GO.UTIL/xlog"
	_ "github.com/eframework-cn/EP.GO.UTIL/xmath"
	"nanomsg.org/go-mangos"
	_ "nanomsg.org/go-mangos/protocol/push"
	_ "nanomsg.org/go-mangos/transport/ipc"
	_ "nanomsg.org/go-mangos/transport/tcp"
)

// 线路连接
type LanClt struct {
	*LanCfg
	Sockets []mangos.Socket // Socket连接
}

// 新建线路连接
func NewLanClt(cfg *LanCfg) *LanClt {
	return nil
}

// 发送数据
//	bytes: 数据
//	idx: 连接索引
func (this *LanClt) Send(bytes []byte, idx int) error {
	return nil
}

// 关闭连接
func (this *LanClt) Close() {
	return
}
