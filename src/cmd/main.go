package main

import (
	"NewThread/src/configs"
	mapper "NewThread/src/mapper/mysql"
	"NewThread/src/middle"
	route "NewThread/src/routes"
	utils "NewThread/src/utils"

	"github.com/spf13/viper"
)

func main() {

	// 读取配置文件
	if err := configs.InitConfig(); err != nil {
		panic(any("Configs init error: " + err.Error()))
	}

	// 创建数据库连接
	if err := mapper.InitMysql(); err != nil {
		panic(any("MySQL init error: " + err.Error()))
	}

	// 创建OSS连接
	if err := utils.InitOSS(); err != nil {
		panic(any("OSS init error: " + err.Error()))
	}

	// 读取公钥
	if err := utils.ReadKey(); err != nil {
		panic(any("Read publicKey error: " + err.Error()))
	}

	port := viper.GetString("PORT")
	// 路由
	r := route.InitRouter(middle.Cors())
	if err := r.Run(":" + port); err != nil {
		panic(any("[Route Run Error] error: " + err.Error()))
	}

}
