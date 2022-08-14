package main

import (
	"context"

	"fancy_painter/proto/fancy_painter"
	"fancy_painter/server/fancy_painter/service"
)

type FancyPainterService struct {
	fancy_painter.UnimplementedFancyPainterServer
}

func (s *FancyPainterService) GetUser(ctx context.Context, req *fancy_painter.UserReq) (*fancy_painter.UserRsp, error) {
	return &fancy_painter.UserRsp{}, nil
}

func (s *FancyPainterService) UserLogin(ctx context.Context, req *fancy_painter.UserLoginReq) (*fancy_painter.UserLoginRsp, error) {
	return service.UserLogin(ctx, req)
}

func (s *FancyPainterService) Helloworld(ctx context.Context, req *fancy_painter.HelloworldReq) (*fancy_painter.HelloworldRsp, error) {
	return service.Helloworld(ctx, req)
}
