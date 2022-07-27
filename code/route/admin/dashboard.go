package admin

import (
	"net/http"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func Dashboard(ctx *gin.Context) {
	ginview.HTML(ctx, http.StatusOK, "dashboard", nil)
}
