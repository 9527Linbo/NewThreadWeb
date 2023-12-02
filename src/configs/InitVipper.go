package configs

import (
	"github.com/spf13/viper"
)

func InitConfig() error {
	// 设置配置文件的名字
	viper.SetConfigName("configs")

	// 设置配置文件的类型
	viper.SetConfigType("yaml")

	// 添加配置文件的路径，指定 目录下寻找

	//服务器的配置文件地址
	viper.AddConfigPath("/www/NewThreadBackend")

	//本地的配置文件地址
	//viper.AddConfigPath("./src/configs")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
