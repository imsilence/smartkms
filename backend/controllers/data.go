package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/imsilence/smartkms/backend/response"
	"github.com/imsilence/smartkms/backend/services"
)

// Encrypt 数据加密
func Encrypt(ctx *gin.Context) {
	appKey := ctx.DefaultPostForm("app_key", "")
	plaintext := ctx.DefaultPostForm("plaintext", "")
	ciphertext, err := services.DataService.Encrypt(appKey, plaintext)
	response.JSON(ctx, ciphertext, err)
}

// Decrypt 数据解密
func Decrypt(ctx *gin.Context) {
	appKey := ctx.DefaultPostForm("app_key", "")
	ciphertext := ctx.DefaultPostForm("ciphertext", "")
	plaintext, err := services.DataService.Decrypt(appKey, ciphertext)
	response.JSON(ctx, plaintext, err)
}
