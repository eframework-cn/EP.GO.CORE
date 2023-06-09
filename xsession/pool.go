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
)

// 获取性能分析器对象
func _GetProfiler() *ProfilerRecord {
	return nil
}

// 缓存性能分析器对象
func _PutProfiler(obj *ProfilerRecord) {
	return
}
