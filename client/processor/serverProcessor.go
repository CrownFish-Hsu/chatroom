package processor

import (
	"basic/chatroom/client/utils"
	"basic/chatroom/util/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

// 显示登陆成功后的列表
func showMenu() {
	fmt.Println("==========login success===========")
	fmt.Println("==========1.Online User List======")
	fmt.Println("==========2.Send SMS==============")
	fmt.Println("==========3.SMS List==============")
	fmt.Println("==========4.Exit==================")
	fmt.Println("Please Select 1-4:")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		outputOnlineUser()
		fmt.Println("show user list")
	case 2:
		fmt.Println("send message")
	case 3:
		fmt.Println("list message")
	case 4:
		fmt.Println("exit")
		os.Exit(0)
	default:
		fmt.Println("input error")
	}
}

func serverConnect(conn net.Conn) {
	tf := &utils.Transfer{Conn: conn}
	for {
		fmt.Println("client keep connection server")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("error =", err)
			return
		}

		switch mes.Type {
		case message.NotifyUserStatusMessageType:
			//1. 取出NotifyUserStatusMessage
			var notifyUserStatusMessage message.NotifyUserStatusMessage
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMessage)

			//2. 保存进client user map
			updateUserStatus(&notifyUserStatusMessage)
		default:
			fmt.Println("mes type undefined, type = ", mes.Type)
		}
		fmt.Println("message=", mes)
	}
}
