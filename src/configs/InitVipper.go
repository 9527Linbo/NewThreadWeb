package configs

import (
	"github.com/spf13/viper"
)

func InitConfig() error {
	// 设置配置文件的名字
	viper.SetConfigName("configs")

	// 设置配置文件的类型
	viper.SetConfigType("yaml")

	//配置文件地址
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
