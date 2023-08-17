package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{pool}
	return
}

func (this *UserDao) GetUserByName(name string) (user *User, err error) {
	conn := this.pool.Get()
	res, err := redis.String(conn.Do("hget", "users", name))
	if err != nil {
		err = ERROR_USER_NON_EXIST
		return
	}

	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("UnMarshal failed, err = ", err)
		return
	}

	return
}

func (this *UserDao) Login(name string, password string) (user *User, err error) {
	user, err = this.GetUserByName(name)
	if err != nil {
		err = ERROR_USER_NON_EXIST
		return
	}

	if password != user.UserPwd {
		err = ERROR_USER_PWD
		return
	}

	return
}

func (this *UserDao) Register(user *User) (err error) {
	_, err = this.GetUserByName(user.UserName)
	if err == nil {
		err = ERROR_USER_EXISTED
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		err = ERROR_JSON_MARSHAL
		return
	}

	conn := this.pool.Get()
	_, err = conn.Do("hset", "users", user.UserName, string(data))
	if err != nil {
		//err = ERROR_REDIS_SAVE
		return
	}

	return
}
