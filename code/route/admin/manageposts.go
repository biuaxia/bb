package admin

import (
	"biuaxia.cn/bb/code/service"
	"biuaxia.cn/bb/code/web"
	"github.com/gin-gonic/gin"
)

func ManagePosts(ctx *gin.Context) {
	contents := service.GetAllContentOmitText("post")
	web.Template(ctx, "manage-posts", gin.H{
		"name":     "管理文章",
		"contents": contents,
	})
}
