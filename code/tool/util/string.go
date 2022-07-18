package util

import "fmt"

func GetMysqlUrl(
	mysqlPort int,
	mysqlHost string,
	mysqlSchema string,
	mysqlUsername string,
	mysqlPassword string,
	mysqlCharset string) string {

	if mysqlCharset == "" {
		mysqlCharset = "utf8"
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlSchema, mysqlCharset)
}
