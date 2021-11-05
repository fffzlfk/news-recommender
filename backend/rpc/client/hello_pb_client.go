package main

import (
	"context"
	"fmt"
	"log"
	"news-api/rpc" //对应的生成文件目录

	"google.golang.org/grpc"
)

func main() {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 函数结束时关闭连接
	defer conn.Close()
	// 创建Waiter服务的客户端
	t := rpc.NewGreeterClient(conn)
	resp, err := t.GetKeywords(context.Background(), &rpc.GetKeywordsReq{Title: "全球征集超2700件作品近150个网络平台播出成都大运会发出“青春的邀约” | 每经网 - 每日经济新闻"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	kws := resp.Keywords
	for _, kw := range kws {
		fmt.Printf("(%v, %v)\n", kw.Word, kw.Weight)
	}
}
