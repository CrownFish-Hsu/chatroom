package main

import (
	"basic/chatroom/server/model"
	"fmt"
	"net"
	"time"
)

func process(conn net.Conn) {
	defer conn.Close()

	processor := &Processor{Conn: conn}
	err := processor.dispatch()
	if err != nil {
		fmt.Println("processor dispatcher err = ", err)
	}
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
func main() {
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()

	listen, err := net.Listen("tcp", "0.0.0.0:9999")
	defer listen.Close()
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}

	for {
		fmt.Println("server waiting client to conn")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err=", err)
			return
		}
		fmt.Println("conn=", conn)

		//启动协程 保持客户端通信
		go process(conn)
	}
}
