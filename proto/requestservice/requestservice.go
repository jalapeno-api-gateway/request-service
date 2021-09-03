package requestservice

import (
	context "context"
	"log"

	"gitlab.ost.ch/ins/jalapeno-api/request-service/influxdb"
	"gitlab.ost.ch/ins/jalapeno-api/request-service/redis"
)

type requestServiceServer struct {
	UnimplementedApiGatewayServer
}

func NewServer() *requestServiceServer {
	s := &requestServiceServer{}
	return s
}

func (s *requestServiceServer) GetLsNodes(request *LsNodeRequest, responseStream ApiGateway_GetLsNodesServer) error {
	log.Printf("SR-App requesting Nodes\n")
	ctx := context.Background()

	documents := []redis.LsNodeDocument{}
	if len(request.Keys) == 0 {
		documents = redis.FetchAllLsNodes(ctx)
	} else {
		documents = redis.FetchLsNodes(ctx, request.Keys)
	}

	for _, document := range documents {
		lsNode := convertToGrpcLsNode(document)
		err := responseStream.Send(&lsNode)
		if err != nil {
			log.Fatal("Unable to send LsNode: ", err)
		}
	}

	return nil
}

func (s *requestServiceServer) GetLsLinks(request *LsLinkRequest, responseStream ApiGateway_GetLsLinksServer) error {
	log.Printf("SR-App requesting Links\n")
	ctx := context.Background()

	documents := []redis.LsLinkDocument{}
	if len(request.Keys) == 0 {
		documents = redis.FetchAllLsLinks(ctx)
	} else {
		documents = redis.FetchLsLinks(ctx, request.Keys)
	}

	for _, document := range documents {
		lsNode := convertToGrpcLsLink(document)
		err := responseStream.Send(&lsNode)
		if err != nil {
			log.Fatal("Unable to send LsLink: ", err)
		}
	}

	return nil
}

func (s *requestServiceServer) GetDataRates(request *DataRateRequest, responseStream ApiGateway_GetDataRatesServer) error {
	log.Printf("SR-App requesting DataRates\n")

	dataRates := influxdb.FetchDataRates(influxdb.InfluxClient, request.Ipv4Addresses)

	for _, dataRate := range dataRates {
		response := convertToGrpcDataRate(dataRate)
		if err := responseStream.Send(&response); err != nil {
			log.Fatalf("Could not return dataRate to request-service, %v", err)
		}
	}
	return nil
}
