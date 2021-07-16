package main

import (
	"context"
	"io"
	"log"
	"net"
	"os"

	graphproto "gitlab.ost.ch/ins/jalapeno-api/request-service/proto/graph-db-feeder"
	rsproto "gitlab.ost.ch/ins/jalapeno-api/request-service/proto/request-service"
	tsproto "gitlab.ost.ch/ins/jalapeno-api/request-service/proto/tsdb-feeder"
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

func (s *requestServiceServer) GetDataRates(ipv4Addresses *rsproto.IPv4Addresses, responseStream rsproto.ApiGateway_GetDataRatesServer) error {
	log.Printf("SR-App requesting DataRates\n")

	//Call GetDataRate on TSDBFeeder
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("TSDB_FEEDER_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	client := tsproto.NewTsdbFeederClient(conn)
	message := &tsproto.IPv4Addresses{Ipv4Address: ipv4Addresses.Ipv4Address}
	stream, err := client.GetDataRates(context.Background(), message)

	if err != nil {
		log.Fatalf("Error when calling GetDataRate on TSDBFeeder: %s", err)
	}

	for {
		dataRate, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetNodes(_) = _, %v", client, err)
		}
		responseStream.Send(&rsproto.DataRate{DataRate: dataRate.DataRate, Ipv4Address: dataRate.Ipv4Address})
	}
	log.Printf("All DataRates returned to sr-app")
	return nil
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
		responseStream.Send(&rsproto.Node{Key: node.Key, Name: node.Name, Asn: node.Asn})
	}
	log.Printf("Sent back all Nodes to SR-App\n")
	return nil
}
