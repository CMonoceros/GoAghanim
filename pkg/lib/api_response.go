package lib

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Context struct {
	*gin.Context
}

func GetContext(c *gin.Context) *Context {
	return &Context{
		c,
	}
}

func (c *Context) PackOK() {
	c.JSON(http.StatusOK, gin.H{
		"errcode": 0,
		"msg":     "ok",
	})
}

func (c *Context) PackArray(data []interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"errcode": 0,
		"msg":     "ok",
		"data":    data,
	})
}

func (c *Context) PackMap(data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"errcode": 0,
		"msg":     "ok",
		"data":    data,
	})
}

func (c *Context) PackError(code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"errcode": code,
		"msg":     msg,
	})
}

func GetResource(data interface{}, provideResource func(interface{}) (res interface{})) (res interface{}) {
	t := reflect.ValueOf(data).Kind()
	if t == reflect.Slice {
		var res []interface{}
		for _, v := range data.([]interface{}) {
			res = append(res, provideResource(v))
		}
		return res
	} else if t == reflect.Struct {
		return provideResource(data)
	} else {
		Log.Panic("Unexpected type:" + t.String())
		return nil
	}
}
