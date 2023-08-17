package processor

import "fmt"

// 服务器只有UserMgr 1个实例

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[string]*UserProcessor
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[string]*UserProcessor, 1024),
	}
}

func (this *UserMgr) addOrEditOnlineUser(up *UserProcessor) {
	this.onlineUsers[up.UserName] = up
}

func (this *UserMgr) deleteOnlineUser(username string) {
	delete(this.onlineUsers, username)
}

func (this *UserMgr) getAllOnlineUsers() map[string]*UserProcessor {
	return this.onlineUsers
}
func (this *UserMgr) getOnlineUserById(username string) (up *UserProcessor, err error) {
	up, ok := this.onlineUsers[username]
	if !ok {
		err = fmt.Errorf("user %s not found", username)
	}

	return
}
