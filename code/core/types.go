package core

type AppConfig struct {
	Mode        string
	Name        string
	Desc        string
	MySQLConfig *MySQLConfig `mapstructure:"mysql"`
	LogConfig   *LogConfig   `mapstructure:"log"`
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

type Menu struct {
	Parent MenuItem
	Child  []MenuItem
}

type MenuItem struct {
	Name, Href string
}
