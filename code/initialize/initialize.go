package initialize

const (
	DEFAULT_CONF_FILE_PATH = "build/conf/bb.yaml"
)

var (
	Conf = new(AppConfig)
)

type AppConfig struct {
	*MySQLConfig `mapstructure:"mysql"`
}

type MySQLConfig struct {
	Port     int
	Host     string
	Schema   string
	Charset  string
	Username string
	Password string
}
