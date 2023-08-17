package model

import "errors"

var (
	ERROR_USER_NON_EXIST = errors.New("user not exist")
	ERROR_USER_EXISTED   = errors.New("user exiseted")
	ERROR_USER_PWD       = errors.New("user password error")

	ERROR_JSON_MARSHAL = errors.New("json marshal error")

	ERROR_REDIS_SAVE = errors.New("save redis error")
)
