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

func (s *requestServiceServer) GetLsNodes(ctx context.Context, request *TopologyRequest) (*LsNodeResponse, error) {
	log.Printf("SR-App requesting Nodes\n")

	var documents []redis.LsNodeDocument
	if len(request.Keys) == 0 {
		documents = redis.FetchAllLsNodes(ctx)
	} else {
		documents = redis.FetchLsNodes(ctx, request.Keys)
	}

	response := &LsNodeResponse{}

	for _, document := range documents {
		lsNode := convertLsNode(document, request.PropertyNames)
		response.LsNodes = append(response.LsNodes, lsNode)
	}
	
	return response, nil
}

func (s *requestServiceServer) GetLsLinks(ctx context.Context, request *TopologyRequest) (*LsLinkResponse, error) {
	log.Printf("SR-App requesting Links\n")

	var documents []redis.LsLinkDocument
	if len(request.Keys) == 0 {
		documents = redis.FetchAllLsLinks(ctx)
	} else {
		documents = redis.FetchLsLinks(ctx, request.Keys)
	}

	response := &LsLinkResponse{}

	for _, document := range documents {
		lsLink := convertLsLink(document, request.PropertyNames)
		response.LsLinks = append(response.LsLinks, lsLink)
	}
	
	return response, nil
}

func (s *requestServiceServer) GetTelemetryData(ctx context.Context, request *TelemetryRequest) (*TelemetryResponse, error) {
	log.Printf("SR-App requesting DataRates\n")

	response := &TelemetryResponse{}

	for _, ipv4address := range request.Ipv4Addresses {
		telemetryData := fetchTelemetryData(ipv4address, request.PropertyNames)
		response.TelemetryData = append(response.TelemetryData, telemetryData)
	}

	return response, nil
}