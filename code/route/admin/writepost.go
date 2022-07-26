package admin

import (
	"net/http"
	"strconv"

	"biuaxia.cn/bb/code/service"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func WritePost(ctx *gin.Context) {
	// 检测是否有id携带
	cid, ok := ctx.GetQuery("cid")
	if !ok || cid == "" {
		ginview.HTML(ctx, http.StatusOK, "write-post", gin.H{
			"hasData": false,
		})
		return
	}

	id, err := strconv.Atoi(cid)
	if err != nil {
		ginview.HTML(ctx, http.StatusOK, "write-post", gin.H{
			"hasData": false,
		})
		return
	}

	content := service.GetContent(uint(id))
	ginview.HTML(ctx, http.StatusOK, "write-post", gin.H{
		"hasData": true,
		"content": content,
	})
	return
}
