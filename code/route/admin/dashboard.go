package admin

import (
	"net/http"

	"biuaxia.cn/bb/code/initialize"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func Dashboard(ctx *gin.Context) {
	ginview.HTML(ctx, http.StatusOK, "dashboard", gin.H{
		"port":     initialize.Conf.Port,
		"host":     initialize.Conf.Host,
		"schema":   initialize.Conf.Schema,
		"charset":  initialize.Conf.Charset,
		"username": initialize.Conf.Username,
		"password": initialize.Conf.Password,
	})
}
