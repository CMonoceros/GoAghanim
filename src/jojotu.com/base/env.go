package base

type Config struct {
	Port    string
	LogPath string
	AppUrl  string
	MySql   MySql
}

type MySql struct {
	User     string
	Password string
	DBName   string
	Args     map[string]string
}

var conf = Config{}

func GetDefaultConfig() Config {
	return conf
}
