package main

import (
	"fmt"
	"log"
	"net"

	"github.com/aldinofrizal/soompi-colly-scrap/config"
	pb "github.com/aldinofrizal/soompi-colly-scrap/news"
	"github.com/aldinofrizal/soompi-colly-scrap/services"
	"github.com/jasonlvhit/gocron"
	"google.golang.org/grpc"
)

func main() {
	go InitScheduler()
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on PORT 9000: %v", err)
	}

	config.InitDatabase()
	s := services.Server{}

	grpcServer := grpc.NewServer()
	pb.RegisterNewsServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve gRPC server on PORT 9000")
	}
	fmt.Println("grpc services running on PORT 9000")
}

func InitScheduler() {
	gocron.Every(1).Minute().Do(services.Scrap)
	<-gocron.Start()
}
