package backend

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// Backend 后端服务
type Backend struct {
	Engine *gin.Engine
	Config *Config
	Db     *gorm.DB
}

// Run 服务启动
func (b *Backend) Run() error {
	errChan := make(chan error, 1)
	stopChan := make(chan os.Signal, 1)

	/*
		SIGINT: ctrl+c
		SIGTERM: kill pid
		SIGKILL: kill -9 pid
	*/
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	if b.Config.Web.HTTP != nil && b.Config.Web.HTTP.Addr != "" {
		go func() {
			// 启动HTTP服务
			logrus.WithFields(logrus.Fields{"addr": b.Config.Web.HTTP.Addr}).Info("smartkms listen http")
			errChan <- b.Engine.Run(b.Config.Web.HTTP.Addr)
		}()
	}

	if b.Config.Web.HTTPS != nil && b.Config.Web.HTTPS.Addr != "" {
		go func() {
			//  启动HTTPS服务
			logrus.WithFields(logrus.Fields{"addr": b.Config.Web.HTTPS.Addr}).Info("smartkms listen https")
			errChan <- b.Engine.RunTLS(
				b.Config.Web.HTTPS.Addr,
				b.Config.Web.HTTPS.CertFile,
				b.Config.Web.HTTPS.KeyFile,
			)
		}()
	}

	logrus.WithFields(logrus.Fields{"pid": os.Getpid()}).Info("smartkms running...")

	select {
	case err := <-errChan:
		return err
	case <-stopChan:
		return nil
	}
}

// NewBackend 创建后端服务
func NewBackend() (*Backend, error) {
	return NewBackendWithConfig(defaultConfig)
}

// NewBackendWithConfigFile 根据配置文件创建后端服务
func NewBackendWithConfigFile(path string) (*Backend, error) {
	config, err := ParseConfig(path)
	if err != nil {
		return nil, err
	}
	return NewBackendWithConfig(config)
}

// NewBackendWithConfig 根据配置创建后端服务
func NewBackendWithConfig(config *Config) (*Backend, error) {
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DisableConsoleColor()

	engine := gin.New()
	engine.Use(gin.Recovery())

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=PRC",
		config.Db.User,
		config.Db.Password,
		config.Db.Host,
		config.Db.Port,
		config.Db.Db,
		config.Db.Charset,
	)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.LogMode(config.Debug)
	db.SingularTable(true)

	if config.Debug {
		engine.Use(gin.Logger())
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.TextFormatter{})
	} else {
		engine.Use(gin.LoggerWithWriter(NewLogger(config.Web.Log)))

		logrus.SetOutput(NewLogger(config.Log))
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetReportCaller(true)
	}

	return &Backend{Engine: engine, Config: config, Db: db}, nil
}

// App 后端应用对象
var App *Backend

// InitApp 初始化后端应用对象
func InitApp() (*Backend, error) {
	var err error
	App, err = NewBackend()
	if err != nil {
		return nil, err
	}
	return App, nil
}

// InitAppWithConfigFile 根据配置文件初始化后端应用对象
func InitAppWithConfigFile(path string) (*Backend, error) {
	var err error
	App, err = NewBackendWithConfigFile(path)
	if err != nil {
		return nil, err
	}
	return App, nil
}
