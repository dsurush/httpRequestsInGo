package config

type server struct {
	Addr string `json:"addr"`
	Name string `json:"name"`
}

type database struct {
	Addr   string `json:"addr"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	DBName string `json:"dbname"`
}

//Config "Просто структура "
type Config struct {
	Server    server    `json:"server"`
	Database  database  `json:"database"`
	Configure configure `json:"configure"`
}

type configure struct {
	Addr string `json:"addr"`
}

var (
	conf Config
)

// Peek "Для перечади данных с указательями "
func Peek() *Config {
	return &conf
}
