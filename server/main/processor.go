package main

import (
	"basic/chatroom/server/processor"
	"basic/chatroom/server/utils"
	"basic/chatroom/util/message"
	"errors"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMessage(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMessageType:
		up := &processor.UserProcessor{Conn: this.Conn}
		err = up.ServerProcessLogin(mes)

	case message.RegisterMessageType:
		up := &processor.UserProcessor{Conn: this.Conn}
		err = up.ServerProcessRegister(mes)

	default:
		fmt.Println("message type is missing...")
	}

	return err
}

func (this *Processor) dispatch() (err error) {
	fmt.Println("start a process...")
	for {
		tf := utils.Transfer{Conn: this.Conn}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				err = errors.New("client exit")
			} else {
				err = errors.New("process read pkg err")
			}
			return err
		}

		err = this.serverProcessMessage(&mes)
		if err != nil {
			return err
		}

		fmt.Println("mes=", mes)
	}
}
