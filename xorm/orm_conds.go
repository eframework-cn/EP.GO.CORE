//go:binary-only-package

// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xorm

import (
	_ "fmt"
	_ "strconv"
	_ "strings"
)

// ExprSep define the expression separation
const (
	ExprSep = "__"
)

type CondValue struct {
	Exprs  []string      `json:"Exprs,omitempty"`
	Args   []interface{} `json:"Args,omitempty"`
	Cond   *Condition    `json:"Cond,omitempty"`
	IsOr   bool
	IsNot  bool
	IsCond bool
	IsRaw  bool   `json:"-"`
	Sql    string `json:"Sql,omitempty"`
}

// Condition struct.
// work for WHERE conditions.
type Condition struct {
	Params []CondValue `json:"Params,omitempty"`
	Limit  int         `json:"Limit,omitempty"`
}

// 对象条件，示例：xorm.NewCondition().And("a__gt",1).And("b",2).And("c__contains","hello")
//	__gt: 大于
//	__gte: 大于等于
//	__lt: 小于
//	__lte: 小于等于
//	__in: 特定（最多65534个KEY）
//	__exact: 等于
//	__ne: 不等于
//	__contains: 包含
//	__startswith: 以...起始
//	__endswith: 以...结束
//	__isnull: 是否为空
func NewCondition() *Condition {
	return nil
}

// Raw add raw sql to condition
func (c Condition) Raw(expr string, sql string) *Condition {
	return nil
}

// And add expression to condition
func (c Condition) And(expr string, args ...interface{}) *Condition {
	return nil
}

// AndNot add NOT expression to condition
func (c Condition) AndNot(expr string, args ...interface{}) *Condition {
	return nil
}

// AndCond combine a condition to current condition
func (c *Condition) AndCond(cond *Condition) *Condition {
	return nil
}

// AndNotCond combine a AND NOT condition to current condition
func (c *Condition) AndNotCond(cond *Condition) *Condition {
	return nil
}

// Or add OR expression to condition
func (c Condition) Or(expr string, args ...interface{}) *Condition {
	return nil
}

// OrNot add OR NOT expression to condition
func (c Condition) OrNot(expr string, args ...interface{}) *Condition {
	return nil
}

// OrCond combine a OR condition to current condition
func (c *Condition) OrCond(cond *Condition) *Condition {
	return nil
}

// OrNotCond combine a OR NOT condition to current condition
func (c *Condition) OrNotCond(cond *Condition) *Condition {
	return nil
}

// IsEmpty check the condition arguments are empty or not.
func (c *Condition) IsEmpty() bool {
	return false
}

// clone clone a condition
func (c Condition) clone() *Condition {
	return nil
}
func init() {
	return
}

// 条件表达式，示例：xorm.Parse("a > {0} && b == {1} && c contains {2}", 1, 2, "hello")
//	>: 大于
//	>=: 大于等于
//	<: 小于
//	<=: 小于等于
//	=: 等于
//	!=: 不等于
//	contains: 包含
//	startswith: 以...起始
//	endswith: 以...结束
//	isnull: 是否为空
//	limit: 限定
func Parse(expression string, params ...interface{}) *Condition {
	return nil
}
func getLRIndex(expression string) map[int]int {
	return nil
}
func doParse(expression string, lrMap map[int]int, left int, right int, andStr string, notStr string, cond *Condition, params ...interface{}) *Condition {
	return nil
}
func parseArr(expression string, _strs []string, andStr string, cond *Condition, params ...interface{}) (string, string, *Condition) {
	return "", "", nil
}
func subStr(str string, from int, to int) string {
	return ""
}
func toInt(str string) int {
	return 0
}
