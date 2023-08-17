package processor

import (
	"basic/chatroom/client/utils"
	"basic/chatroom/util/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

type UserProcessor struct {
}

func (this *UserProcessor) Login(username string, password string) (err error) {
	fmt.Println("username", username)
	fmt.Println("password", password)
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("net dail error = ", err)
		return err
	}
	defer conn.Close()

	var msg message.Message
	var loginMsg message.LoginMessage
	msg.Type = message.LoginMessageType
	loginMsg.UserName = username
	loginMsg.UserPassword = password

	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("login msg json marshal error = ", err)
		return err
	}
	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("msg json marshal error = ", err)
		return err
	}

	var msgLen uint32
	msgLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[:4], msgLen)
	//send length
	n, err := conn.Write(bytes[:4])
	if n != 4 || err != nil {
		fmt.Println("send msg length error, n=", n, ", error = ", err)
		return err
	}
	fmt.Println("send msg len=", msgLen, ", byte=", bytes)

	time.Sleep(time.Second * 3)

	//send data
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("send msg data error, n=", n, ", error = ", err)
		return err
	}

	fmt.Println("send msg data=", string(data))

	//wait read from Server
	tf := &utils.Transfer{Conn: conn}
	mes, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("send msg data error, n=", n, ", error = ", err)
		return err
	}

	var loginRes message.LoginResponseMessage
	err = json.Unmarshal([]byte(mes.Data), &loginRes)
	if loginRes.Code != 200 {
		fmt.Println("json.Unmarshal Error, err=", loginRes.Error)
		return err
	}

	//显示当前用户列表
	for _, v := range loginRes.UserLists {
		fmt.Printf("online username %s\n", v)
	}

	go serverConnect(conn)
	for {
		showMenu()
	}

	return nil
}

func (this *UserProcessor) Register(username string, password string) (err error) {
	//1.server建立连接
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("net dail error = ", err)
		return err
	}
	defer conn.Close()

	//2.准备msg发送到server,定义类型
	var msg message.Message
	msg.Type = message.RegisterMessageType

	//3.创建RegisterMessage类型结构体
	var registerMsg message.RegisterMessage
	registerMsg.User.UserName = username
	registerMsg.User.UserPwd = password

	//4.将registerMsg序列化
	data, err := json.Marshal(registerMsg)
	if err != nil {
		fmt.Println("register msg json marshal error = ", err)
		return err
	}

	//5.把序列化后的data塞进msg
	msg.Data = string(data)

	//6.将msg再序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("register msg json marshal error = ", err)
		return err
	}

	tf := &utils.Transfer{Conn: conn}
	//7.发送data到server, 发送长度和消息
	err = tf.WritePkg(data)
	fmt.Println("send register msg data=", string(data))

	//8.等待server返回结果,mes即为RegisterResponseMessage
	mes, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("send register msg data error, error = ", err)
		return err
	}

	//9.序列化mes的data
	var registerRes message.RegisterResponseMessage
	err = json.Unmarshal([]byte(mes.Data), &registerRes)
	if registerRes.Code != 200 {
		fmt.Println("json.Unmarshal Error, err=", registerRes.Error)
	} else {
		fmt.Println("register success")
	}

	//10.强制退出
	os.Exit(0)

	return
}
