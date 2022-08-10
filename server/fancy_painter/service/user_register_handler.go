package service

import (
	"context"
	svrHttp "fancy_painter/common/http"
	"fancy_painter/proto/fancy_painter"
	"fancy_painter/server/fancy_painter/aliyun"
	"fancy_painter/server/fancy_painter/config"
	"fmt"
	"log"
)

// UserLogin 用户登录接口，根据登录接口，获取用户资料
func UserLogin(ctx context.Context, req *fancy_painter.UserLoginReq) (*fancy_painter.UserLoginRsp, error) {
	log.Printf("in UserLogin, req: %v", req)
	rsp := &fancy_painter.UserLoginRsp{
		RetCode: 0,
		RetMsg:  "SUCCESS",
	}

	if req.GetToken() == "" {
		log.Printf("param error, req: %v", req)
		rsp.RetCode, rsp.RetMsg = -1, "param error"
		return rsp, fmt.Errorf("param error")
	}

	svrHttp.NewHttpClient("http://baidu.com").Get(ctx)
	dccessKeyID, accessKeySecret, accessToken := config.GetAccessKeyID(), config.GetAccessKeySecret(), req.GetToken()
	mobile, err := aliyun.GetMobile(&dccessKeyID, &accessKeySecret, &accessToken)
	if err != nil {
		log.Printf("GetMobile error, err:%v", err)
		rsp.RetCode, rsp.RetMsg = -2, "GetMobile error"
		return rsp, fmt.Errorf("GetMobile error")
	}
	rsp.UserId = mobile
	return rsp, nil
}
