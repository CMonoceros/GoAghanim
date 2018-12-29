package lib

import (
	"fmt"
	"reflect"
)

// Error 自定义错误
type Error struct {
	Code  int
	Error error
}

func (e Error) String() string {
	return fmt.Sprintf("Code: %v  =>  Error: %v", e.Code, e.Error)
}

// DefaultCode 默认错误码
const DefaultCode = 45000

// LibraryCode 依赖库错误码
const LibraryCode = 55000

// CheckAndThrowError 检查并扔出错误
func CheckAndThrowError(sysError error, code int) {
	if sysError != nil {
		ThrowError(sysError, code)
	}
}

// ThrowError 扔出错误
func ThrowError(sysError error, code int) {
	err := Error{
		Code:  code,
		Error: sysError,
	}
	Logger.Error(err)
	panic(err)
}

// CheckAndReportError 检查并打印错误，然后进行回调
func CheckAndReportError(err error, callback func()) {
	if err != nil {
		Logger.Error(err)
		callback()
	}
}

// CatchErrorFunc 捕获错误，如为自定义错误类型，则恢复
func CatchErrorFunc() func() {
	return func() {
		err := recover()
		if err != nil && reflect.TypeOf(err) == reflect.TypeOf(Error{}) {
			fmt.Println(err)
		} else {
			panic(err)
		}
	}
}
