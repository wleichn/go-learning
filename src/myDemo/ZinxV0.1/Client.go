package main

import (
	"fmt"
	"net"
	"time"
)

/*
	模拟客户端
*/
func main(){
	fmt.Println("client start..")

	time.Sleep(1 *time.Second)

	// 直接连接服务器，获得一个conn连接
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err", err)
		return
	}

	for{
		// 连接调用write 写数据
		_, err := conn.Write([]byte("Hello Zinx V0.1.."))
		if err != nil {
			fmt.Println("client write err", err)
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if  err != nil {
			fmt.Println("client read err", err)
			return
		}
		fmt.Printf("server call back: %s, cnt = %d \n", buf, cnt)

		// cpu阻塞
		time.Sleep(1 * time.Second)
	}
}