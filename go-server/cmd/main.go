package main

import (
	"go-server/config"
	"go-server/pb"
	"go-server/repository"
	"go-server/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	log.Println("RUN APP!!")

	userRepo := repository.NewUserRepository()
	UserService := service.NewUserService(userRepo)
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, UserService)

	log.Println("Starting RPC server at", config.ServiceUserPort)

	l, err := net.Listen("tcp", config.ServiceUserPort)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.ServiceUserPort, err)
	}

	log.Fatal(grpcServer.Serve(l))

}
