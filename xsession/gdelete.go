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

// 删除数据（会话内同步，远端异步）
//	notice: 该操作会维护数据同步锁，条件删除请使用GClear
//	model: 数据模型
func GDelete(model xorm.ITable) {
	return
}
