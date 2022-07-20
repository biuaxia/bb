package admin

import (
	"net/http"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func WritePost(ctx *gin.Context) {
	// 检测是否有id携带
	ginview.HTML(ctx, http.StatusOK, "write-post", nil)
}
