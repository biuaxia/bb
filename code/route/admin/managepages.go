package admin

import (
	"biuaxia.cn/bb/code/service"
	"biuaxia.cn/bb/code/web"
	"github.com/gin-gonic/gin"
)

func ManagePages(ctx *gin.Context) {
	contents := service.GetAllContentOmitText("page")
	web.Template(ctx, "manage-pages", gin.H{
		"name":     "管理独立页面",
		"contents": contents,
	})
}
