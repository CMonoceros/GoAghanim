package model

import (
	"cmonoceros.com/GoAghanim/pkg/lib"
	"cmonoceros.com/GoAghanim/pkg/web"
	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserTel     string `gorm:"type:varchar(32)"`
	UserTelZone int    `gorm:"type:smallint(6)"`
	APIToken    string `gorm:"type:varchar(60)"`
	UserAlias   string `gorm:"type:varchar(16)"`
	Password    string `gorm:"type:varchar(256)"`
}

// TableName 用户表名
func (User) TableName() string {
	return "user"
}

// UserAuth 验证用户是否存在
func UserAuth(token string) (isAccess bool, user User) {
	db := lib.GetSQLDb(web.ENV)
	defer func() {
		err := db.Close()
		lib.CheckAndThrowError(err, lib.DefaultCode)
	}()
	isAccess = !db.First(&user, "api_token = ?", token).RecordNotFound()
	return isAccess, user
}
