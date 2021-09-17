package requestservice

import (
	context "context"
	"log"

	"gitlab.ost.ch/ins/jalapeno-api/request-service/redis"
)

type requestServiceServer struct {
	UnimplementedRequestServiceServer
}

func NewServer() *requestServiceServer {
	s := &requestServiceServer{}
	return s
}

func (s *requestServiceServer) GetLsNodes(request *TopologyRequest, responseStream RequestService_GetLsNodesServer) error {
	log.Printf("SR-App requesting Nodes\n")
	ctx := context.Background()

	documents := []redis.LsNodeDocument{}
	if len(request.Keys) == 0 {
		documents = redis.FetchAllLsNodes(ctx)
	} else {
		documents = redis.FetchLsNodes(ctx, request.Keys)
	}

	for _, document := range documents {
		lsNode := convertLsNode(document, request.PropertyNames)
		err := responseStream.Send(lsNode)
		if err != nil {
			log.Fatal("Unable to send LsNode: ", err)
		}
	}

	return nil
}

func (s *requestServiceServer) GetLsLinks(request *TopologyRequest, responseStream RequestService_GetLsLinksServer) error {
	log.Printf("SR-App requesting Links\n")
	ctx := context.Background()

	documents := []redis.LsLinkDocument{}
	if len(request.Keys) == 0 {
		documents = redis.FetchAllLsLinks(ctx)
	} else {
		documents = redis.FetchLsLinks(ctx, request.Keys)
	}

	for _, document := range documents {
		lsNode := convertLsLink(document, request.PropertyNames)
		err := responseStream.Send(lsNode)
		if err != nil {
			log.Fatal("Unable to send LsLink: ", err)
		}
	}

	return nil
}

func (s *requestServiceServer) GetTelemetryData(request *TelemetryRequest, responseStream RequestService_GetTelemetryDataServer) error {
	log.Printf("SR-App requesting DataRates\n")

	for _, ipv4address := range request.Ipv4Addresses {
		response := fetchTelemetryResponse(ipv4address, request.PropertyNames)
		if err := responseStream.Send(response); err != nil {
			log.Fatalf("Could not return TelemetryData to SR-App, %v", err)
		}
	}
	return nil
}