package index

import (
	"net/http"
	"strconv"

	"biuaxia.cn/bb/code/service"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func ViewContentByArchives(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	content := service.ViewContentByArchives(uint(id))
	ginview.HTML(c, http.StatusOK, "content/view", gin.H{
		"content": content,
	})
}
