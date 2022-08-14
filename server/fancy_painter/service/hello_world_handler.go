package service

import (
	"context"
	"fancy_painter/proto/fancy_painter"
	"fmt"
	"log"
)

// Helloworld ..
func Helloworld(ctx context.Context, req *fancy_painter.HelloworldReq) (*fancy_painter.HelloworldRsp, error) {
	log.Printf("in Helloworld, req: %v", req)

	rsp := &fancy_painter.HelloworldRsp{
		Msg: fmt.Sprintf("你好,%s", req.GetName()),
	}

	return rsp, nil
}
