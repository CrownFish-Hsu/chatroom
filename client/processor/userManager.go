package processor

import (
	client_model "basic/chatroom/client/model"
	"basic/chatroom/server/model"
	"basic/chatroom/util/message"
	"fmt"
)

var onlineUsers map[string]*model.User = make(map[string]*model.User, 10)

var curUser client_model.CurUser //用户登录后初始化

func updateUserStatus(message *message.NotifyUserStatusMessage) {
	user, ok := onlineUsers[message.UserName]
	if !ok {
		user = &model.User{UserName: message.UserName}
	}

	user.UserStatus = message.Status
	onlineUsers[message.UserName] = user
}

// 显示客户端在线用户
func outputOnlineUser() {
	for name, _ := range onlineUsers {
		fmt.Println("current online user name = ", name)
	}
}
