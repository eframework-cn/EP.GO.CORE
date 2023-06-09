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

package xsession

import (
	_ "fmt"
	_ "sync"

	"github.com/eframework-cn/EP.GO.CORE/xorm"
	_ "github.com/eframework-cn/EP.GO.UTIL/xlog"
)

const (
	PIPE_CHAN_MAX = 100000 // 推送队列最大长度
)

// map[string]map[string]*GlobalObject
// map[string]bool
// map[string]map[string]int
// map[int]map[string]map[string]*SessionObject
// map[int]map[string]bool
// map[string]*DeleteRecord
// map[int64]int64
// map[int64]*ProfilerRecord
// map[int64]*PipeRecord
// 会话自增ID
// 管道推送计数
// 会话对象维护
type SessionObject struct {
	Raw     xorm.ITable     // 原始对象
	Ptr     xorm.ITable     // 工作对象
	SWrite  bool            // 写入标识
	SDelete bool            // 删除标识
	SClear  bool            // 清除标识
	Cond    *xorm.Condition // 条件对象
}

// 全局对象维护
type GlobalObject struct {
	Ptr     xorm.ITable // 原始对象
	SDelete bool        // 删除标识
}

// 管道对象维护
type PipeObject struct {
	Raw     xorm.ITable     // 原始对象
	SWrite  bool            // 写入标识
	SDelete bool            // 删除标识
	SClear  bool            // 清除标识
	SDB     bool            // 推送至DB
	SRedis  bool            // 推送至Redis
	Cond    *xorm.Condition // 条件对象
}

// 管道推送批次
type PipeRecord struct {
	ID       int64           // 批次ID
	Tag      interface{}     // 批次标签
	Time     int             // 开始时间
	Objects  []*PipeObject   // 管道对象列表
	Profiler *ProfilerRecord // 性能分析器
}

// 删除对象记录
type DeleteRecord struct {
	Chan  chan bool // 通道
	Count int32     // 数量
}

// Defer函数类型
type DeferFun func()

// 性能分析器
type ProfilerRecord struct {
	Tag  interface{} // 会话标签
	Time int         // 开始时间
	RW   bool        // 是否读写
	Log  int         // 日志层级
	TID  int64       // 线程ID
	SID  int64       // 会话ID
	// 日志前缀
}

// 重置对象
func (this *ProfilerRecord) Reset() {
	return
}

// 日志前缀
func (this *ProfilerRecord) LogPrefix() string {
	return ""
}

// 输出会话日志
//	priority: Debug(7) > Info(6) > Notice(5) > Warn(4) > Error(3) > Critical(2) > Alert(1) > Emergency(0)
func (this *ProfilerRecord) GLogEmergency(f interface{}, v ...interface{}) {
	return
}

// 输出会话日志
//	priority: Debug(7) > Info(6) > Notice(5) > Warn(4) > Error(3) > Critical(2) > Alert(1) > Emergency(0)
func (this *ProfilerRecord) GLogAlert(f interface{}, v ...interface{}) {
	return
}

// 输出会话日志
//	priority: Debug(7) > Info(6) > Notice(5) > Warn(4) > Error(3) > Critical(2) > Alert(1) > Emergency(0)
func (this *ProfilerRecord) GLogCritical(f interface{}, v ...interface{}) {
	return
}

// 输出会话日志
//	priority: Debug(7) > Info(6) > Notice(5) > Warn(4) > Error(3) > Critical(2) > Alert(1) > Emergency(0)
func (this *ProfilerRecord) GLogError(f interface{}, v ...interface{}) {
	return
}

// 输出会话日志
//	priority: Debug(7) > Info(6) > Notice(5) > Warn(4) > Error(3) > Critical(2) > Alert(1) > Emergency(0)
func (this *ProfilerRecord) GLogWarn(f interface{}, v ...interface{}) {
	return
}

// 输出会话日志
//	priority: Debug(7) > Info(6) > Notice(5) > Warn(4) > Error(3) > Critical(2) > Alert(1) > Emergency(0)
func (this *ProfilerRecord) GLogNotice(f interface{}, v ...interface{}) {
	return
}

// 输出会话日志
//	priority: Debug(7) > Info(6) > Notice(5) > Warn(4) > Error(3) > Critical(2) > Alert(1) > Emergency(0)
func (this *ProfilerRecord) GLogInfo(f interface{}, v ...interface{}) {
	return
}

// 输出会话日志
//	priority: Debug(7) > Info(6) > Notice(5) > Warn(4) > Error(3) > Critical(2) > Alert(1) > Emergency(0)
func (this *ProfilerRecord) GLogDebug(f interface{}, v ...interface{}) {
	return
}
