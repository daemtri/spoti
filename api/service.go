package api

import (
	"context"
	"embed"
	"errors"
)

//go:embed service.go message.go
var Specific embed.FS

var (
	ErrUnknown          = errors.New("[9999]未知错误")
	ErrInvalidArgument  = errors.New("[1001]请求参数错误")
	ErrPermissionDenied = errors.New("[1002]禁止操作")

	ErrPersonNotFound = errors.New("[4001]用户未找到")
)

// Calc Calc/v1接口服务
// @produce application/json
// @consume application/json
// @error ErrInvalidArgument ErrPermissionDenied
type Calc interface {

	// Add 加法
	Add(ctx context.Context, arg AddArg) (*AddReply, error)

	// Div 除法
	Div(ctx context.Context, arg DivArg) (*DivReply, error)

	// Mul 乘法
	Mul(ctx context.Context, arg MulArg) (*MulReply, error)

	// GetUserInfo 获取用户信息（此接口为了演示自定义Method和路由）
	// @http GET /user/{id}
	// @error ErrPersonNotFound
	GetUserInfo(ctx context.Context, arg GetUserInfoArg) (*GetUserInfoReply, error)
}
