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

// 读取数据
//	priority: 会话内存 > 全局缓存(LCache == true) > Redis(LRedis == true) > DB(LDB == true)
//	notice: 远端读取（LCache == false、该会话首次读取、条件查找），会触发数据同步锁，可能会严重影响效率（等待GDelete操作完成），可预加载（GList）该模型优化之
//	model: 数据模型
//	_cond: 条件对象
func GRead(model xorm.ITable, _cond ...*xorm.Condition) xorm.ITable {
	return nil
}
