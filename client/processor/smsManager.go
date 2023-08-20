package processor

import (
	"basic/chatroom/util/message"
	"encoding/json"
	"fmt"
)

func outputGroupMessage(mes *message.Message) {
	// 1.反序列化
	var smsMessage message.SmsMessage
	err := json.Unmarshal([]byte(mes.Data), &smsMessage)
	if err != nil {
		fmt.Println("outputGroupMessage Unmarshal err=", err.Error())
		return
	}

	// 2.显示信息
	info := fmt.Sprintf("user %s says %s", smsMessage.UserName, smsMessage.Content)
	fmt.Println(info)
	fmt.Println()
}
