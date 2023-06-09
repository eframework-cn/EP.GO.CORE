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
	_ "sync"

	"github.com/eframework-cn/EP.GO.CORE/xorm"
)

// 获取DB指定Column（列）的最大值，并且自增（指定delta值或+1），若未指定Column则自增主键
//	model: 数据模型
//	columnAndDelta(1): column/delta（column需为int类型）
//	columnAndDelta(2): column&delta（column需为int类型）
func GIncrease(model xorm.ITable, columnAndDelta ...interface{}) int {
	return 0
}

// 获取DB指定Column（列）的最大值，若未指定Column则获取主键的最大值
//	notice: 异步写入可能会导致该值不同步
//	column: 列名（column需为int类型）
func GMaxIndex(model xorm.ITable, column ...string) int {
	return 0
}

// 获取DB指定Column（列）的最小值，若未指定Column则获取主键的最小值
//	notice: 异步写入可能会导致该值不同步
//	column: 列名（column需为int类型）
func GMinIndex(model xorm.ITable, column ...string) int {
	return 0
}
