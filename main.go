package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jellycheng/gosupport"
	"gowebunit/utils"
	"net/http"
)

var (
	port = flag.String("port", "9908", "指定监听端口，默认9908")
)

func main() {
	flag.Parse()
	listenPort := ":9908"
	if *port != "" {
		listenPort = fmt.Sprintf(":%s", *port)
	}
	r := gin.Default()
	r.Any("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": gosupport.EmptyStruct(),
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": gosupport.EmptyStruct(),
		})
	})

	// 接收电信通知
	r.Any("/ximalaya/callback", func(c *gin.Context) {
		utils.SaveGphParams(c.Writer, c.Request, "./log/gowebunit.log")
		res := ""
		if c.Request.Method == http.MethodPost {
			res = utils.DianBoResp(0, "f8b76524b99511ed8001000000000000")
		} else {
			res = "success"
		}

		c.String(http.StatusOK, res)
	})

	// 模拟喜马渠道方接收回调
	r.POST("/ximachannel/callback", func(c *gin.Context) {
		utils.SaveGphParams(c.Writer, c.Request, "./log/gowebunit.log")

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gosupport.EmptyStruct(),
		})

	})

	//404
	r.NoRoute(func(c *gin.Context) {
		utils.SaveGphParams(c.Writer, c.Request, "./log/gowebunit.log")
		res := "success"
		c.String(http.StatusOK, res)
	})

	_ = r.Run(listenPort)
}
