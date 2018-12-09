package main

import (
	"context"
	"flag"
	"log"

	grpc2 "github.com/emurmotol/pbgrpc/ethereum/client/grpc"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

var serverAddr = flag.String("addr", ":8082", "gRPC server address")

func main() {
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("grpc.Dial", err)
	}
	defer conn.Close()

	var options map[string][]grpc1.ClientOption
	ethereumService, err := grpc2.New(conn, options)
	if err != nil {
		log.Fatalln("grpc2.New", err)
	}

	ethereumService.CreateAccount(context.Background(), "123123123")
}
