package admin

import (
	"net/http"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func ManagePosts(ctx *gin.Context) {
	ginview.HTML(ctx, http.StatusOK, "manage-posts", nil)
}
