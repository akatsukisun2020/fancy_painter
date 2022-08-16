package service

import (
	"context"
	"fmt"
	"log"

	"github.com/akatsukisun2020/fancy_painter/proto/fancy_painter"
	"github.com/akatsukisun2020/fancy_painter/server/fancy_painter/dao"
)

// SmartDrawImg ..
func SmartDrawImg(ctx context.Context, req *fancy_painter.SmartDrawImgReq) (*fancy_painter.SmartDrawImgRsp, error) {
	rsp := new(fancy_painter.SmartDrawImgRsp)
	log.Printf("in SmartDrawImg, req: %v", req)

	// 参数校验
	if err := CheckSmartDrawImgParam(ctx, req, rsp); err != nil {
		return rsp, err
	}

	// 校验身份
	user, err := dao.GetUserInfoByMobile(ctx, req.GetPhoneNumber())
	if err != nil {
		log.Printf("GetUserInfoByMobile error, req: %v, err:%v", req, err)
		rsp.RetCode, rsp.RetMsg = -3, "GetUserInfoByMobile error"
		return rsp, err
	}
	if user.UserId != req.GetUserId() {
		log.Printf("identity verification error, req: %v, user:%v", req, user)
		rsp.RetCode, rsp.RetMsg = -4, "identity verification failed"
		return rsp, err
	}

	// 录入任务

	return nil, nil
}

// CheckSmartDrawImgParam 校验参数
func CheckSmartDrawImgParam(ctx context.Context, req *fancy_painter.SmartDrawImgReq, rsp *fancy_painter.SmartDrawImgRsp) error {
	if req.GetPhoneNumber() == "" || req.GetUserId() == "" || req.GetSmartDrawImgTask().GetArtist() == "" ||
		req.GetSmartDrawImgTask().GetTargetStyle() == "" || req.GetSmartDrawImgTask().GetImgHeight() == 0 ||
		req.GetSmartDrawImgTask().GetImgWidth() == 0 {
		log.Printf("param error, req: %v", req)
		rsp.RetCode, rsp.RetMsg = -1, "param error"
		return fmt.Errorf("param error")
	}

	return nil
}
