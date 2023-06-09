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
	_ "fmt"

	_ "github.com/eframework-cn/EP.GO.UTIL/xstring"
)

// 线路配置
type LanCfg struct {
	Name     string // 名称
	Addr     string // tcp://$ip:$port
	Raw      string // $ip:$port
	IP       string // IP
	Port     int    // 端口
	GO       int    // 逻辑线程数
	MaxRx    int    // 最大接收字节数（KB）
	MsgProto string // msg消息协议类型，可选pb/json，默认pb
	CgiProto string // cgi消息协议类型，可选pb/json，默认json
}

// 创建线路配置
//	name: 线路名称
//	addr: 线路地址
func NewLanCfg(name, addr string) *LanCfg {
	return nil
}

// 服务器ID（$name@tcp://$ip:$port）
func (this *LanCfg) ServerID() string {
	return ""
}
