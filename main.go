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
	port    = flag.String("port", "", "指定监听端口")
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
		listenPort = ":" + appPort
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

	r.Any("/sleep30/return", func(c *gin.Context) {
		gosupport.Sleep(30)
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": gosupport.EmptyStruct(),
		})
	})

	// 通用通知回调应答协议-成功
	r.Any("/commonnotify/callback01", func(c *gin.Context) {
		// 写日志文件
		utils.SaveGphParams(c.Writer, c.Request, fmt.Sprintf("%s%s.log", goenv.GetString("LOG_DIR"), goenv.GetString("APP_NAME")))
		// 响应成功
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gosupport.EmptyStruct(),
		})
	})

	// 通用通知回调应答协议-失败
	r.Any("/commonnotify/callback02", func(c *gin.Context) {
		// 写日志文件
		utils.SaveGphParams(c.Writer, c.Request, fmt.Sprintf("%s%s.log", goenv.GetString("LOG_DIR"), goenv.GetString("APP_NAME")))
		// 响应成功
		c.JSON(200, gin.H{
			"code": 100,
			"msg":  "fail，处理失败",
			"data": gosupport.EmptyStruct(),
		})
	})

	// 接收电信通知
	r.Any("/ximalaya/callback", func(c *gin.Context) {
		utils.SaveGphParams(c.Writer, c.Request, fmt.Sprintf("%s%s.log", goenv.GetString("LOG_DIR"), goenv.GetString("APP_NAME")))
		res := ""
		if c.Request.Method == http.MethodPost {//点播通知
			cpOrderId := c.PostForm("cp_order_id")
			res = utils.DianBoResp(0, cpOrderId)
		} else {//包月通知
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
