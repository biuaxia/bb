package admin

import (
	"net/http"

	"biuaxia.cn/bb/code/core"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func Dashboard(ctx *gin.Context) {
	ginview.HTML(ctx, http.StatusOK, "dashboard", gin.H{
		"port":     core.Conf.Port,
		"host":     core.Conf.Host,
		"schema":   core.Conf.Schema,
		"charset":  core.Conf.Charset,
		"username": core.Conf.Username,
		"password": core.Conf.Password,
	})
}
