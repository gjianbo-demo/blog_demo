package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"gitee.com/gjianbo/web/global"
	"gitee.com/gjianbo/web/internal/model"
	"gitee.com/gjianbo/web/internal/routers"
	"gitee.com/gjianbo/web/pkg/setting"
	"github.com/gin-gonic/gin"
)

var (
	isVersion    bool
	buildTime    string
	buildVersion string
	gitCommitID  string
)

func setupSetting() error {

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

func setupDBEngine() error {

	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSettings)

	return err
}

func init() {

	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}

	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()

}
func main() {

	if isVersion {
		fmt.Printf("build_time: %s\r\n", buildTime)
		fmt.Printf("build_version: %s\r\n", buildVersion)
		fmt.Printf("git_commit_id: %s\r\n", gitCommitID)
		return
	}
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
