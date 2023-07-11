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
	_ "io/ioutil"
	_ "net"
	_ "net/http"
	_ "os"

	"github.com/eframework-cn/EP.GO.CORE/xorm"
	_ "github.com/eframework-cn/EP.GO.UTIL/xfs"
	_ "github.com/eframework-cn/EP.GO.UTIL/xjson"
	"github.com/eframework-cn/EP.GO.UTIL/xlog"
	_ "github.com/eframework-cn/EP.GO.UTIL/xstring"
)

const (
	ENV_DEV  = "dev"  // 开发
	ENV_TEST = "test" // 内测
	ENV_BETA = "beta" // 公测
	ENV_LIVE = "live" // 线上
)

var (
	GCfg *SvrCfg // 全局配置
)

// 服务配置
type SvrCfg struct {
	Raw   []byte                    `json:"-"`               // 原始数据
	Env   string                    `json:"env"`             // 环境标识（必要：dev[开发]/test[内测]/beta[公测]/live[线上]）
	Csl   *CslCfg                   `json:"csl"`             // 中心配置（必要）
	Lan   *LanCfg                   `json:"lan"`             // 线路配置（必要）
	Log   map[string]*xlog.LogCfg   `json:"log"`             // 日志配置（必要）
	Redis *xorm.RedisCfg            `json:"redis,omitempty"` // Redis配置（可选）
	Mysql map[string]*xorm.MysqlCfg `json:"mysql,omitempty"` // Mysql配置（可选）
}

// 初始化
func (this *SvrCfg) Init() bool {
	return false
}
func allocPort() int {
	return 0
}
