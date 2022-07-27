package initialize

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"

	"biuaxia.cn/bb/code/core"
	"biuaxia.cn/bb/code/route/action"
	"biuaxia.cn/bb/code/route/admin"
	"biuaxia.cn/bb/code/route/index"
	"biuaxia.cn/bb/code/route/openapi"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"gitlab.com/golang-commonmark/markdown"
	"go.uber.org/zap"
)

var md = markdown.New(markdown.HTML(true),
	markdown.Tables(true),
	markdown.Linkify(true),
	markdown.Typographer(true),
	markdown.XHTMLOutput(true))

var (
	adminFuncMap = template.FuncMap{
		"md2HTML":                 md2HTML,
		"parseTime":               parseTime,
		"parseDate":               parseDate,
		"parseDateTime":           parseDateTime,
		"parseDateTimeMinute":     parseDateTimeMinute,
		"parseDateTimeMillsecond": parseDateTimeMillsecond,
		"allMenu":                 allMenu,
	}
	indexFuncMap = template.FuncMap{
		"md2HTML":                 md2HTML,
		"parseTime":               parseTime,
		"parseDate":               parseDate,
		"parseDateTime":           parseDateTime,
		"parseDateTimeMinute":     parseDateTimeMinute,
		"parseDateTimeMillsecond": parseDateTimeMillsecond,
		"specifiedStr":            specifiedStr,
		"specifiedStr20":          specifiedStr20,
	}
	openapiFuncMap = template.FuncMap{
		"copy": func() string {
			return time.Now().Format("2006")
		},
	}
)

func md2HTML(s string) template.HTML {
	htmlStr := md.RenderToString([]byte(s))
	zap.L().Debug("md2HTML", zap.String("s", s), zap.String("htmlStr", htmlStr))
	return template.HTML(htmlStr)
}

func specifiedStr(s string, num int) string {
	count := utf8.RuneCountInString(s)
	if count <= num {
		return s
	}
	runes := []rune(s)
	return string(runes[:num])
}

func specifiedStr20(s string) string {
	return specifiedStr(s, 20)
}

func allMenu() template.HTML {
	allMenus := core.AllMenus
	var outStr []string
	for _, menu := range allMenus {
		parent := menu.Parent
		items := menu.Child

		parentHtml := fmt.Sprintf(`<li class="parent"><a href="%s">%s</a></li>`, parent.Href, parent.Name)

		var itemsStr []string
		for _, item := range items {
			itemsStr = append(itemsStr, fmt.Sprintf(`<li><a href="%s">%s</a></li>`, item.Href, item.Name))
		}

		outStr = append(outStr, fmt.Sprintf(`<ul class="root">
            %s
            <ul class="child">
				%s
            </ul>
        </ul>`, parentHtml, strings.Join(itemsStr, "\r\n\t\t\t\t")))
	}
	return template.HTML(strings.Join(outStr, "\r\n\t\t"))
}

func parseTime(ts string) string {
	t := core.ParseLocalDateTime(ts)
	return t.Format(core.LOCALTIME_FORMAT_LAYOUT)
}
func parseDate(ts string) string {
	t := core.ParseLocalDateTime(ts)
	return t.Format(core.LOCALDATE_FORMAT_LAYOUT)
}
func parseDateTime(ts string) string {
	t := core.ParseLocalDateTime(ts)
	return t.Format(core.LOCALDATETIME_FORMAT_LAYOUT)
}
func parseDateTimeMinute(ts string) string {
	t := core.ParseLocalDateTime(ts)
	return t.Format(core.LOCALDATETIME_MINUTE_FORMAT_LAYOUT)
}
func parseDateTimeMillsecond(ts string) string {
	t := core.ParseLocalDateTime(ts)
	return t.Format(core.LOCALDATETIME_MILLSECOND_FORMAT_LAYOUT)
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
	zap.L().Info("初始化路由完成")
}

func IndexGroup() {
	indexMiddleware := ginview.NewMiddleware(goview.Config{
		Root:         "build/html/templates/frontend", // 前端模板路径
		Extension:    ".html",                         // 模板文件后缀
		Master:       "layouts/master",                // 渲染主文件
		Funcs:        indexFuncMap,                    // 自定义函数
		DisableCache: true,                            // 禁用缓存
	}) // 前端默认渲染
	indexGroup := core.Route.Group("/", indexMiddleware)

	indexGroup.GET("/", index.Index)
	indexGroup.GET("/index.php", index.Index)

	indexGroup.GET("/index.php/:id", index.ViewPage)
	indexGroup.GET("/index.php/archives/:id", index.ViewContentByArchives)
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

		adminGroup.GET("/write-page", admin.WritePage)
		adminGroup.GET("/write-page.php", admin.WritePage)

		adminGroup.GET("/manage-posts", admin.ManagePosts)
		adminGroup.GET("/manage-posts.php", admin.ManagePosts)

		adminGroup.GET("/manage-pages", admin.ManagePages)
		adminGroup.GET("/manage-pages.php", admin.ManagePages)
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
		r.POST("/contents-page-edit", action.ContentsPageEdit)
	}
}
