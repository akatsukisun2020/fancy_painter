package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"fancy_painter/proto/fancy_painter"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func mainforgrpc() {
	// 创建一个监听句柄
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	// 这里是创建的一个服务server，并不是一个service！！
	// 一个server，可以提供多个service。
	// 每一个service，代表对外提供的几种服务！！
	svr := grpc.NewServer()

	// 这个实际上就是通过生成的stub存根，
	// (1) 将对应的存根的service的结构体和名字，先注册到server中
	// 将对应的接口实现，先“赋值”
	fancy_painter.RegisterFancyPainterServer(svr, &FancyPainterService{})
	defer func() {
		svr.Stop()
		listen.Close()
	}()

	fmt.Println("Serving 8001...")
	_ = svr.Serve(listen)
	fmt.Println("Serving Quit...")
}

// 提供http接口服务
func main() {
	s := grpc.NewServer()
	fancy_painter.RegisterFancyPainterServer(s, &FancyPainterService{})

	// 1. 启动grpc服务
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	log.Println("serving grpc on 0.0.0.0")
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve:%v", err)
		}
	}()

	// 2. 启动http服务
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	err = fancy_painter.RegisterFancyPainterHandler(context.Background(), gwmux, conn) // 这里是handler不是service
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}

	log.Println("Serving grpc-gateway on http:0.0.0.0 :8080")
	log.Fatalln(gwServer.ListenAndServe())
}
