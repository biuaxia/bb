package admin

import (
	"net/http"

	"biuaxia.cn/bb/code/service"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	allContentTotal := service.AllContentTotal("post")
	ginview.HTML(ctx, http.StatusOK, "index", gin.H{
		"allContentTotal": allContentTotal,
	})
}
