package web

import "cmonoceros.com/GoAghanim/pkg/lib"

// ENV 配置文件
var ENV = lib.Env{FileName: "env/gin.env"}

// LogAPIConfig 接口日志配置文件
var LogAPIConfig = lib.LogFileConfig{
	BaseLevel:  "gin-api-info-data-",
	ErrorLevel: "gin-api-error-data-",
}

// LogAdminConfig 后台日志配置文件
var LogAdminConfig = lib.LogFileConfig{
	BaseLevel:  "gin-admin-info-data-",
	ErrorLevel: "gin-admin-error-data-",
}

// LogRequestConfig 请求日志配置文件
var LogRequestConfig = lib.LogFileConfig{
	BaseLevel:  "gin-request-info-data-",
	ErrorLevel: "gin-request-error-data-",
}
