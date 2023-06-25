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
	_ "io/ioutil"
	_ "net/http"
	_ "sync"
	_ "time"

	_ "github.com/eframework-cn/EP.GO.UTIL/xcollect"
	_ "github.com/eframework-cn/EP.GO.UTIL/xjson"
	_ "github.com/eframework-cn/EP.GO.UTIL/xlog"
	_ "github.com/eframework-cn/EP.GO.UTIL/xrun"
	_ "github.com/eframework-cn/EP.GO.UTIL/xstring"
	consulapi "github.com/hashicorp/consul/api"
)

// 中心配置
type CslCfg struct {
	Ns       string `json:"ns"`                   // 命名空间（必要）
	Addr     string `json:"addr"`                 // 注册地址（必要）
	Logout   int    `json:"logout" default:"60"`  // 超时注销（可选，默认：60，开发阶段应设置较大值避免注销）
	Timeout  int    `json:"timeout" default:"5"`  // 检测超时（可选，默认：5，开发阶段应设置较大值避免注销）
	Interval int    `json:"interval" default:"5"` // 检测间隔（可选，默认：5，开发阶段应设置较大值避免注销）
	Fetch    int    `json:"fetch" default:"5"`    // 拉取间隔（可选，默认：5）
}

var (
	CslClt *consulapi.Client // Consul连接
	// Consul刷新
)

func regConsul() error {
	return nil
}
func unregConsul(id string) error {
	return nil
}
func waitConsul() chan byte {
	return nil
}
func fetchConsul(onUpdate func(map[string][]string)) {
	return
}

// 推送KV（键值对）至Consul Storage
//	key: 键
//	value: 值
//	version: 版本号，以'VERSION_'为前缀
//	block-是否阻塞
func PostKV(key string, value string, version string, block ...bool) bool {
	return false
}

// 从Consul Storage中拉取KV（键值对）（阻塞）
//	key: 键
func PullKV(key string) []byte {
	return []byte{}
}

// 订阅Consul Storage中的KV（键值对）（阻塞）
//	key: 键（注意订阅的Key需要设置版本，以'VERSION_'为前缀）
//	interval: 间歇时间
func SubKV(key string, interval int, onUpdate func(data []byte)) {
	return
}
