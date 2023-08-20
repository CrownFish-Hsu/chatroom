package processor

import (
	"basic/chatroom/server/utils"
	"basic/chatroom/util/message"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcessor struct {
}

func (this *SmsProcessor) SendGroupMessage(mes *message.Message) (err error) {
	// 取出Message的内容
	var smsMessage message.SmsMessage
	err = json.Unmarshal([]byte(mes.Data), &smsMessage)
	if err != nil {
		fmt.Println("sendGroupMessage Unmarshal mes err = ", err)
		return err
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("sendGroupMessage Marshal mes err = ", err)
		return err
	}

	// 遍历usersOnline
	for name, up := range userMgr.onlineUsers {
		if smsMessage.UserName == name {
			continue
		}

		this.sendMessageToEachUser(data, up.Conn)
	}

	return nil
}

func (this *SmsProcessor) sendMessageToEachUser(data []byte, conn net.Conn) {
	// 创建transfer，发送data
	tf := &utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("sendMessageToEachUser err = ", err.Error())
	}

	return
}
