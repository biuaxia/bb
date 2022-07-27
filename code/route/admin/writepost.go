package admin

import (
	"strconv"

	"biuaxia.cn/bb/code/service"
	"biuaxia.cn/bb/code/web"
	"github.com/gin-gonic/gin"
)

func WritePost(ctx *gin.Context) {
	// 检测是否有id携带
	cid, ok := ctx.GetQuery("cid")
	if !ok || cid == "" {
		web.Template(ctx, "write-post", gin.H{
			"name":    "撰写新文章",
			"hasData": false,
		})
		return
	}

	id, err := strconv.Atoi(cid)
	if err != nil {
		web.Template(ctx, "write-post", gin.H{
			"name":    "撰写新文章",
			"hasData": false,
		})
		return
	}

	content := service.GetContent(uint(id))
	web.Template(ctx, "write-post", gin.H{
		"name":    "编辑文章",
		"hasData": true,
		"content": content,
	})
	return
}
