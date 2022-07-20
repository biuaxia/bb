package initialize

const (
	DEFAULT_CONF_FILE_PATH = "build/conf/bb.yaml"
)

var (
	Conf = new(AppConfig)
)

type AppConfig struct {
	*MySQLConfig `mapstructure:"mysql"`
	*LogConfig   `mapstructure:"log"`
}

type LogConfig struct {
	MaxAge     int
	MaxSize    int
	MaxBackups int
	Level      string
	Filename   string
}

type MySQLConfig struct {
	Port     int
	Host     string
	Schema   string
	Charset  string
	Username string
	Password string
}
