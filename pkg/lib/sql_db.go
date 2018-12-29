package lib

import (
	"github.com/jinzhu/gorm"
)

// GetSQLDb 获取mysql连接
func GetSQLDb(env Env) *gorm.DB {
	arg := env.Get("MYSQL_USER") + ":" +
		env.Get("MYSQL_PASSWORD") + "@/" +
		env.Get("MYSQL_DB_NAME") + "?" +
		env.Get("MYSQL_DB_ARGS")
	db, err := gorm.Open("mysql", arg)
	CheckAndThrowError(err, LibraryCode)
	return db
}
