package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jellycheng/goenv"
	"github.com/jellycheng/gosupport"
	"gowebunit/internal/controller"
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
	listenPort := ":9909"
	if appPort != "" {
		listenPort = ":" + appPort
	}
	if *port != "" {
		listenPort = fmt.Sprintf(":%s", *port)
	}

	r := gin.Default()
	controller.RegisterRoutes(r)
	_ = r.Run(listenPort)
}
