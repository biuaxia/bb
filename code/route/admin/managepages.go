package admin

import (
	"net/http"

	"biuaxia.cn/bb/code/service"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func ManagePages(ctx *gin.Context) {
	contents := service.GetAllContentOmitText("page")
	ginview.HTML(ctx, http.StatusOK, "manage-pages", gin.H{
		"contents": contents,
	})
}
