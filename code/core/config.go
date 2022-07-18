package core

const (
	DEFAULT_SERVER_PORT      = 6179
	DEFAULT_CONFIG_FILE_PATH = "build/config/bb.yaml"
)

type Config interface {
	ServerPort() int
	MysqlUrl() string
}
