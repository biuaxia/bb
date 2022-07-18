package admin

import (
	"net/http"

	"biuaxia.cn/bb/code/rest"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	rest.BaseController
}

func (ac *AdminController) RegisterRoutes() map[string]gin.HandlerFunc {
	routeMap := make(map[string]gin.HandlerFunc)

	routeMap["/"] = ac.AdminIndex
	routeMap["/index.php"] = ac.AdminIndex

	return routeMap
}

func (ac *AdminController) AdminIndex(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"content": "adminIndex",
	})
}
