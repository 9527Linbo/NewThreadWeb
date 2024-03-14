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
