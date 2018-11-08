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

func GetDefaultConfig() Config {
	return CONF
}
