package initialize

import (
	"html/template"
	"net/http"
	"time"

	"biuaxia.cn/bb/code/action"
	"biuaxia.cn/bb/code/core"
	"biuaxia.cn/bb/code/route/admin"
	"biuaxia.cn/bb/code/route/index"
	"biuaxia.cn/bb/code/route/openapi"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

var (
	adminFuncMap = template.FuncMap{
		"parseDateTime": parseDateTime,
	}
	openapiFuncMap = template.FuncMap{
		"copy": func() string {
			return time.Now().Format("2006")
		},
	}
)

func parseDateTime(t time.Time) string {
	return t.Format(core.LOCALDATETIME_FORMAT_LAYOUT)
}

func Router() {
	core.Route = gin.New()

	core.Route.Use(logger(), recovery(true))

	// 404
	core.Route.NoRoute(func(c *gin.Context) {
		// controller.ResponseError(c, http.StatusNotFound, "not found")
	})

	// 静态资源服务
	core.Route.Static("/assets", "build/html/assets")

	// 心跳检测
	core.Route.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	IndexGroup()
	AdminGroup()
	OpenapiGroup()
	DispatcherGroup()
}

func IndexGroup() {
	indexMiddleware := ginview.NewMiddleware(goview.Config{
		Root:         "build/html/templates/frontend", // 前端模板路径
		Extension:    ".html",                         // 模板文件后缀
		Master:       "layouts/master",                // 渲染主文件
		Funcs:        template.FuncMap{},              // 自定义函数
		DisableCache: true,                            // 禁用缓存
	}) // 前端默认渲染
	indexGroup := core.Route.Group("/", indexMiddleware)

	indexGroup.GET("/", index.Index)
}

func AdminGroup() {
	adminMiddleware := ginview.NewMiddleware(goview.Config{
		Root:         "build/html/templates/backend", // 后端模板路径
		Extension:    ".html",                        // 模板文件后缀
		Master:       "layouts/master",               // 渲染主文件
		Funcs:        adminFuncMap,                   // 自定义函数
		DisableCache: true,                           // 禁用缓存
	}) // 后端默认渲染
	adminGroup := core.Route.Group("/admin", adminMiddleware)

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

func OpenapiGroup() {
	openapiMiddleware := ginview.NewMiddleware(goview.Config{
		Root:         "build/html/templates/openapi", // 前端模板路径
		Extension:    ".html",                        // 模板文件后缀
		Master:       "layouts/master",               // 渲染主文件
		Partials:     []string{"partials/ad"},        // 局部模板?
		Funcs:        openapiFuncMap,                 // 自定义函数
		DisableCache: true,                           // 禁用缓存
	}) // 前端默认渲染
	openapiGroup := core.Route.Group("/openapi", openapiMiddleware)

	openapiGroup.GET("/", openapi.Index)
}

func DispatcherGroup() {
	r := core.Route.Group("/index.php")

	{
		ActionGroup(r)
	}
}

func ActionGroup(rg *gin.RouterGroup) {
	r := rg.Group("/action")

	{
		r.POST("/contents-post-edit", action.ContentsPostEdit)
	}
}