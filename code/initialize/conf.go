package initialize

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Configer() {
	viper.SetConfigFile(DEFAULT_CONF_FILE_PATH) // 配置文件(相对项目的路径)，eg: build/conf/bb.yaml
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("viper.ReadInConfig() failed, err: %s", err))
	}
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("viper.Unmarshal() failed, err: %s", err))
	}

	fmt.Printf("读取 %s 配置文件完成\n", DEFAULT_CONF_FILE_PATH)

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生变化...")
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("viper.OnConfigChange -> viper.Unmarshal() failed, err: %s", err))
		}
	})
}
