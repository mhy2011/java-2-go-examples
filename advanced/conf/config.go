package conf

// 配置信息
type Config struct {
	Server Server `json:"server"`
	DB     DB     `json:"db"`
}

// 服务配置信息
type Server struct {
	Port        int    `json:"port"`
	Name        string `json:"name"`
	ContextPath string `json:"contextPath"`
}

type DB struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"dbName"`
}
