package route

import (
	"html/template"

	"biuaxia.cn/bb/code/route/admin"
	"biuaxia.cn/bb/code/route/index"
	"biuaxia.cn/bb/code/route/openapi"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func IndexGroup(router *gin.Engine) {
	indexMiddleware := ginview.NewMiddleware(goview.Config{
		Root:         "build/html/templates/frontend", // 前端模板路径
		Extension:    ".html",                         // 模板文件后缀
		Master:       "layouts/master",                // 渲染主文件
		Funcs:        template.FuncMap{},              // 自定义函数
		DisableCache: true,                            // 禁用缓存
	}) // 前端默认渲染
	indexGroup := router.Group("/", indexMiddleware)

	indexGroup.GET("/", index.Index)
}

func AdminGroup(router *gin.Engine) {
	adminMiddleware := ginview.NewMiddleware(goview.Config{
		Root:         "build/html/templates/backend", // 后端模板路径
		Extension:    ".html",                        // 模板文件后缀
		Master:       "layouts/master",               // 渲染主文件
		Funcs:        adminFuncMap,                   // 自定义函数
		DisableCache: true,                           // 禁用缓存
	}) // 后端默认渲染
	adminGroup := router.Group("/admin", adminMiddleware)

	{
		adminGroup.GET("/", admin.Index)
		adminGroup.GET("/index.php", admin.Index)

		adminGroup.GET("/dashboard", admin.Dashboard)
		adminGroup.GET("/dashboard.php", admin.Dashboard)

		adminGroup.GET("/write-post", admin.WritePost)
		adminGroup.GET("/write-post.php", admin.WritePost)

		adminGroup.GET("/manage-posts", admin.ManagePosts)
		adminGroup.GET("/manage-posts.php", admin.ManagePosts)
	}
}

func OpenapiGroup(router *gin.Engine) {
	openapiMiddleware := ginview.NewMiddleware(goview.Config{
		Root:         "build/html/templates/openapi", // 前端模板路径
		Extension:    ".html",                        // 模板文件后缀
		Master:       "layouts/master",               // 渲染主文件
		Partials:     []string{"partials/ad"},        // 局部模板?
		Funcs:        openapiFuncMap,                 // 自定义函数
		DisableCache: true,                           // 禁用缓存
	}) // 前端默认渲染
	openapiGroup := router.Group("/openapi", openapiMiddleware)

	openapiGroup.GET("/", openapi.Index)
}
