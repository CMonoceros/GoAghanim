package controller

import (
	"cmonoceros.com/GoAghanim/pkg/lib"
	"cmonoceros.com/GoAghanim/pkg/web/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"reflect"
)

func UserLogin(c *gin.Context) {
	var user model.User
	password := c.PostForm("password")
	token := c.Query("api_token")

	db := lib.GetSqlDb()
	defer db.Close()
	isNotFound := db.First(&user, "api_token = ?", token).
		RecordNotFound()
	if isNotFound {
		lib.GetContext(c).PackError(500, "Can't find user")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		lib.Log.Error(err)
		lib.GetContext(c).PackError(500, "Can't find user")
		return
	}

	lib.GetContext(c).PackMap(lib.GetResource(user, provideUserResource))
}

func provideUserResource(res interface{}) interface{} {
	resType := reflect.TypeOf(res)
	reqType := reflect.TypeOf(model.User{})
	if resType != reqType {
		lib.Log.Panic("Require " + reqType.String() + ",but get " + resType.String())
		return nil
	}
	user := res.(model.User)
	return map[string]interface{}{
		"api_token":     user.ApiToken,
		"user_tel":      user.UserTel,
		"user_tel_zone": user.UserTelZone,
		"user_alias":    user.UserAlias,
	}
}
