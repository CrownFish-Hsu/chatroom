package processor

import (
	"basic/chatroom/server/model"
	"basic/chatroom/server/utils"
	"basic/chatroom/util/message"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcessor struct {
	Conn     net.Conn
	UserName string
}

func (this *UserProcessor) ServerProcessLogin(mes *message.Message) (err error) {
	//先取出mes.Data,并反序列化
	var loginMes message.LoginMessage
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("serverProcessLogin Unmarshal Failed, err = ", err)
		return
	}

	var resMes message.Message
	resMes.Data = message.LoginResponseMessageType

	//声明一个返回结构体
	var loginResponseMes message.LoginResponseMessage
	user, err := model.MyUserDao.Login(loginMes.UserName, loginMes.UserPassword)
	if err == nil {
		loginResponseMes.Code = 200

		// 用户登录成功，放进全局userMgr
		this.UserName = user.UserName
		userMgr.addOrEditOnlineUser(this)
		this.NotifyOthersOnlineUser(user.UserName)
		//将username放入loginResponseMes UserLists
		for k, _ := range userMgr.onlineUsers {
			loginResponseMes.UserLists = append(loginResponseMes.UserLists, k)
		}
	} else {
		switch err {
		case model.ERROR_USER_NON_EXIST:
			loginResponseMes.Code = 40001
			loginResponseMes.Error = err.Error()
		case model.ERROR_USER_PWD:
			loginResponseMes.Code = 40002
			loginResponseMes.Error = err.Error()
		default:
			loginResponseMes.Code = 40000
			loginResponseMes.Error = err.Error()
		}
	}

	//序列化loginResponseMes
	data, err := json.Marshal(loginResponseMes)
	if err != nil {
		return
	}

	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		return
	}

	tf := &utils.Transfer{Conn: this.Conn}
	err = tf.WritePkg(data)
	return
}

func (this *UserProcessor) ServerProcessRegister(mes *message.Message) (err error) {
	//1.取出mes.Data,并反序列化成registerMes
	var registerMes message.RegisterMessage
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("ServerProcessRegister Unmarshal Failed, err = ", err)
		return
	}

	//2.声明message返回类型, 定义
	var resMes message.Message
	resMes.Type = message.RegisterResponseMessageType

	//3.声明一个LoginResponseMessage类型的返回结构体
	var registerResponseMes message.RegisterResponseMessage

	//4.注册
	err = model.MyUserDao.Register(&registerMes.User)
	if err == nil {
		registerResponseMes.Code = 200
	} else {
		switch err {
		case model.ERROR_USER_EXISTED:
			registerResponseMes.Code = 50001
			registerResponseMes.Error = err.Error()
		case model.ERROR_JSON_MARSHAL:
			registerResponseMes.Code = 10000
			registerResponseMes.Error = err.Error()
		default:
			registerResponseMes.Code = 50000
			registerResponseMes.Error = err.Error()
		}
	}

	//5.序列化registerResponseMes
	data, err := json.Marshal(registerResponseMes)
	if err != nil {
		return
	}

	//6.将data转成字符串再序列化resMes
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		return
	}

	//7.将data发送给client
	tf := &utils.Transfer{Conn: this.Conn}
	err = tf.WritePkg(data)
	return
}

// userId通知其他用户上线
func (this *UserProcessor) NotifyOthersOnlineUser(userName string) {
	for name, up := range userMgr.onlineUsers {
		if name == userName {
			continue
		}

		up.notifyMeToOther(userName)
	}
}

func (this *UserProcessor) notifyMeToOther(userName string) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMessageType

	var notifyMes message.NotifyUserStatusMessage
	notifyMes.UserName = userName
	notifyMes.Status = message.UserOnline
	data, err := json.Marshal(notifyMes)
	if err != nil {
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		return
	}

	tf := &utils.Transfer{Conn: this.Conn}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("UserProcessorError, error=", err)
		return
	}
}
