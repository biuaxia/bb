package main

import (
	"biuaxia.cn/bb/code/initialize"
)

func main() {
	// 1. viper读取配置
	initialize.Configer()
	// 2. 日志门面
	initialize.Logger()
	// 3. gin 集成 zap
	router := initialize.Router()

	router.Run(":6179") // 运行在 6179 端口
}
