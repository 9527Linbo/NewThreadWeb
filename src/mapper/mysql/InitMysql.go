package mapper

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitMysql() (err error) {

	// 设置配置文件的名字
	viper.SetConfigName("configs")
	// 设置配置文件的类型
	viper.SetConfigType("yaml")
	// 添加配置文件的路径，指定 目录下寻找
	//viper.AddConfigPath("/usr/local/src")
	viper.AddConfigPath("./src/configs")
	configsErr := viper.ReadInConfig()
	if configsErr != nil {
		panic(any("configsErr: " + configsErr.Error()))
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.url"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"),
	)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
