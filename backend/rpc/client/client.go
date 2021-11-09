package client

import (
	"context"
	"log"
	"news-api/rpc" //对应的生成文件目录

	"google.golang.org/grpc"
)

var conn *grpc.ClientConn
var t rpc.GreeterClient

func Dial() {
	// 建立连接到gRPC服务
	var err error
	conn, err = grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 创建Waiter服务的客户端
	t = rpc.NewGreeterClient(conn)
}

func Close() {
	// 函数结束时关闭连接
	conn.Close()
}

func GetKeywords(title string) map[string]float32 {
	resp, err := t.GetKeywords(context.Background(), &rpc.GetKeywordsReq{Title: title})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	mp := make(map[string]float32)
	for _, kw := range resp.Keywords {
		mp[kw.Word] = kw.Weight
	}

	return mp
}
