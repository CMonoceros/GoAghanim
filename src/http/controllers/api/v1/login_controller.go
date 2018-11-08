package v1

import (
	"cmonoceros.com/GoAghanim/src/base"
	"cmonoceros.com/GoAghanim/src/http/response"
	"cmonoceros.com/GoAghanim/src/services/user/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"reflect"
)

func UserLogin(c *gin.Context) {
	var user models.User
	password := c.PostForm("password")
	token := c.Query("api_token")

	db := base.GetSqlDb()
	defer db.Close()
	isNotFound := db.First(&user, "api_token = ?", token).
		RecordNotFound()
	if isNotFound {
		response.GetContext(c).PackError(500, "Can't find user")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		base.Log.Error(err)
		response.GetContext(c).PackError(500, "Can't find user")
		return
	}

	response.GetContext(c).PackMap(response.GetResource(user, provideUserResource))
}

func provideUserResource(res interface{}) interface{} {
	resType := reflect.TypeOf(res)
	reqType := reflect.TypeOf(models.User{})
	if resType != reqType {
		base.Log.Panic("Require " + reqType.String() + ",but get " + resType.String())
		return nil
	}
	user := res.(models.User)
	return map[string]interface{}{
		"api_token":     user.ApiToken,
		"user_tel":      user.UserTel,
		"user_tel_zone": user.UserTelZone,
		"user_alias":    user.UserAlias,
	}
}
