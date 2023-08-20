package model

import (
	"basic/chatroom/server/model"
	"net"
)

type CurUser struct {
	Conn net.Conn
	model.User
}
