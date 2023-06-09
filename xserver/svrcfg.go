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
	"github.com/eframework-cn/EP.GO.UTIL/xconfig"
	_ "github.com/eframework-cn/EP.GO.UTIL/xlog"
)

// 服务配置
type SvrCfg struct {
	Raw              xconfig.Configer
	Env              string  // 环境标识: 测试，内测，生产
	LanCfg           *LanCfg // 线路配置
	LinkServer       string  // 需要连接的内部服务器
	ConsulAddr       string  // Consul中心地址
	ConsulHttp       string  // Consul检测地址
	ConsulTimeout    string  // Consul超时时间
	ConsulInterval   string  // Consul访问间隔
	ConsulDeregister string  // Consul延迟注销
}

// 初始化
//	config: 配置内容
func (this *SvrCfg) Init(config string) bool {
	return false
}

// 服务ID
func (this *SvrCfg) SvrID() string {
	return ""
}

// 服务名称
func (this *SvrCfg) SvrName() string {
	return ""
}

// 是否调试环境
func (this *SvrCfg) IsDebug() bool {
	return false
}
