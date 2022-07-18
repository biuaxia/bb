package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WritePostController struct {
	AdminController
}

func (ac *WritePostController) RegisterRoutes() map[string]gin.HandlerFunc {
	routeMap := make(map[string]gin.HandlerFunc)

	routeMap["/write-post"] = ac.WritePostIndex
	routeMap["/write-post.php"] = ac.WritePostIndex

	return routeMap
}

func (ac *AdminController) WritePostIndex(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"content": "writePost",
	})
}
