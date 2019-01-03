package lib

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

// Context 默认Context
type Context struct {
	*gin.Context
}

// GetContext 获取默认Context
func GetContext(c *gin.Context) *Context {
	return &Context{
		c,
	}
}

// PackOK 正常返回
func (c *Context) PackOK() {
	c.JSON(http.StatusOK, gin.H{
		"errcode": 0,
		"msg":     "ok",
	})
}

// PackArray 正常返回数组
func (c *Context) PackArray(data []interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"errcode": 0,
		"msg":     "ok",
		"data":    data,
	})
}

// PackMap 正常返回Map
func (c *Context) PackMap(data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"errcode": 0,
		"msg":     "ok",
		"data":    data,
	})
}

// PackError 错误返回
func (c *Context) PackError(code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"errcode": code,
		"msg":     msg,
	})
}

// GetResource 获取返回的资源文件
func GetResource(data interface{}, provideResource func(interface{}) interface{}) (res interface{}) {
	t := reflect.ValueOf(data).Kind()
	if t == reflect.Slice {
		var res = make([]interface{}, 15)
		for _, v := range data.([]interface{}) {
			res = append(res, provideResource(v))
		}
		return res
	} else if t == reflect.Struct {
		return provideResource(data)
	} else {
		CheckAndThrowError(errors.New("Unexpected type:"+t.String()), LibraryCode)
		return nil
	}
}
