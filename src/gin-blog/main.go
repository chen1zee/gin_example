package main

import (
	"fmt"
	"gin_example/src/gin-blog/models"
	"gin_example/src/gin-blog/pkg/logging"
	"gin_example/src/gin-blog/pkg/setting"
	"gin_example/src/gin-blog/routers"
	"net/http"
)

// @title gin_example API
// @version 1.0
// @description This is a sample server for gin
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
