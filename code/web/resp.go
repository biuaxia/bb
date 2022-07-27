package web

import (
	"encoding/json"
	"net/http"

	"biuaxia.cn/bb/code/core"
	"biuaxia.cn/bb/code/service"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Success(c *gin.Context, code int, msg any, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func Err(c *gin.Context, httpCode int, code int, msg string, jsonStr any) {
	c.JSON(httpCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": jsonStr,
	})
}

func Template(c *gin.Context, name string, data gin.H) {
	// 附加标题、描述
	data["title"] = core.Conf.Name
	data["desc"] = core.Conf.Desc

	ginview.HTML(c, http.StatusOK, name, data)
}

func Redirect(c *gin.Context, location string) {
	c.Redirect(http.StatusFound, location)
}

type TemplateOptions struct {
	OutLastContent bool // 最新文章
	OutPage        bool // 独立页面
	OutConf        bool // 配置内容
}

func (t TemplateOptions) Template(c *gin.Context, name string, data gin.H) {
	if t.OutLastContent {
		contents := service.GetAllContentOmitText("post")
		data["lastContent"] = contents
	}
	if t.OutPage {
		pages := service.GetAllContentOmitText("page")
		data["pages"] = pages
	}
	if t.OutConf {
		marshal, _ := json.MarshalIndent(core.Conf, "", "    ")
		data["conf"] = string(marshal)
	}

	// 附加标题、描述
	data["title"] = core.Conf.Name
	data["desc"] = core.Conf.Desc

	zap.L().Debug("TemplateOptions.Template", zap.Any("data", data))

	ginview.HTML(c, http.StatusOK, name, data)
}
