package main

import "zinx/znet"
/*
基于zinx框架的服务端应用程序
 */
func main(){
	// 创建一个server实例，使用Zinx的api
	s := znet.NewServer("[zinx V0.2]")
	// 启动server
	s.Serve()

}
