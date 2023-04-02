package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/aldinofrizal/soompi-colly-scrap/news"
	"github.com/aldinofrizal/soompi-colly-scrap/services"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting app")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on PORT 9000: %v", err)
	}
	fmt.Println("generating listener app")

	s := services.Server{}

	grpcServer := grpc.NewServer()
	pb.RegisterNewsServiceServer(grpcServer, &s)
	fmt.Println("generating grpc server")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve gRPC server on PORT 9000")
	}
	fmt.Println("grpc services running on PORT 9000")

	services.GetNews()
}
