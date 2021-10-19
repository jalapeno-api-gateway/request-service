package requestservice

import (
	context "context"
	"log"

	"github.com/jalapeno-api-gateway/jagw-core/model/class"
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
)

type requestServiceServer struct {
	jagw.UnimplementedRequestServiceServer
}

func NewServer() *requestServiceServer {
	s := &requestServiceServer{}
	return s
}

func (s *requestServiceServer) GetLsNodes(ctx context.Context, request *jagw.TopologyRequest) (*jagw.LsNodeResponse, error) {
	log.Printf("SR-App requesting Nodes\n")

	documents := fetchDocuments(ctx, request.Keys, class.LsNode)

	response := &jagw.LsNodeResponse{}

	for _, document := range documents {
		lsNode := convertLsNode(document, request.PropertyNames)
		response.LsNodes = append(response.LsNodes, lsNode)
	}
	
	return response, nil
}

func (s *requestServiceServer) GetLsLinks(ctx context.Context, request *jagw.TopologyRequest) (*jagw.LsLinkResponse, error) {
	log.Printf("SR-App requesting Links\n")

	documents := fetchDocuments(ctx, request.Keys, class.LsLink)

	response := &jagw.LsLinkResponse{}

	for _, document := range documents {
		lsLink := convertLsLink(document, request.PropertyNames)
		response.LsLinks = append(response.LsLinks, lsLink)
	}
	
	return response, nil
}

func (s *requestServiceServer) GetLsPrefixes(ctx context.Context, request *jagw.TopologyRequest) (*jagw.LsPrefixResponse, error) {
	log.Printf("SR-App requesting LSPrefix\n")

	documents := fetchDocuments(ctx, request.Keys, class.LsPrefix)

	response := &jagw.LsPrefixResponse{}

	for _, document := range documents {
		lsPrefix := convertLsPrefix(document, request.PropertyNames)
		response.LsPrefixes = append(response.LsPrefixes, lsPrefix)
	}
	
	return response, nil
}

func (s *requestServiceServer) GetLsSrv6Sids(ctx context.Context, request *jagw.TopologyRequest) (*jagw.LsSrv6SidResponse, error) {
	log.Printf("SR-App requesting LSSRv6SID\n")

	documents := fetchDocuments(ctx, request.Keys, class.LsSrv6Sid)

	response := &jagw.LsSrv6SidResponse{}

	for _, document := range documents {
		lsSRv6SID := convertLsSrv6Sid(document, request.PropertyNames)
		response.LsSrv6Sids = append(response.LsSrv6Sids, lsSRv6SID)
	}
	
	return response, nil
}

func (s *requestServiceServer) GetLsNodeEdges(ctx context.Context, request *jagw.TopologyRequest) (*jagw.LsNodeEdgeResponse, error) {
	log.Printf("SR-App requesting LSNodeEdges\n")

	documents := fetchDocuments(ctx, request.Keys, class.LsNodeEdge)

	response := &jagw.LsNodeEdgeResponse{}

	for _, document := range documents {
		lsNodeEdge := convertLsNodeEdge(document, request.PropertyNames)
		response.LsNodeEdges = append(response.LsNodeEdges, lsNodeEdge)
	}
	
	return response, nil
}

func (s *requestServiceServer) GetTelemetryData(ctx context.Context, request *jagw.TelemetryRequest) (*jagw.TelemetryResponse, error) {
	log.Printf("SR-App requesting DataRates\n")

	response := &jagw.TelemetryResponse{}

	for _, interfaceId := range request.InterfaceIds {
		telemetryData := fetchTelemetryData(interfaceId, request.PropertyNames)
		response.TelemetryData = append(response.TelemetryData, telemetryData)
	}

	return response, nil
}