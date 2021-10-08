package controllers

import (
	"Mooncakes/logic"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func GetRankHandler(c *gin.Context) {
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("rank: getCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}

	data, err := logic.GetRank()
	if err != nil {
		zap.L().Error("GetRank failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	me, err := logic.GetMyRank(userID)
	if err != nil {
		zap.L().Error("get my rank failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, gin.H{
		"msg":  "success",
		"me":   me,
		"rank": data,
	})

	return
}
