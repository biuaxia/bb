package admin

import (
	"net/http"

	"biuaxia.cn/bb/code/service"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func ManagePosts(ctx *gin.Context) {
	contents := service.GetAllContent()
	ginview.HTML(ctx, http.StatusOK, "manage-posts", gin.H{
		"contents": contents,
	})
}
