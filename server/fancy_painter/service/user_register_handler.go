package service

import (
	"context"
	"fmt"
	"log"

	"github.com/akatsukisun2020/fancy_painter/mainline/proto/fancy_painter"
	"github.com/akatsukisun2020/fancy_painter/mainline/server/fancy_painter/aliyun"
	"github.com/akatsukisun2020/fancy_painter/mainline/server/fancy_painter/config"
	"github.com/akatsukisun2020/fancy_painter/mainline/server/fancy_painter/dao"

	"github.com/google/uuid"
)

// UserLogin 用户登录接口，根据登录接口，获取用户资料  ==> TODO： 做缓存限制频率，避免击垮后台mysql
func UserLogin(ctx context.Context, req *fancy_painter.UserLoginReq) (*fancy_painter.UserLoginRsp, error) {
	log.Printf("in UserLogin, req: %v", req)
	rsp := &fancy_painter.UserLoginRsp{
		RetCode: 0,
		RetMsg:  "SUCCESS",
	}

	// 参数校验
	if req.GetToken() == "" {
		log.Printf("param error, req: %v", req)
		rsp.RetCode, rsp.RetMsg = -1, "param error"
		return rsp, fmt.Errorf("param error")
	}

	// token校验
	dccessKeyID, accessKeySecret, accessToken := config.GetAccessKeyID(), config.GetAccessKeySecret(), req.GetToken()
	mobile, err := aliyun.GetMobile(&dccessKeyID, &accessKeySecret, &accessToken)
	if err != nil {
		log.Printf("GetMobile error, err:%v", err)
		rsp.RetCode, rsp.RetMsg = -2, "GetMobile error"
		return rsp, fmt.Errorf("GetMobile error")
	}
	rsp.UserId = mobile

	// 创建用户id
	userid := uuid.New().String()
	err = dao.InsertUserIDByMobile(ctx, mobile, userid) // 不存在则插入，存在则忽略
	if err != nil {
		log.Printf("InsertUserIDByMobile error, err:%v", err)
		rsp.RetCode, rsp.RetMsg = -3, "InsertUserIDByMobile error"
		return rsp, fmt.Errorf("InsertUserIDByMobile error")
	}

	return rsp, nil
}
