package index

import (
	"strconv"

	"biuaxia.cn/bb/code/service"
	"biuaxia.cn/bb/code/web"
	"github.com/gin-gonic/gin"
)

func ViewContentByArchives(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	content := service.ViewContentByArchives(uint(id))

	web.TemplateOptions{
		OutLastContent: true,
	}.Template(c, "content/view", gin.H{
		"content": content,
	})
}
