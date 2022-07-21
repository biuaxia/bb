package main

import (
	"biuaxia.cn/bb/code/core"
	"biuaxia.cn/bb/code/initialize"
)

func main() {
	// 1. viper读取配置
	initialize.Configer()
	// 2. 日志门面
	initialize.Logger()
	// 3. gin 集成 zap
	initialize.Router()
	// 4. gorm 集成
	initialize.MustConnect2Db()
	// 5. 自动建表
	initialize.AutoMigrate()

	core.Route.Run(":6179") // 运行在 6179 端口
}
