package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jellycheng/goenv"
	"github.com/jellycheng/gosupport"
	"gowebunit/dto"
	"gowebunit/utils"
	"net/http"
)

var (
	port    = flag.String("port", "9908", "指定监听端口，默认9908")
	envFile = flag.String("config", "./.env", "指定配置文件")
)

func main() {
	flag.Parse()
	if gosupport.IsFile(*envFile) {
		err := goenv.LoadEnv(*envFile)
		if err != nil {
			fmt.Println("配置文件解析错误", err.Error())
			return
		}
	}
	appPort := goenv.GetString("APP_PORT")
	listenPort := ":9908"
	if appPort != "" {
		listenPort = appPort
	}
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
		utils.SaveGphParams(c.Writer, c.Request, fmt.Sprintf("%s%s.log", goenv.GetString("LOG_DIR"), goenv.GetString("APP_NAME")))
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
		utils.SaveGphParams(c.Writer, c.Request, fmt.Sprintf("%s%s.log", goenv.GetString("LOG_DIR"), goenv.GetString("APP_NAME")))
		if c.Request.Method == http.MethodPost {
			//appid := goenv.GetString("xima_dev_appid")
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
			isOk := false
			isOk = notifyObj.CheckSign(appsecret)
			if isOk {
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

	})

	//404
	r.NoRoute(func(c *gin.Context) {
		utils.SaveGphParams(c.Writer, c.Request, fmt.Sprintf("%s%s.log", goenv.GetString("LOG_DIR"), goenv.GetString("APP_NAME")))
		res := "success"
		c.String(http.StatusOK, res)
	})

	_ = r.Run(listenPort)
}
