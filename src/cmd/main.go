package main

import (
	mapper "NewThread/src/mapper/mysql"
	route "NewThread/src/routes"
	utils "NewThread/src/utils"
)

func main() {

	// 创建数据库连接
	if err := mapper.InitMysql(); err != nil {
		panic(any("MySQL init error: " + err.Error()))
	}

	// 创建OSS连接
	if err := utils.InitOSS(); err != nil {
		panic(any("OSS init error: " + err.Error()))
	}

	//路由
	r := route.InitRouter()
	if err := r.Run(); err != nil {
		panic(any("[Route Run Error] error: " + err.Error()))
	}
}
