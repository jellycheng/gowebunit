package controller

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.Any("/", Index)
	r.GET("/ping", Ping)
	r.GET("/favicon.ico", Favicon)
	r.Any("/sleep30/return", Sleep30)

	// 通用通知回调应答协议
	r.Any("/commonnotify/callback01", CommonNotifySuccess)
	r.Any("/commonnotify/callback02", CommonNotifyFail)

	// 喜马拉雅相关回调
	r.Any("/ximalaya/callback", XimalayaCallback)
	r.POST("/ximachannel/callback", XimaChannelCallback)

	r.NoRoute(NoRoute)
}
