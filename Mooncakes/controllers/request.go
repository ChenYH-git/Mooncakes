package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrorChangeInt    = errors.New("强制转换为int失败")
	ErrorUserNotLogin = errors.New("用户未登录")
)

const CtxUserIDKey = "userID"

func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorChangeInt
		return
	}
	return
}
