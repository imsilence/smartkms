package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 响应对象
type Response struct {
	Code   int         `json:"code"`
	Text   string      `json:"text"`
	Result interface{} `json:"result"`
}

// JSON 响应
func JSON(ctx *gin.Context, result interface{}, err error) {
	if err == nil {
		ctx.JSON(http.StatusOK, Ok(result))
	} else {
		ctx.JSON(http.StatusOK, InternalServerError(err))
	}
}

// Ok 构建Response对象
func Ok(result interface{}) *Response {
	return &Response{
		Code:   200,
		Text:   "",
		Result: result,
	}
}

// InternalServerError 构建Response对象
func InternalServerError(err error) *Response {
	return &Response{
		Code:   500,
		Text:   err.Error(),
		Result: nil,
	}
}
