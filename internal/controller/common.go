package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jellycheng/goenv"
	"github.com/jellycheng/gosupport"
	"gowebunit/pkg/utils"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gosupport.EmptyStruct(),
	})
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gosupport.EmptyStruct(),
	})
}

func Favicon(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func Sleep30(c *gin.Context) {
	gosupport.Sleep(30)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gosupport.EmptyStruct(),
	})
}

func NoRoute(c *gin.Context) {
	utils.SaveGphParams(c.Writer, c.Request, logFile())
	c.String(http.StatusOK, "success")
}

func logFile() string {
	return fmt.Sprintf("%s%s.log", goenv.GetString("LOG_DIR"), goenv.GetString("APP_NAME"))
}
