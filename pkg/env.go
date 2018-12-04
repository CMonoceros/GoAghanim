package pkg

var CONF = Config{
	Port:    "8000",
	LogPath: "logs",
	AppUrl:  "jojotoo.io",
	MySql: MySql{
		User:     "root",
		Password: "zhangjunming1921",
		DBName:   "jojotoo_test",
		Args: map[string]string{
			"charset":   "utf8mb4",
			"parseTime": "True",
		},
	},
}
