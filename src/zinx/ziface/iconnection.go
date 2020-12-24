package ziface

import "net"

type IConnection interface{
	// 启动连接
	Start()

	// 停止连接
	Stop()

	// 获取当前连接绑定的socket conn
	GETTCPConnection() *net.TCPConn

	// 获取当前连接的id
	GetConnID() uint32

	// 获取远程客户端tcp状态
	RemoteAddr() net.Addr

	// 发送数据
	Send(data []byte) error
}

// 定义一个处理连接业务的方法

type HandleFunc func(*net.TCPConn, []byte, int) error
