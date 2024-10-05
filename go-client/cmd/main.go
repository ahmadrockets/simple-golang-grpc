package main

import (
	"context"
	"fmt"
	"go-client/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	grpcServerUrl = "localhost:9001"
)

func main() {

	conn, err := grpc.Dial(grpcServerUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:9001: %v", err)
	}

	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// get list all user
	fmt.Println("=========== CALL GetAllUser ===========")
	userLists, err := c.GetAllUser(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("error calling function GetAllUser: %v", err)
	}
	log.Println("USERLISTS ", userLists)

	// register users
	fmt.Println("=========== CALL Register ===========")
	userReg, err := c.Register(ctx, &pb.User{
		Name:    "Andi Rohmanto",
		Address: "Jl Mangkubumi no 10",
		Phone:   "08129389",
		Email:   "andirohmanto@gmail.com",
	})
	if err != nil {
		log.Fatalf("error calling function Register: %v", err)
	}
	log.Println("USER REGISTER ", userReg)

}
