package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jellycheng/goenv"
	"github.com/jellycheng/gosupport"
	"gowebunit/internal/dto"
	"gowebunit/pkg/utils"
)

// XimalayaCallback 接收电信通知
func XimalayaCallback(c *gin.Context) {
	utils.SaveGphParams(c.Writer, c.Request, logFile())
	res := ""
	if c.Request.Method == http.MethodPost { //点播通知
		cpOrderId := c.PostForm("cp_order_id")
		res = utils.DianBoResp(0, cpOrderId)
	} else { //包月通知
		res = "success"
	}

	c.String(http.StatusOK, res)
}

// XimaChannelCallback 模拟喜马渠道方接收回调
func XimaChannelCallback(c *gin.Context) {
	utils.SaveGphParams(c.Writer, c.Request, logFile())
	if c.Request.Method == http.MethodPost {
		appsecret := goenv.GetString("xima_dev_appsecret")
		notifyObj := dto.XimaNotifyDto{}
		err := c.ShouldBindBodyWith(&notifyObj, binding.JSON)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "body内容解析错误：" + err.Error(),
				"data": gosupport.EmptyStruct(),
			})
			return
		}
		if notifyObj.CheckSign(appsecret) {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "success",
				"data": gosupport.EmptyStruct(),
			})
		} else {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "签名错误",
				"data": gosupport.EmptyStruct(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "不支持的请求通知方式：" + c.Request.Method,
			"data": gosupport.EmptyStruct(),
		})
	}
}
