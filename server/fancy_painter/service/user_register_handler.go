package service

import (
	"context"
	svrHttp "fancy_painter/common/http"
	"fancy_painter/proto/fancy_painter"
	"log"
)

// UserLogin 用户登录接口，根据登录接口，获取用户资料
func UserLogin(ctx context.Context, req *fancy_painter.UserLoginReq) (*fancy_painter.UserLoginRsp, error) {
	log.Printf("req: %v", req)
	svrHttp.NewHttpClient("http://baidu.com").Get(ctx)
	return &fancy_painter.UserLoginRsp{}, nil
}
