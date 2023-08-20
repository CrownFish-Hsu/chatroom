package processor

import (
	"basic/chatroom/client/utils"
	"basic/chatroom/util/message"
	"encoding/json"
	"fmt"
)

type SmsProcessor struct {
}

func (this *SmsProcessor) sendGroupMessage(content string) (err error) {
	//1. 创建消息结构体
	var mes message.Message
	mes.Type = message.SmsMessageType

	//2. 创建smsMessage消息实例
	var smsMessage message.SmsMessage
	smsMessage.Content = content
	smsMessage.UserName = curUser.UserName
	smsMessage.UserStatus = curUser.UserStatus

	//3.smsMessage序列化
	data, err := json.Marshal(smsMessage)
	if err != nil {
		fmt.Println("sendGroupMessage smsMessage Marshal err = ", err.Error())
		return err
	}

	mes.Data = string(data)
	//4.sms再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("sendGroupMessage mes Marshal err = ", err.Error())
		return err
	}

	//5. 将mes发送给server
	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("sendGroupMessage WritePkg err = ", err.Error())
		return err
	}

	return
}
