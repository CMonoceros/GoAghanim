package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"jojotu.com/base"
	"jojotu.com/http/response"
	"jojotu.com/services/user/models"
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
	var a [4]models.User
	a[0] = user
	a[1] = user
	base.Log.Info(response.GetResource(a, provideUserResource))
	base.Log.Info(response.GetResource(user, provideUserResource))
	response.GetContext(c).PackMap(user)
}

func provideUserResource(res map[string]interface{}) interface{} {
	//resType := reflect.TypeOf(res)
	//reqType := reflect.TypeOf(models.User{})
	//if resType != reqType {
	//	base.Log.Panic("Require " + reqType.String() + ",but get " + resType.String())
	//	return nil
	//}
	return struct {
		api_token     string
		user_tel      string
		user_tel_zone int
		user_alias    string
	}{
		api_token:     res["ApiToken"].(string),
		user_tel:      res["UserTel"].(string),
		user_tel_zone: res["UserTelZone"].(int),
		user_alias:    res["UserAlias"].(string),
	}
}
