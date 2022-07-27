package admin

import (
	"biuaxia.cn/bb/code/service"
	"biuaxia.cn/bb/code/web"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	allContentTotal := service.AllContentTotal("post")
	web.Template(ctx, "index", gin.H{
		"name":            "网站概要",
		"allContentTotal": allContentTotal,
	})
}
