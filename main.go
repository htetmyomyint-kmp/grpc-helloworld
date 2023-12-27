package main

import (
	"log"
	"net"

	"github.com/htetmyomyint-kmp/grpc-helloworld/client"
	proto "github.com/htetmyomyint-kmp/grpc-helloworld/proto/sqrt"
	"github.com/htetmyomyint-kmp/grpc-helloworld/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	startServer()

}

func startServer() {
	gs := grpc.NewServer()

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Panicln("net listen", err)
	}

	reflection.Register(gs)

	gc := server.NewServer()
	proto.RegisterSqrtServer(gs, gc)

	go client.MakeSqrtCall()

	if err := gs.Serve(l); err != nil {
		log.Panicln("grpc serve err", err)
	}

}
