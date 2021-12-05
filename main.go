package main

import (
	"log"
	"net/http"
	"time"

	"gitee.com/gjianbo/web/global"
	"gitee.com/gjianbo/web/internal/routers"
	"gitee.com/gjianbo/web/pkg/setting"
	"github.com/gin-gonic/gin"
)

func SetupSetting() error {

	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServetSettingS)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSettingS)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSettings)
	if err != nil {
		return err
	}

	global.ServetSettingS.ReadTimeout *= time.Second
	global.ServetSettingS.WriteTimtout *= time.Second

	return nil
}
func init() {

	err := SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}

}
func main() {

	gin.SetMode(global.ServetSettingS.RunMode)

	routrer := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServetSettingS.HttpPort,
		Handler:        routrer,
		ReadTimeout:    global.ServetSettingS.ReadTimeout,
		WriteTimeout:   global.ServetSettingS.WriteTimtout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
	/* 	r := gin.Default()
	   	r.GET("/ping", func(c *gin.Context) {
	   		c.JSON(200, gin.H{"message": "pong"})
	   	})
	   	r.Run() */

}
