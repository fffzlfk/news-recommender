package client

import (
	"context"
	"fmt"
	"log"
	"news-api/rpc" //对应的生成文件目录
	"os"

	"google.golang.org/grpc"
)

type Client struct {
	gc rpc.GreeterClient
}

// var conn *grpc.ClientConn
// var gc rpc.GreeterClient

func NewClient() (*Client, func()) {
	// 建立连接到gRPC服务
	dsn := fmt.Sprintf("%s:50052", os.Getenv("KEY_WORDS_HOST"))
	conn, err := grpc.Dial(dsn, grpc.WithInsecure())
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
		log.Printf("could not greet: %v\n", err)
	}

	mp := make(map[string]float32)
	for _, kw := range resp.Keywords {
		mp[kw.Word] = kw.Weight
	}

	return mp
}
