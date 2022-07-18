package rest

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaseBean struct{}

func (bb *BaseBean) Init() {}

func (bb *BaseBean) Bootstrap() {}

func (bb *BaseBean) Cleanup() {}

func (bb *BaseBean) PanicError(err error) {
	zap.L().Error("Panic err", zap.Error(err))
	panic(err)
}

type BaseController struct {
	BaseBean
}

func (bb *BaseController) Init() {
	bb.BaseBean.Init()
}

func (bb *BaseController) RegisterRoutes() map[string]gin.HandlerFunc {
	return make(map[string]gin.HandlerFunc)
}
