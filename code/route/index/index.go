package index

import (
	"net/http"

	"biuaxia.cn/bb/code/service"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	contents := service.GetAllContentOmitText()
	ginview.HTML(ctx, http.StatusOK, "index", gin.H{
		"contents": contents,
	})
}
