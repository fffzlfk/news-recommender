package client

import (
	"context"
	"log"
	"news-api/rpc" //对应的生成文件目录

	"google.golang.org/grpc"
)

type Client struct {
	gc rpc.GreeterClient
}

// var conn *grpc.ClientConn
// var gc rpc.GreeterClient

func NewClient() (*Client, func()) {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	closeFunc := func() {
		conn.Close()
	}

	// 创建Waiter服务的客户端
	gc := rpc.NewGreeterClient(conn)

	return &Client{
		gc: gc,
	}, closeFunc
}

func (c *Client) GetKeywords(title string) map[string]float32 {
	resp, err := c.gc.GetKeywords(context.Background(), &rpc.GetKeywordsReq{Title: title})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	mp := make(map[string]float32)
	for _, kw := range resp.Keywords {
		mp[kw.Word] = kw.Weight
	}

	return mp
}
