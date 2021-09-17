package main

import (
	"log"
	"net"
	"os"

	"gitlab.ost.ch/ins/jalapeno-api/request-service/helpers"
	"gitlab.ost.ch/ins/jalapeno-api/request-service/influxdb"
	"gitlab.ost.ch/ins/jalapeno-api/request-service/proto/requestservice"
	"gitlab.ost.ch/ins/jalapeno-api/request-service/redis"
	"google.golang.org/grpc"
)

func main() {
	log.Print("Starting Request Service ...")
	redis.InitializeRedisClient()
	influxdb.InitializeInfluxClient()

	serverAddress := os.Getenv("APP_SERVER_ADDRESS")
	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", serverAddress, err)
	}

	grpcServer := grpc.NewServer()

	signals := helpers.WatchInterruptSignals()
	go func() {
		<-signals
		grpcServer.Stop()
	}()

	requestservice.RegisterRequestServiceServer(grpcServer, requestservice.NewServer())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
