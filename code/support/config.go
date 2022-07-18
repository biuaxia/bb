package support

import (
	"log"

	"biuaxia.cn/bb/code/core"
	"biuaxia.cn/bb/code/tool/util"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type BbConfig struct {
	serverPort int
	mysqlUrl   string
	item       *ConfigItem
}

type ConfigItem struct {
	// server port
	ServerPort int
	// mysql configurations.
	// mysql port
	MysqlPort int
	// mysql host
	MysqlHost string
	// mysql schema
	MysqlSchema string
	// mysql username
	MysqlUsername string
	// mysql password
	MysqlPassword string
	// mysql charset
	MysqlCharset string
	// log config
	// log maxSize
	LogMaxSize int
	// log maxAge
	LogMaxAge int
	// log maxBackups
	LogMaxBackups int
	// log level
	LogLevel string
	// log filename
	LogFilename string
}

type LoggerConfig struct {
	Level      string
	Filename   string
	MaxSize    int `mapstructure:"max_size"`
	MaxAge     int `mapstructure:"max_age"`
	MaxBackups int `mapstructure:"max_backups"`
}

func (bbc *BbConfig) ServerPort() int {
	return bbc.serverPort
}

func (bbc *BbConfig) MysqlUrl() string {
	return bbc.mysqlUrl
}

func (bbc *ConfigItem) validate() bool {
	if bbc.ServerPort == 0 {
		log.Println("ServerPort is not configured")
		return false
	}

	if bbc.MysqlUsername == "" {
		log.Println("MysqlUsername is not configured")
		return false
	}

	if bbc.MysqlPassword == "" {
		log.Println("MysqlPassword is not configured")
		return false
	}

	if bbc.MysqlHost == "" {
		log.Println("MysqlHost is not configured")
		return false
	}

	if bbc.MysqlPort == 0 {
		log.Println("MysqlPort is not configured")
		return false
	}

	if bbc.MysqlSchema == "" {
		log.Println("MysqlSchema is not configured")
		return false
	}
	if bbc.MysqlCharset == "" {
		log.Println("MysqlCharset is not configured")
		return false
	}

	return true
}

func (bbc *BbConfig) Init() {
	bbc.serverPort = core.DEFAULT_SERVER_PORT
	bbc.ReadFromConfigFilePath(core.DEFAULT_CONFIG_FILE_PATH)
}

func (bbc *BbConfig) ReadFromConfigFilePath(filePath string) {
	viper.SetConfigFile(filePath)

	if err := viper.ReadInConfig(); err != nil {
		zap.L().Error("viper.ReadInConfig() failed", zap.Error(err))
		return
	}

	if err := viper.Unmarshal(&bbc.item); err != nil {
		zap.L().Error("viper.Unmarshal() failed", zap.Error(err))
		return
	}

	if bbc.item.ServerPort != 0 {
		bbc.serverPort = bbc.item.ServerPort
	}

	itemValidate := bbc.item.validate()
	if !itemValidate {
		zap.L().Error("config file not integrity, installation will start!")
		return
	}

	bbc.mysqlUrl = util.GetMysqlUrl(bbc.item.MysqlPort,
		bbc.item.MysqlHost,
		bbc.item.MysqlSchema,
		bbc.item.MysqlUsername,
		bbc.item.MysqlPassword,
		bbc.item.MysqlCharset)

	zap.L().Info("读取配置完成", zap.String("filePath", filePath))

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.L().Info("配置文件发生变化...")
		if err := viper.Unmarshal(bbc.item); err != nil {
			zap.L().Error("viper.OnConfigChange -> viper.Unmarshal() failed", zap.Error(err))
			return
		}
	})
}
