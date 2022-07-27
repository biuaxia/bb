package web

import (
	"encoding/json"
	"net/http"

	"biuaxia.cn/bb/code/core"
	"biuaxia.cn/bb/code/service"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
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

func Template(c *gin.Context, name string, data any) {
	ginview.HTML(c, http.StatusOK, "dashboard", data)
}

func Redirect(c *gin.Context, location string) {
	c.Redirect(http.StatusFound, location)
}

type TemplateOptions struct {
	OutLastContent bool // 最新文章
	OutConf        bool // 配置内容
}

func (t TemplateOptions) Template(c *gin.Context, name string, data gin.H) {
	marshal, _ := json.Marshal(data)
	var res map[string]interface{}
	_ = json.Unmarshal(marshal, &res)

	switch {
	case t.OutLastContent:
		contents := service.GetAllContentOmitText()
		res["lastContent"] = contents
	case t.OutConf:
		marshal, _ := json.MarshalIndent(core.Conf, "", "    ")
		res["conf"] = string(marshal)
	}

	h := gin.H{}
	for k, v := range res {
		h[k] = v
	}

	ginview.HTML(c, http.StatusOK, name, h)
}
