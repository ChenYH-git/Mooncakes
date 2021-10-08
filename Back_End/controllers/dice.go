package controllers

import (
	"Mooncakes/logic"
	"Mooncakes/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func DiceResHandler(c *gin.Context) {
	p := new(models.DiceRes)

	if err := c.Bind(p); err != nil {
		zap.L().Error("Bind DiceRes failed", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("get current userID failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}

	notOK, err := logic.CheckChances(userID)
	if err != nil {
		zap.L().Error("CheckChances failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	if notOK {
		ResponseError(c, CodeNoChances)
		return
	}

	if err := logic.AddScore(userID, p); err != nil {
		zap.L().Error("addscore failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, gin.H{
		"msg": "success",
	})
}
