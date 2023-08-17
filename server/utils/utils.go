package utils

import (
	"basic/chatroom/util/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// 将方法绑定进结构体
type Transfer struct {
	Conn net.Conn
	Buff [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//等待client发送信息，如果client没有write，则协程一直阻塞
	fmt.Println("ReadPkg waiting client write ip=", this.Conn.RemoteAddr().String())

	_, err = this.Conn.Read(this.Buff[:4])
	if err != nil {
		return
	}

	//要截取:n

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buff[:4])
	n, err := this.Conn.Read(this.Buff[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("ReadPkg n != pkgLen, n=", n, ", pkgLen=", pkgLen, ", err=", err)
		return
	}

	//pkgLen 反序列化 message.Message
	err = json.Unmarshal(this.Buff[:pkgLen], &mes)
	if err != nil {
		fmt.Println("ReadPkg unmarshal error, err=", err)
		return
	}

	fmt.Println("ReadPkg n = pkgLen, n=", n, ", pkgLen=", pkgLen, ", err=", err)
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送长度给对方
	var msgLen uint32
	msgLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[:4], msgLen)
	//send length
	n, err := this.Conn.Write(bytes[:4])
	if n != 4 || err != nil {
		fmt.Println("WritePkg length error, n=", n, ", error = ", err)
		return err
	}

	//发送消息
	n, err = this.Conn.Write(data)
	if n != int(msgLen) || err != nil {
		fmt.Println("WritePkg data error, n=", n, ", error = ", err)
		return err
	}

	fmt.Println("WritePkg data=", string(data))
	return nil
}
