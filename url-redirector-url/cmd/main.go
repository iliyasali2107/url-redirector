package main

import (
	"fmt"
	"log"
	"net"
	"url-redirecter-url/pkg/config"
	"url-redirecter-url/pkg/db"
	"url-redirecter-url/pkg/pb"
	"url-redirecter-url/pkg/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed to config", err)
	}

	storage := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listening")
	}

	fmt.Println("Url service is on: ", c.Port)

	srv := service.NewService(storage)

	grpcServer := grpc.NewServer()

	pb.RegisterURLServiceServer(grpcServer, srv)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve: ", err)
	}
}
