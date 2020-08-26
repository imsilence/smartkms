package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/imsilence/smartkms/backend/controllers"
)

// Register 注册路由
func Register(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	{
		v1.GET("/kek/apply", controllers.Apply)
		v1.POST("/data/encrypt", controllers.Encrypt)
		v1.POST("/data/decrypt", controllers.Decrypt)
	}
}
