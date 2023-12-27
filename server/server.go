package server

import (
	"context"
	"log"
	"math"

	proto "github.com/htetmyomyint-kmp/grpc-helloworld/proto/sqrt"
)

type SqrtServer struct{}

func NewServer() *SqrtServer {
	return &SqrtServer{}
}

func (s *SqrtServer) GetSquareRoot(ctx context.Context, req *proto.MRequest) (*proto.MResponse, error) {
	log.Printf("GetSquareRoot %+v \n", req)

	return &proto.MResponse{Num: math.Sqrt(float64(req.Num))}, nil
}
