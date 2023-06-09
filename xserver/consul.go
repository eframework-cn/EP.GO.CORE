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
	_ "encoding/json"
	_ "errors"
	_ "fmt"
	_ "io/ioutil"
	_ "net/http"
	_ "strings"
	_ "time"

	_ "github.com/eframework-cn/EP.GO.UTIL/xjson"
	_ "github.com/eframework-cn/EP.GO.UTIL/xlog"
	_ "github.com/eframework-cn/EP.GO.UTIL/xrun"
	_ "github.com/eframework-cn/EP.GO.UTIL/xstring"
	consulapi "github.com/hashicorp/consul/api"
)

const (
	CONSUL_RESP_OK    = "ok"     // Consul响应200
	CONSUL_CHECK_PATH = "/check" // Consul心跳检测
)

var (
	CslClt *consulapi.Client // Consul连接
	CslCfg *consulapi.Config // Consul配置
)

func startConsul(onUpdate func(map[string][]string)) {
	return
}
func httpConsul(addr string) {
	return
}
func regConsul(name string, nodeID string, addr string, port int, regAddr string, timeout string, interval string, unregAfter string) error {
	return nil
}
func unregConsul(nodeID string) error {
	return nil
}
func pullConsul(url string, servers string, onUpdate func(map[string][]string)) {
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
