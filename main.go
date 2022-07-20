package main

import (
	"biuaxia.cn/bb/code/initialize"
	"biuaxia.cn/bb/code/route"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. viper读取配置
	initialize.Configer()
	// 2. 日志门面
	initialize.Logger()

	router := gin.Default()

	// 静态资源服务
	router.Static("/assets", "build/html/assets")

	route.IndexGroup(router)
	route.AdminGroup(router)
	route.OpenapiGroup(router)

	router.Run(":6179") // 运行在 6179 端口
}
