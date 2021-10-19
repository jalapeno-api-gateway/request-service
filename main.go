package main

import (
	"log"
	"net"
	"os"

	"github.com/jalapeno-api-gateway/request-service/helpers"
	"github.com/jalapeno-api-gateway/request-service/influxdb"
	"github.com/jalapeno-api-gateway/request-service/requestservice"
	"github.com/jalapeno-api-gateway/request-service/redis"
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
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

	jagw.RegisterRequestServiceServer(grpcServer, requestservice.NewServer())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
