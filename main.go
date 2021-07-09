package main

import (
	"context"
	"io"
	"log"
	"net"
	"os"

	graphproto "gitlab.ost.ch/ins/jalapeno-api/request-service/proto/graph-db-feeder"
	rsproto "gitlab.ost.ch/ins/jalapeno-api/request-service/proto/request-service"
	"google.golang.org/grpc"
)

type requestServiceServer struct {
	rsproto.UnimplementedApiGatewayServer
}

func newServer() *requestServiceServer {
	s := &requestServiceServer{}
	return s
}

func main() {
	log.Print("Starting Request Service ...")
	//Start gRPC server for SR-Apps
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
	grpcServer := grpc.NewServer()
	rsproto.RegisterApiGatewayServer(grpcServer, newServer())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}

func (s *requestServiceServer) GetNodes(nodeIds *rsproto.NodeIds, responseStream rsproto.ApiGateway_GetNodesServer) error {
	log.Printf("SR-App requesting Nodes\n")
	log.Printf("Requesting all Nodes from GraphDBFeeder\n")

	//Call GetNodes on GraphDBFeeder
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("GRAPH_DB_FEEDER_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	client := graphproto.NewGraphDbFeederClient(conn)
	message := &graphproto.NodeIds{Ids: nodeIds.Ids}
	stream, err := client.GetNodes(context.Background(), message)

	if err != nil {
		log.Fatalf("Error when calling GetNodes on GraphDBFeeder: %s", err)
	}

	for {
		node, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetNodes(_) = _, %v", client, err)
		}
		responseStream.Send(&rsproto.Node{Id: node.Id, Name: node.Name, Asn: node.Asn})
	}
	return nil

}
