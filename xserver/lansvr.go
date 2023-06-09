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
	_ "math/rand"
	"sync"

	_ "github.com/eframework-cn/EP.GO.UTIL/xlog"
	"nanomsg.org/go-mangos"
	_ "nanomsg.org/go-mangos/protocol/pull"
	_ "nanomsg.org/go-mangos/transport/ipc"
	_ "nanomsg.org/go-mangos/transport/tcp"
)

// 线路服务
type LanSvr struct {
	*LanCfg
	mangos.Socket
	Clients  sync.Map // 连接池（map[string][]*LanClt）
	ClientID sync.Map // 连接映射（map[string]*LanClt）
	SClosed  bool     // 是否关闭
}

// 新建线路服务
//	cfg: 线路配置
func NewLanSvr(cfg *LanCfg) *LanSvr {
	return nil
}

// 线路接收
func (this *LanSvr) Recv() ([]byte, error) {
	return []byte{}, nil
}

// 线路关闭
func (this *LanSvr) Close() {
	return
}

// 路由更新
//	smap: 路由表
func (this *LanSvr) Update(smap map[string][]string) {
	return
}

// 选择所有指定类型的线路
//	svr: 服务类型
func (this *LanSvr) SelectAll(svr string) []*LanClt {
	return []*LanClt{}
}

// 随机选择指定类型的线路
//	svr: 服务类型
func (this *LanSvr) SelectRand(svr string) *LanClt {
	return nil
}

// 发送数据
//	svr: 服务类型
//	bytes: 数据
//	idx: 连接索引
func (this *LanSvr) SendData(svr string, bytes []byte, idx int) error {
	return nil
}
