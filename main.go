package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/changecoder/webservice/common/setting"
	"github.com/changecoder/webservice/dao/db"
	"github.com/changecoder/webservice/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	db.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routers := routers.InitRouter()
	address := fmt.Sprintf(":%d", setting.ServerSetting.SitePort)
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout

	server := &http.Server{
		Addr:         address,
		Handler:      routers,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	log.Printf("[info] start http server listening %s", address)

	server.ListenAndServe()
}
