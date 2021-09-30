package requestservice

import (
	context "context"
	"log"

	"github.com/jalapeno-api-gateway/jagw-core/model/class"
)

type requestServiceServer struct {
	UnimplementedRequestServiceServer
}

func NewServer() *requestServiceServer {
	s := &requestServiceServer{}
	return s
}

func (s *requestServiceServer) GetLSNodes(ctx context.Context, request *TopologyRequest) (*LSNodeResponse, error) {
	log.Printf("SR-App requesting Nodes\n")

	documents := fetchDocuments(ctx, request.Keys, class.LSNode)

	response := &LSNodeResponse{}

	for _, document := range documents {
		lsNode := convertLSNode(document, request.PropertyNames)
		response.LsNodes = append(response.LsNodes, lsNode)
	}
	
	return response, nil
}

func (s *requestServiceServer) GetLSLinks(ctx context.Context, request *TopologyRequest) (*LSLinkResponse, error) {
	log.Printf("SR-App requesting Links\n")

	documents := fetchDocuments(ctx, request.Keys, class.LSLink)

	response := &LSLinkResponse{}

	for _, document := range documents {
		lsLink := convertLSLink(document, request.PropertyNames)
		response.LsLinks = append(response.LsLinks, lsLink)
	}
	
	return response, nil
}

func (s *requestServiceServer) GetLSPrefixes(ctx context.Context, request *TopologyRequest) (*LSPrefixResponse, error) {
	log.Printf("SR-App requesting LSPrefix\n")

	documents := fetchDocuments(ctx, request.Keys, class.LSPrefix)

	response := &LSPrefixResponse{}

	for _, document := range documents {
		lsPrefix := convertLSPrefix(document, request.PropertyNames)
		response.LsPrefixes = append(response.LsPrefixes, lsPrefix)
	}
	
	return response, nil
}

func (s *requestServiceServer) GetLSSRv6SID(ctx context.Context, request *TopologyRequest) (*LSSRv6SIDResponse, error) {
	log.Printf("SR-App requesting LSSRv6SID\n")

	documents := fetchDocuments(ctx, request.Keys, class.LSSRv6SID)

	response := &LSSRv6SIDResponse{}

	for _, document := range documents {
		lsSRv6SID := convertLSSRv6SID(document, request.PropertyNames)
		response.LsSRv6SIDs = append(response.LsSRv6SIDs, lsSRv6SID)
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