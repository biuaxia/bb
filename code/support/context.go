package support

import (
	"reflect"

	"biuaxia.cn/bb/code/core"
	"biuaxia.cn/bb/code/rest/admin"
	"biuaxia.cn/bb/code/rest/index"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BbContext struct {
	db            *gorm.DB
	BeanMap       map[string]core.Bean
	ControllerMap map[string]core.Controller
	Router        *BbRouter
}

func (bbc *BbContext) GetDB() *gorm.DB {
	return bbc.db
}

func (bbc *BbContext) Cleanup() {
	for _, bean := range bbc.BeanMap {
		bean.Cleanup()
	}
}

func (bbc *BbContext) GetControllerMap() map[string]core.Controller {
	return bbc.ControllerMap
}

func (bbc *BbContext) Run(addr ...string) {
	bbc.Router.Engine.Run(addr...)
}

func (bbc *BbContext) registerBean(bean core.Bean) {
	typeOf := reflect.TypeOf(bean)
	typeName := typeOf.String()

	if _, ok := bbc.BeanMap[typeName]; ok {
		zap.L().Error("has been registerd, skip", zap.String("typeName", typeName))
	} else {
		bbc.BeanMap[typeName] = bean
		if controller, ok1 := bean.(core.Controller); ok1 {
			bbc.ControllerMap[typeName] = controller
		}
	}
}

func (bbc *BbContext) registerBeans() {
	bbc.registerBean(new(admin.AdminController))
	bbc.registerBean(new(admin.WritePostController))

	bbc.registerBean(new(index.IndexController))

	zap.L().Info("registerBeans success")
}

func (bbc *BbContext) GetBean(bean core.Bean) core.Bean {
	typeOf := reflect.TypeOf(bean)
	typeName := typeOf.String()

	if val, ok := bbc.BeanMap[typeName]; ok {
		return val
	} else {
		zap.L().Error("not registered", zap.String("typeName", typeName))
		return nil
	}
}

func (bbc *BbContext) initBeans() {
	for _, bean := range bbc.BeanMap {
		bean.Init()
	}

	zap.L().Info("initBeans success")
}

func (bbc *BbContext) Init() {
	bbc.BeanMap = make(map[string]core.Bean)
	bbc.ControllerMap = make(map[string]core.Controller)

	bbc.registerBeans()

	bbc.initBeans()

	zap.L().Info("Bean容器加载完成")

	bbc.Router = registerRoutes()

	zap.L().Info("初始化路由完成")

	registerGin(bbc.Router)

	zap.L().Info("注册路由完成")
}

func (bbc *BbContext) OpenDb() {
	var err error = nil

	// bbc.db, err = gorm.Open("mysql", core.CONFIG.MysqlUrl())

	if err != nil {
		zap.L().Error("failed to connect mysql database", zap.Error(err))
		return
	}

	// whether open the db sql log. (only true when debug)
	// bbc.db.LogMode(false)
}

func (bbc *BbContext) CloseDb() {
	if bbc.db != nil {
		sqlDB, err := bbc.db.DB()
		if err != nil {
			zap.L().Error("occur error when closing db", zap.Error(err))
			return
		}
		err = sqlDB.Close()
		if err != nil {
			zap.L().Error("failed to close db", zap.Error(err))
			return
		}
	}
}

func (bbc *BbContext) Destroy() {
	bbc.CloseDb()
}
