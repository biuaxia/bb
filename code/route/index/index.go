package index

import (
	"biuaxia.cn/bb/code/service"
	"biuaxia.cn/bb/code/web"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	contents := service.GetAllContentOmitText()

	web.TemplateOptions{
		OutLastContent: true,
	}.Template(ctx, "index", gin.H{
		"contents": contents,
	})
}
