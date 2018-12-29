package lib

import "github.com/joho/godotenv"

// Env 配置文件结构
type Env struct {
	FileName string
}

// Get 返回指定配置
func (env *Env) Get(key string) (value string) {
	envMap, err := godotenv.Read(env.FileName)
	CheckAndThrowError(err, LibraryCode)
	return envMap[key]
}
