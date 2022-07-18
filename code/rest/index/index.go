package index

import (
	"net/http"

	"biuaxia.cn/bb/code/rest"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	rest.BaseController
}

func (ic *IndexController) RegisterRoutes() map[string]gin.HandlerFunc {
	routeMap := make(map[string]gin.HandlerFunc)

	routeMap["/"] = ic.Index
	routeMap["/index.php"] = ic.Index

	return routeMap
}

func (ic *IndexController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"content": "index",
	})
}
