package models

import (
	"cmonoceros.com/GoAghanim/src/base"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserTel     string `gorm:"type:varchar(32)"`
	UserTelZone int    `gorm:"type:smallint(6)"`
	ApiToken    string `gorm:"type:varchar(60)"`
	UserAlias   string `gorm:"type:varchar(16)"`
	Password    string `gorm:"type:varchar(256)"`
}

func (User) TableName() string {
	return "user"
}

func Auth(token string) (isAccess bool, user User) {
	db := base.GetSqlDb()
	defer db.Close()
	isAccess = !db.First(&user, "api_token = ?", token).
		RecordNotFound()
	return isAccess, user
}
