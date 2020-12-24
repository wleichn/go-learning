package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

// iServer的实现，定义一个Server的服务器模块
type Server struct{
	// 服务器名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的ip
	IP string
	// 服务器监听端口
	Port int
}
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP:%s, Port %d, is Starting\n", s.IP, s.Port)

	go func(){
		// 1 获取一个TCP的套接字
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("ResolveTCPAddr err ", err)
			return
		}

		// 2 监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, "err ", err)
		}


		// 3 阻塞的等待客户端连接，处理客户端连接业务
		for {
			// 如果客户端有链接过来， 阻塞返回
			conn, err := listener.AcceptTCP()
			if err != nil{
				fmt.Println("Accept err ", err)
			}

			// 已经与客户端建立连接， 实现业务
			go func(){
				for{
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil{
						fmt.Println("conn read err ", err)
						return
					}

					// 回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("Write buf back err ", err)
						return
					}
				}
			}()
		}
	}()

}

func (s *Server) Stop() {
	// TODO 将服务器资源、状态或者一些已经开辟的连接信息停止回收
}

func (s *Server) Serve() {
	// 启动server的服务功能
	s.Start()

	// TODO 做一些启动服务器后的额外业务  框架扩展

	//阻塞状态
	select{}
}
/*
初始化服务器的方法
 */
func NewServer(name string) ziface.IServer{
	s := &Server{
		Name : name,
		IPVersion : "tcp4",
		IP: "0.0.0.0",
		Port: 8999,
	}

	return s
}
