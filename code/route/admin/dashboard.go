package admin

import (
	"net/http"

	"biuaxia.cn/bb/code/core"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func Dashboard(ctx *gin.Context) {
	ginview.HTML(ctx, http.StatusOK, "dashboard", gin.H{
		"port":     core.Conf.MySQLConfig.Port,
		"host":     core.Conf.MySQLConfig.Host,
		"schema":   core.Conf.MySQLConfig.Schema,
		"charset":  core.Conf.MySQLConfig.Charset,
		"username": core.Conf.MySQLConfig.Username,
		"password": core.Conf.MySQLConfig.Password,
	})
}
