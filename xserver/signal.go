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
	_ "os"
	_ "os/signal"
	_ "syscall"

	_ "github.com/eframework-cn/EP.GO.UTIL/xrun"
)

func doWatch(rch chan<- string) {
	return
}
func WatchSignal() <-chan string {
	return nil
}
