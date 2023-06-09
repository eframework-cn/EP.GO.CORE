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
	"github.com/eframework-cn/EP.GO.CORE/xorm"
)

// 清除数据（会话内同步，远端异步）
//	model: 数据模型
//	_cond: 条件对象
func GClear(model xorm.ITable, _cond ...*xorm.Condition) {
	return
}
