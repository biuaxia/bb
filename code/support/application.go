package support

import (
	"fmt"

	"biuaxia.cn/bb/code/core"
	"go.uber.org/zap"
)

type BbApplication struct {
	BeanMap       map[string]core.Bean
	ControllerMap map[string]core.Controller
}

func (bba *BbApplication) Start() {
	defer func() {
		if err := recover(); err != nil {
			zap.L().Error("has error", zap.Any("err", err))
		}
	}()

	bba.HandleWeb()
}

func (bba *BbApplication) HandleWeb() {
	bbLogger := &BbLogger{}
	bbLogger.Init()

	bbConfig := &BbConfig{}
	core.CONFIG = bbConfig
	bbConfig.Init()
	// 重载 log
	bbLogger.InitByConfig(bbConfig)

	bbContext := &BbContext{}
	core.CONTEXT = bbContext
	bbContext.Init()
	defer bbContext.Destroy()

	core.CONTEXT.Run(fmt.Sprintf(":%d", core.CONFIG.ServerPort()))
}
