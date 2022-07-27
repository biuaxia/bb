package index

import (
	"strconv"
	"strings"

	"biuaxia.cn/bb/code/service"
	"biuaxia.cn/bb/code/web"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ViewContentByArchives(c *gin.Context) {
	idStr := c.Param("id")
	zap.L().Debug("ViewContentByArchives", zap.String("idStr", idStr))
	idStr = strings.TrimSuffix(idStr, ".html")
	zap.L().Debug("ViewContentByArchives", zap.String("idStr", idStr))
	id, _ := strconv.Atoi(idStr)
	content := service.ViewContentByArchives(uint(id))

	web.TemplateOptions{
		OutLastContent: true,
		OutPage:        true,
	}.Template(c, "contentview", gin.H{
		"name":    content.Title,
		"content": content,
	})
}

func ViewPage(c *gin.Context) {
	idStr := c.Param("id")
	zap.L().Debug("ViewPage", zap.String("idStr", idStr))
	idStr = strings.TrimSuffix(idStr, ".html")
	zap.L().Debug("ViewPage", zap.String("idStr", idStr))
	id, _ := strconv.Atoi(idStr)
	content := service.ViewContentByArchives(uint(id))

	web.TemplateOptions{
		OutLastContent: true,
		OutPage:        true,
	}.Template(c, "pageview", gin.H{
		"name":    content.Title,
		"content": content,
	})
}
