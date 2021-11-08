package main

import (
	"context"
	"fmt"
	"log"
	"news-api/rpc" //对应的生成文件目录
	"news-api/utils/similarity"

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
	respA, err := t.GetKeywords(context.Background(), &rpc.GetKeywordsReq{Title: "习近平：以共同但有区别的责任原则为基石，全面有效落实《联合国气候变化框架公约》及其《巴黎协定》 | 每经网 - 每日经济新闻"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	respB, err := t.GetKeywords(context.Background(), &rpc.GetKeywordsReq{Title: "中疾控专家：儿童接种新冠疫苗后要留观30分钟，避免剧烈运动_新华报业网 - 新华报业网"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	mpA := make(map[string]float32)
	for _, kw := range respA.Keywords {
		mpA[kw.Word] = kw.Weight
	}

	mpB := make(map[string]float32)
	for _, kw := range respB.Keywords {
		mpB[kw.Word] = kw.Weight
	}

	fmt.Println(similarity.Sim(mpA, mpB))
}
