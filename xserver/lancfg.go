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
	Name  string   `json:"name"`                 // 线路名称（必要）
	Addr  string   `json:"addr"`                 // 线路地址（可选：固定地址/动态分配，开发阶段应固定端口）
	Ctrl  string   `json:"ctrl"`                 // 运维地址（可选：固定地址/动态分配）
	GO    int      `json:"go" default:"1"`       // 逻辑线程数（可选，默认：1）
	MaxRx int      `json:"maxrx" default:"4096"` // 最大接收字节数（可选，单位：KB，默认：4096）
	Msg   string   `json:"msg" default:"pb"`     // msg消息协议类型（可选：pb/json，默认：pb）
	Cgi   string   `json:"cgi" default:"json"`   // cgi消息协议类型（可选：pb/json，默认：json）
	Link  []string `json:"link"`                 // 关联线路（可选）
}

// 线路解析
//	id: 线路ID（name@ip:port）
//	ip: 线路IP
//	port: 线路端口
func (this *LanCfg) Resolve() (id string, ip string, port int) {
	return id, ip, port
}

// 线路ID（name@ip:port）
func (this *LanCfg) ID() string {
	return ""
}

// 线路IP
func (this *LanCfg) IP() string {
	return ""
}

// 线路端口
func (this *LanCfg) Port() int {
	return 0
}
