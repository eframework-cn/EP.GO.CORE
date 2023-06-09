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

// 支持建立Socket、WebSocket两种长链接服务，处理了多种网络异常，应用层开箱即用.
package xconn

import (
	"net"
	"net/http"
	"sync"
	_ "sync/atomic"
	_ "time"

	_ "github.com/eframework-cn/EP.GO.UTIL/xlog"
	_ "github.com/eframework-cn/EP.GO.UTIL/xrun"
	_ "golang.org/x/net/websocket"
)

const (
	MAXID = 999999999 // 最大ID
)

// 网络服务
type Server struct {
	Addr      string       // 前端连接地址
	IsWS      bool         // 是否面向WebSocket连接（默认true）
	Key       string       // WebSocket密钥
	Cert      string       // WebSocket证书
	WSServer  *http.Server // WebSocket服务
	MaxConn   int          // 最大连接数（超过该值将不再接收新的连接）（默认5000）
	MaxLoad   int          // 单连接每秒最大的请求数（超过该值则视为DDOS，主动断开连接）（默认100QPS）
	MaxBody   int          // 单连接最大的消息体（不包含头部）（默认1024 * 1024 = 1048576）
	Timeout   int          // 连接超时时间（超过该值则主动断开连接）（默认15秒）
	ChClose   chan bool    // 服务关闭标识
	Clients   sync.Map     // map[int]IClient
	Pool      sync.Pool    // IClient缓存池
	OnlineNum int64        // 当前在线数量
	// 当前最大ID
	NewClient func() interface{}    // 连接对象构造
	OnAccept  func(IClient)         // 接收到新连接
	OnRemove  func(IClient)         // 移除一个连接
	OnReceive func(IClient, []byte) // 接收连接数据
}

// 新建网络服务
func NewServer() *Server {
	return nil
}

// 前端连接地址
func (this *Server) SetAddr(addr string) *Server {
	return nil
}

// 是否面向WebSocket连接（默认true，若未设置key和cert则使用http监听）
// 生成密钥：openssl genrsa -out server.key 2048
// 生成证书：openssl req -new -x509 -key server.key -out server.crt -days 365
// 参数说明：Common Name = localhost/domain/ip
func (this *Server) SetIsWS(ws bool, key string, cert string) *Server {
	return nil
}

// 最大连接数（超过该值将不再接收新的连接）（默认5000）
func (this *Server) SetMaxConn(maxConn int) *Server {
	return nil
}

// 单连接每秒最大的请求数（超过该值则视为DDOS，主动断开连接）（默认100QPS）
func (this *Server) SetMaxLoad(maxLoad int) *Server {
	return nil
}

// 单连接最大的消息体（不包含头部）（默认1024 * 1024 = 1048576）
func (this *Server) SetMaxBody(maxBody int) *Server {
	return nil
}

// 连接超时时间（超过该值则主动断开连接）（默认15秒）
func (this *Server) SetTimeout(timeout int) *Server {
	return nil
}

// 连接对象构造
func (this *Server) SetNewClient(newClient func() interface{}) *Server {
	return nil
}

// 接收到新连接
func (this *Server) SetOnAccept(onAccept func(IClient)) *Server {
	return nil
}

// 移除一个连接
func (this *Server) SetOnRemove(onRemove func(IClient)) *Server {
	return nil
}

// 接收连接数据
func (this *Server) SetOnReceive(onReceive func(IClient, []byte)) *Server {
	return nil
}

// 启动服务
func (this *Server) Start() *Server {
	return nil
}

// 停止服务
func (this *Server) Stop() {
	return
}

// 自增ID
func (this *Server) MaxID() int {
	return 0
}

// 添加连接
func (this *Server) AddClient(conn net.Conn) {
	return
}

// 移除连接
func (this *Server) DelClient(client IClient) {
	return
}

// 获取连接
func (this *Server) GetClient(id int) IClient {
	return nil
}
