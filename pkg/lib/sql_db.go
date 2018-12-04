package lib

import (
	"cmonoceros.com/GoAghanim/pkg"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetSqlDb() *gorm.DB {
	mysqlConf := pkg.GetDefaultConfig().MySql
	arg := mysqlConf.User + ":" + mysqlConf.Password + "@/" + mysqlConf.DBName + "?"
	for k, v := range mysqlConf.Args {
		arg += k + "=" + v + "&"
	}
	db, err := gorm.Open("mysql", arg)
	if err != nil {
		Log.Panic(err)
	}
	return db
}
