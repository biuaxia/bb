package openapi

import (
	"net/http"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ginview.HTML(ctx, http.StatusOK, "index", gin.H{
		"title":   "openapi",
		"content": "openapi准备好了~",
	})
}
