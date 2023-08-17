package message

import "basic/chatroom/server/model"

const (
	LoginMessageType            = "LoginMessage"
	LoginResponseMessageType    = "LoginResponseMessage"
	RegisterMessageType         = "RegisterMessage"
	RegisterResponseMessageType = "RegisterResponseMessage"
)

type LoginMessage struct {
	UserName     string
	UserPassword string
}

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

type LoginResponseMessage struct {
	Code      int      `json:"code"`
	Error     string   `json:"error"`
	UserLists []string `json:"userLists"`
}

// user结构体类型
type RegisterMessage struct {
	User model.User `json:"user"`
}

type RegisterResponseMessage struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
