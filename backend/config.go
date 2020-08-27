package backend

import (
	"github.com/spf13/viper"
)

// LogConfig Log配置
type LogConfig struct {
	Filename   string `mapstructure:"filename"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	Compress   bool   `mapstructure:"compress"`
	LocalTime  bool   `mapstructure:"local_time"`
}

// SetDefaults 设置默认值
func (c *LogConfig) SetDefaults() {
	if c.Filename == "" {
		c.Filename = "logs/smartkms.log"
	}
}

// HTTPConfig HTTP服务配置
type HTTPConfig struct {
	Addr string `mapstructure:"addr"`
}

// SetDefaults 设置默认值
func (c *HTTPConfig) SetDefaults() {
	if c.Addr == "" {
		c.Addr = defaultConfig.Web.HTTP.Addr
	}
}

// HTTPSConfig HTTPS服务配置
type HTTPSConfig struct {
	Addr     string `mapstructure:"addr"`
	CertFile string `mapstructure:"cert_file"`
	KeyFile  string `mapstructure:"key_file"`
}

// SetDefaults 设置默认值
func (c *HTTPSConfig) SetDefaults() {
	if c.Addr == "" {
		c.Addr = defaultConfig.Web.HTTPS.Addr
	}

	if c.CertFile == "" {
		c.CertFile = defaultConfig.Web.HTTPS.CertFile
	}

	if c.KeyFile == "" {
		c.KeyFile = defaultConfig.Web.HTTPS.KeyFile
	}
}

// WebConfig Web服务配置
type WebConfig struct {
	HTTP  *HTTPConfig  `mapstructure:"http"`
	HTTPS *HTTPSConfig `mapstructure:"https"`
	Log   *LogConfig   `mapstructure:"log"`
}

// SetDefaults 设置默认值
func (c *WebConfig) SetDefaults() {
	if c.HTTP == nil && c.HTTPS == nil {
		c.HTTP = defaultConfig.Web.HTTP
	}

	if c.HTTP != nil {
		c.HTTP.SetDefaults()
	}

	if c.HTTPS != nil {
		c.HTTPS.SetDefaults()
	}

	if c.Log == nil {
		c.Log = defaultConfig.Web.Log
	} else {
		c.Log.SetDefaults()
	}
}

// DbConfig 数据库配置
type DbConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Db       string `mapstructure:"db"`
	Charset  string `mapstructure:"charset"`
}

// SetDefaults 设置默认值
func (c *DbConfig) SetDefaults() {
	if c.Host == "" {
		c.Host = defaultConfig.Db.Host
	}

	if c.Port <= 0 {
		c.Port = defaultConfig.Db.Port
	}

	if c.User == "" {
		c.User = defaultConfig.Db.User
	}

	if c.Password == "" {
		c.Password = defaultConfig.Db.Password
	}

	if c.Db == "" {
		c.Db = defaultConfig.Db.Db
	}

	if c.Charset == "" {
		c.Charset = defaultConfig.Db.Charset
	}
}

// Config 后端配置
type Config struct {
	Key   string     `mapstructure:"key"`
	Debug bool       `mapstructure:"debug"`
	Log   *LogConfig `mapstructure:"log"`
	Db    *DbConfig  `mapstructure:"db"`
	Web   *WebConfig `mapstructure:"web"`
}

// SetDefaults 设置默认值
func (c *Config) SetDefaults() {
	if c.Key == "" {
		c.Key = defaultConfig.Key
	}

	if c.Log == nil {
		c.Log = defaultConfig.Log
	} else {
		c.Log.SetDefaults()
	}

	if c.Db == nil {
		c.Db = defaultConfig.Db
	} else {
		c.Db.SetDefaults()
	}

	if c.Web == nil {
		c.Web = defaultConfig.Web
		c.Web.HTTPS = nil
	} else {
		c.Web.SetDefaults()
	}
}

// ParseConfig 解析后端配置文件
func ParseConfig(path string) (*Config, error) {
	var config Config

	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	config.SetDefaults()

	return &config, nil
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
			Addr: "127.0.0.1:8080",
		},
		HTTPS: &HTTPSConfig{
			Addr:     "127.0.0.1:8443",
			CertFile: "etc/ssl/server.pem",
			KeyFile:  "etc/ssl/key.pem",
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
		Password: "",
		Db:       "smartkms",
		Charset:  "utf8mb4",
	},
}
