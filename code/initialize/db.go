package initialize

import (
	"fmt"
	"log"
	"os"
	"time"

	"biuaxia.cn/bb/code/core"
	"biuaxia.cn/bb/code/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logs "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Connect2Db() (err error) {
	var loggers logs.Interface
	if core.Conf.Mode == "dev" || core.Conf.Mode == "" {
		// 进入开发模式，日志输出到终端
		loggers = logs.Default.LogMode(logs.Info)
	} else {
		loggers = logs.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
			logs.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  logs.Silent, // 日志级别
				IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,        // 启用彩色打印
			},
		)
	}

	core.GormDB, err = gorm.Open(mysqlDialector(), &gorm.Config{
		Logger: loggers,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}

	db, err := core.GormDB.DB()
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return
}

func MustConnect2Db() {
	err := Connect2Db()
	if err != nil {
		log.Panic(err)
	}
	zap.L().Info("初始化DB完成")
}

func mysqlDialector() gorm.Dialector {
	// username password host port schema charset
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		core.Conf.MySQLConfig.Username,
		core.Conf.MySQLConfig.Password,
		core.Conf.MySQLConfig.Host,
		core.Conf.MySQLConfig.Port,
		core.Conf.MySQLConfig.Schema,
		core.Conf.MySQLConfig.Charset)
	return mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	})
}

// AutoMigrate 检测表结构是否存在，不存在则新建
func AutoMigrate() {
	core.GormDB.AutoMigrate(&model.Content{})
}
