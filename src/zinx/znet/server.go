package znet

import "zinx/ziface"

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

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {

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
