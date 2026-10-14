package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jellycheng/gosupport"
	"gowebunit/pkg/utils"
)

// CommonNotifySuccess 通用通知回调应答协议-成功
func CommonNotifySuccess(c *gin.Context) {
	utils.SaveGphParams(c.Writer, c.Request, logFile())
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": gosupport.EmptyStruct(),
	})
}

// CommonNotifyFail 通用通知回调应答协议-失败
func CommonNotifyFail(c *gin.Context) {
	utils.SaveGphParams(c.Writer, c.Request, logFile())
	c.JSON(200, gin.H{
		"code": 100,
		"msg":  "fail，处理失败",
		"data": gosupport.EmptyStruct(),
	})
}
