package admin

import (
	"biuaxia.cn/bb/code/web"
	"github.com/gin-gonic/gin"
)

func Dashboard(ctx *gin.Context) {
	web.TemplateOptions{
		OutConf: true,
	}.Template(ctx, "dashboard", gin.H{
		"name": "仪表盘",
	})
}
