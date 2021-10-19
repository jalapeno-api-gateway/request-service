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

	documents := fetchDocuments(ctx, request.Keys, class.LSNode)

	response := &jagw.LsNodeResponse{}

	for _, document := range documents {
		lsNode := convertLSNode(document, request.PropertyNames)
		response.LsNodes = append(response.LsNodes, lsNode)
	}
	
	return response, nil
}

func (s *requestServiceServer) GetLsLinks(ctx context.Context, request *jagw.TopologyRequest) (*jagw.LsLinkResponse, error) {
	log.Printf("SR-App requesting Links\n")

	documents := fetchDocuments(ctx, request.Keys, class.LSLink)

	response := &jagw.LsLinkResponse{}

	for _, document := range documents {
		lsLink := convertLSLink(document, request.PropertyNames)
		response.LsLinks = append(response.LsLinks, lsLink)
	}
	
	return response, nil
}

func (s *requestServiceServer) GetLsPrefixes(ctx context.Context, request *jagw.TopologyRequest) (*jagw.LsPrefixResponse, error) {
	log.Printf("SR-App requesting LSPrefix\n")

	documents := fetchDocuments(ctx, request.Keys, class.LSPrefix)

	response := &jagw.LsPrefixResponse{}

	for _, document := range documents {
		lsPrefix := convertLSPrefix(document, request.PropertyNames)
		response.LsPrefixes = append(response.LsPrefixes, lsPrefix)
	}
	
	return response, nil
}

func (s *requestServiceServer) GetLsSrv6Sids(ctx context.Context, request *jagw.TopologyRequest) (*jagw.LsSrv6SidResponse, error) {
	log.Printf("SR-App requesting LSSRv6SID\n")

	documents := fetchDocuments(ctx, request.Keys, class.LSSRv6SID)

	response := &jagw.LsSrv6SidResponse{}

	for _, document := range documents {
		lsSRv6SID := convertLSSRv6SID(document, request.PropertyNames)
		response.LsSrv6Sids = append(response.LsSrv6Sids, lsSRv6SID)
	}
	
	return response, nil
}

func (s *requestServiceServer) GetLsNodeEdge(ctx context.Context, request *jagw.TopologyRequest) (*jagw.LsNodeEdgeResponse, error) {
	log.Printf("SR-App requesting LSNodeEdges\n")

	documents := fetchDocuments(ctx, request.Keys, class.LSSRv6SID)

	response := &jagw.LsNodeEdgeResponse{}

	for _, document := range documents {
		lsNodeEdge := convertLSNodeEdge(document, request.PropertyNames)
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