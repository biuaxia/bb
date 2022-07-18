package main

import (
	"html/template"
	"net/http"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 静态资源服务
	router.Static("/assets", "build/html/assets")

	router.HTMLRender = ginview.New(goview.Config{
		Root:         "build/html/templates/frontend", // 前端模板路径
		Extension:    ".html",                         // 模板文件后缀
		Master:       "layouts/master",                // 渲染主文件
		Funcs:        template.FuncMap{},              // 自定义函数
		DisableCache: true,                            // 禁用缓存
	}) // 前端默认渲染
	router.GET("/", func(ctx *gin.Context) {
		ginview.HTML(ctx, http.StatusOK, "index", gin.H{
			"title":   "前端",
			"content": "前端准备好了~",
		})
	})

	adminMiddleware := ginview.NewMiddleware(goview.Config{
		Root:         "build/html/templates/backend", // 后端模板路径
		Extension:    ".html",                        // 模板文件后缀
		Master:       "layouts/master",               // 渲染主文件
		Funcs:        template.FuncMap{},             // 自定义函数
		DisableCache: true,                           // 禁用缓存
	}) // 后端默认渲染
	adminGroup := router.Group("/admin", adminMiddleware)
	adminGroup.GET("/", func(ctx *gin.Context) {
		ginview.HTML(ctx, http.StatusOK, "index", gin.H{
			"title":   "后端",
			"content": "后端准备好了~",
		})
	})
	adminGroup.GET("/write-post.php", func(ctx *gin.Context) {
		ginview.HTML(ctx, http.StatusOK, "write-post", nil)
	})

	router.Run(":6179") // 运行在 6179 端口
}
