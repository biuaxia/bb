package support

import (
	"html/template"
	"reflect"

	"biuaxia.cn/bb/code/core"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BbRouter struct {
	Engine        *gin.Engine
	routeMap      map[string]gin.HandlerFunc
	adminRouteMap map[string]gin.HandlerFunc
}

func registerGin(bbr *BbRouter) {
	router := gin.Default()
	// 静态资源服务
	router.Static("/assets", "build/html/assets")

	// bbc.Router
	router.HTMLRender = ginview.New(goview.Config{
		Root:         "build/html/templates/frontend", // 前端模板路径
		Extension:    ".html",                         // 模板文件后缀
		Master:       "layouts/master",                // 渲染主文件
		Funcs:        template.FuncMap{},              // 自定义函数
		DisableCache: true,                            // 禁用缓存
	}) // 前端默认渲染
	for k, v := range bbr.routeMap {
		router.GET(k, v)
	}

	adminMiddleware := ginview.NewMiddleware(goview.Config{
		Root:         "build/html/templates/backend", // 后端模板路径
		Extension:    ".html",                        // 模板文件后缀
		Master:       "layouts/master",               // 渲染主文件
		Funcs:        template.FuncMap{},             // 自定义函数
		DisableCache: true,                           // 禁用缓存
	}) // 后端默认渲染
	adminGroup := router.Group("/admin", adminMiddleware)
	for k, v := range bbr.adminRouteMap {
		adminGroup.GET(k, v)
	}

	bbr.Engine = router
}

func registerRoutes() *BbRouter {
	router := &BbRouter{
		routeMap:      make(map[string]gin.HandlerFunc),
		adminRouteMap: make(map[string]gin.HandlerFunc),
	}

	for _, controller := range core.CONTEXT.GetControllerMap() {
		routes := controller.RegisterRoutes()

		typeOf := reflect.TypeOf(controller)
		typeName := typeOf.String()

		if typeName == "*admin.AdminController" {
			routes := controller.RegisterRoutes()
			for k, v := range routes {
				router.adminRouteMap[k] = v
				zap.L().Info("admin RegisterRoutes success",
					zap.String("controller", typeName),
					zap.String("router", k))
			}
		} else {
			for k, v := range routes {
				router.routeMap[k] = v
				zap.L().Info("RegisterRoutes success",
					zap.String("controller", typeName),
					zap.String("router", k))
			}
		}
	}

	zap.L().Info("registerRoutes success")

	return router
}
