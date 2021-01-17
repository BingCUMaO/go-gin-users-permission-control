package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"bingcu.top/go-gin-users-permission-control/models"
	"bingcu.top/go-gin-users-permission-control/pkg/gredis"
	"bingcu.top/go-gin-users-permission-control/pkg/logging"
	"bingcu.top/go-gin-users-permission-control/pkg/setting"
	"bingcu.top/go-gin-users-permission-control/routers"
	"bingcu.top/go-gin-users-permission-control/pkg/util"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

// @title User Permission Control Gin API
// @version 1.0
// @description An example of gin for users permission control
// @termsOfService https://bingcu.top/go-gin-users-permission-control
// @license.name MIT
// @license.url https://bingcu.top/go-gin-users-permission-control/blob/master/LICENSE
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
