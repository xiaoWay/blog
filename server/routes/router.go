package routes

import (
	"github.com/xiaoWay/blog/config"
	"github.com/xiaoWay/blog/dao"
	"github.com/xiaoWay/blog/utils"
	"log"
	"net/http"
	"time"
)

// 初始化全局变量

func InitGlobalVariable() {
	// 初始化 Viper
	utils.InitViper()
	// 初始化 Logger
	utils.InitLogger()
	// 初始化数据库 DB
	dao.DB = utils.InitMySQLDB()
	// 初始化 Redis
	utils.InitRedis()
	// 初始化 Casbin
	utils.InitCasbin(dao.DB)
}

// 后台服务
func BackendServer() *http.Server {
	backPort := config.Cfg.Server.BackPort
	log.Printf("后台服务启动于 %s 端口", backPort)
	return &http.Server{
		Addr:         backPort,
		Handler:      BackRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

// 前台服务
