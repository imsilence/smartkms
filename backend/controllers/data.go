package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/imsilence/smartkms/backend/common/response"
	"github.com/imsilence/smartkms/backend/request"
	"github.com/imsilence/smartkms/backend/services"
)

// Encrypt 数据加密
func Encrypt(ctx *gin.Context) {
	req := request.EncryptRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.JSONBadRequest(ctx, err)
	} else {
		ciphertext, err := services.DataService.Encrypt(&req)
		response.JSON(ctx, ciphertext, err)
	}
}

// Decrypt 数据解密
func Decrypt(ctx *gin.Context) {
	req := &request.DecryptRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		response.JSONBadRequest(ctx, err)
	} else {
		plaintext, err := services.DataService.Decrypt(req)
		response.JSON(ctx, plaintext, err)
	}
}
