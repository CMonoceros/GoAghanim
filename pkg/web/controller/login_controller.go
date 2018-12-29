package controller

import (
	"cmonoceros.com/GoAghanim/pkg/lib"
	"cmonoceros.com/GoAghanim/pkg/web"
	"cmonoceros.com/GoAghanim/pkg/web/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"reflect"
)

// UserLogin 用户登录逻辑
func UserLogin(c *gin.Context) {
	var user model.User
	var context = lib.GetContext(c)

	password := context.PostForm("password")
	token := context.Query("api_token")

	db := lib.GetSQLDb(web.ENV)
	defer func() {
		err := db.Close()
		lib.CheckAndThrowError(err, lib.DefaultCode)
	}()
	isNotFound := db.First(&user, "api_token = ?", token).RecordNotFound()
	if isNotFound {
		context.PackError(500, "Can't find user")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		lib.Logger.Error(err)
		context.PackError(500, "Can't find user")
		return
	}

	context.PackMap(lib.GetResource(user, provideUserResource))
}

func provideUserResource(res interface{}) interface{} {
	resType := reflect.TypeOf(res)
	reqType := reflect.TypeOf(model.User{})
	if resType != reqType {
		lib.Logger.Panic("Require " + reqType.String() + ",but get " + resType.String())
		return nil
	}
	user := res.(model.User)
	return map[string]interface{}{
		"api_token":     user.APIToken,
		"user_tel":      user.UserTel,
		"user_tel_zone": user.UserTelZone,
		"user_alias":    user.UserAlias,
	}
}
