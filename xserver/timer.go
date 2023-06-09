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
	"sync"
	_ "sync/atomic"
	"time"

	_ "github.com/eframework-cn/EP.GO.CORE/xproto"
	_ "github.com/eframework-cn/EP.GO.CORE/xsession"
	_ "github.com/eframework-cn/EP.GO.UTIL/xlog"
	_ "github.com/eframework-cn/EP.GO.UTIL/xos"
	_ "github.com/eframework-cn/EP.GO.UTIL/xrun"
	_ "github.com/eframework-cn/EP.GO.UTIL/xtime"
)

const (
	UPDATE_SLEEP time.Duration = 10 * time.Millisecond // 帧刷新间隔
)

var (
	MainTID int64 = -1 // 主线程ID
)

// 定时器句柄
type TimerRecord struct {
	Timers   sync.Map // 定时器映射
	TimerID  int64    // 定时器ID
	LastTime int      // 上次时间
}

// 自增ID
func (this *TimerRecord) MaxID() int64 {
	return 0
}

// 定时器对象
type TimerEntity struct {
	ID      int         // 定时器ID
	Func    func()      // 定时器回调
	Time    int         // 定时时间
	RawTime int         // 初始时间
	Repeat  bool        // 循环调用
	Crash   bool        // 是否崩溃
	RW      bool        // 是否读写
	Tag     interface{} // 日志标签
	Log     int         // 日志层级
}

// 设置会话的可读性（默认为可读可写）
func (this *TimerEntity) SetRW(sig bool) *TimerEntity {
	return nil
}

// 设置会话的标签
func (this *TimerEntity) SetTag(tag interface{}) *TimerEntity {
	return nil
}

// 设置会话的日志层级
func (this *TimerEntity) SetLog(log int) *TimerEntity {
	return nil
}
func procStart(proc *Proc) {
	return
}
func procUpdate() {
	return
}
func procStop(proc *Proc) {
	return
}

// 设置超时调用（务必在逻辑线程中调用或指定线程ID）
//	fun: 回调函数
//	timeout: 超时时间（秒）
//	tid: 线程ID
func SetTimeout(fun func(), timeout float32, tid ...int64) *TimerEntity {
	return nil
}

// 取消超时调用（务必在逻辑线程中调用或指定线程ID）
//	id: 定时器ID
//	tid: 线程ID
func ClearTimeout(id int, tid ...int64) {
	return
}

// 设置间歇调用（务必在逻辑线程中调用或指定线程ID）
//	fun: 回调函数
//	interval: 间歇时间（秒）
//	tid: 线程ID
func SetInterval(fun func(), interval float32, tid ...int64) *TimerEntity {
	return nil
}

// 取消间歇调用（务必在逻辑线程中调用或指定线程ID）
//	id: 定时器ID
//	tid: 线程ID
func ClearInterval(id int, tid ...int64) {
	return
}

// 在逻辑主线程中调用（返回的chan可用于阻塞当前线程）
//	fun: 回调函数
func RunInMain(fun func()) chan bool {
	return nil
}

// 在指定逻辑线程中调用（返回的chan可用于阻塞当前线程）
//	tid: 线程ID
//	fun: 回调函数
func RunIn(tid int64, fun func()) chan bool {
	return nil
}

// 在当前逻辑线程中的下一帧调用
//	fun: 回调函数
func RunInNext(fun func()) *TimerEntity {
	return nil
}
