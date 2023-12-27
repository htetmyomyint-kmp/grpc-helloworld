package client

import (
	"context"
	"log"
	"time"

	proto "github.com/htetmyomyint-kmp/grpc-helloworld/proto/sqrt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SqrtClient struct{}

func NewSqrtClient() *SqrtClient {
	return &SqrtClient{}
}

func MakeSqrtCall() {
	time.Sleep(time.Second * 2)
	conn, err := grpc.Dial("localhost:9092", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicln("grpc server dial err", err)
	}
	defer conn.Close()

	client := proto.NewSqrtClient(conn)

	resp, err := client.GetSquareRoot(context.TODO(), &proto.MRequest{Num: 49})

	if err != nil {
		log.Panicln("grpc call err", err)
	}

	log.Println("response", resp.Num)

}
