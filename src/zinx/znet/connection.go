package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

/*
连接模块
**/
type Connection struct{
	// 当前连接的套接字
	Conn *net.TCPConn

	// 当前连接的ID
	ConnID uint32

	// 当前连接的状态
	isClosed bool

	// 当前连接所绑定的业务处理方法
	handleApi ziface.HandleFunc

	// 告知连接已经停止的chanel
	ExitChan chan bool
}

// 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn: conn,
		ConnID: connID,
		handleApi: callback_api,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

// 当前连接的读业务方法
func (c *Connection)StartReader(){
	fmt.Println("Reader Connection is running...")
	defer fmt.Println("connID = ", c.ConnID, "Reader is exit, remote addr is,", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//  读取客户端数据到buf中
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil{
			fmt.Println("rev buff err:", err)
			// 当前包读取失败继续读下一个
			continue
		}

		// 调用当前连接绑定的HandleAPI（）
		if err := c.handleApi(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID", c.ConnID, "handle is error ", err)
			break
		}
	}
}

// 启动连接
func (c *Connection) Start(){
	fmt.Println("Conn Start().. ConnId = ", c.ConnID)

	// 启动从当前连接读数据的业务
	go c.StartReader()

	// TODO 启动从当前连接写数据的业务
}

// 停止连接
func (c *Connection) Stop(){
	fmt.Println("Conn Stop().. ConnID = ", c.ConnID)

	// 如果当前连接已经关闭
	if c.isClosed == true{
		return
	}

	// 关闭sockect连接
	c.Conn.Close()

	// 回收资源
	close(c.ExitChan)
}

// 获取当前连接绑定的socket conn
func (c *Connection) GETTCPConnection() *net.TCPConn{
	return c.Conn
}

// 获取当前连接的id
func (c *Connection) GetConnID() uint32{
	return c.ConnID
}

// 获取远程客户端tcp状态
func (c *Connection) RemoteAddr() net.Addr{
	return c.Conn.RemoteAddr()
}

// 发送数据
func (c *Connection) Send(data []byte) error{
	return nil
}