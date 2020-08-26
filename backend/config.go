package backend

// LogConfig Log配置
type LogConfig struct {
	Filename   string
	MaxAge     int
	MaxSize    int
	MaxBackups int
	Compress   bool
	LocalTime  bool
}

// HTTPConfig HTTP服务配置
type HTTPConfig struct {
	Addr string
}

// HTTPSConfig HTTPS服务配置
type HTTPSConfig struct {
	Addr     string
	CertFile string
	KeyFile  string
}

// WebConfig Web服务配置
type WebConfig struct {
	HTTP  *HTTPConfig
	HTTPS *HTTPSConfig
	Log   *LogConfig
}

// DbConfig 数据库配置
type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Db       string
	Charset  string
}

// Config 后端配置
type Config struct {
	Key   string
	Debug bool
	Log   *LogConfig
	Db    *DbConfig
	Web   *WebConfig
}

// ParseConfig 解析后端配置文件
func ParseConfig(path string) (*Config, error) {
	return defaultConfig, nil
}

var defaultConfig = &Config{
	Key:   "8d2582a1ce5637109263b433aa8b0ae7ae2e67ab0aed6f6a3f96a64e8b7c69e6",
	Debug: true,
	Log: &LogConfig{
		Filename:   "logs/backend.log",
		MaxSize:    5,
		MaxBackups: 7,
		Compress:   true,
		LocalTime:  false,
	},
	Web: &WebConfig{
		HTTP: &HTTPConfig{
			Addr: "0.0.0.0:8080",
		},
		Log: &LogConfig{
			Filename:   "logs/web.log",
			MaxSize:    5,
			MaxBackups: 7,
			Compress:   true,
			LocalTime:  false,
		},
	},
	Db: &DbConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "881019",
		Db:       "smartkms",
		Charset:  "utf8mb4",
	},
}
