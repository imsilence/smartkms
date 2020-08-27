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

// JSONBadRequest 响应
func JSONBadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, BadRequest(err))
}

// JSONUnauthorized 响应
func JSONUnauthorized(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Unauthorized())
}

// Ok 构建Response对象
func Ok(result interface{}) *Response {
	return &Response{
		Code:   200,
		Text:   "OK",
		Result: result,
	}
}

// InternalServerError 构建Response对象
func InternalServerError(err error) *Response {
	return &Response{
		Code:   http.StatusBadRequest,
		Text:   err.Error(),
		Result: nil,
	}
}

// BadRequest 构建Response对象
func BadRequest(err error) *Response {
	return &Response{
		Code:   http.StatusBadRequest,
		Text:   err.Error(),
		Result: nil,
	}
}

// Unauthorized 构建Response对象
func Unauthorized() *Response {
	return &Response{
		Code:   http.StatusUnauthorized,
		Text:   http.StatusText(http.StatusUnauthorized),
		Result: nil,
	}
}
