package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/imsilence/smartkms/backend/response"
	"github.com/imsilence/smartkms/backend/services"
)

// Apply 申请KEK密钥
func Apply(ctx *gin.Context) {
	kek, err := services.KekService.Apply()
	response.JSON(ctx, kek, err)
}
