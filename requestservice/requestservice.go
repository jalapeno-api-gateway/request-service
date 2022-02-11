package requestservice

import (
	"context"
	"fmt"

	"github.com/jalapeno-api-gateway/jagw-core/model/class"
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"github.com/sirupsen/logrus"
)

type requestServiceServer struct {
	jagw.UnimplementedRequestServiceServer
}

func NewServer() *requestServiceServer {
	s := &requestServiceServer{}
	return s
}

func (s *requestServiceServer) GetLsNodes(ctx context.Context, request *jagw.TopologyRequest) (*jagw.LsNodeResponse, error) {
	logger := logrus.WithFields(logrus.Fields{"clientIp": getClientIp(ctx), "grpcFunction": "GetLsNodes"})
	logger.Debug("Incoming request.")

	documents := fetchDocuments(ctx, logger, request.Keys, class.LsNode)

	logger.Debug("Preparing response.")
	response := &jagw.LsNodeResponse{}

	for _, document := range documents {
		lsNode := convertLsNode(logger, document, request.PropertyNames)
		response.LsNodes = append(response.LsNodes, lsNode)
	}
	
	logger.Debug("Sending response.")
	return response, nil
}

func (s *requestServiceServer) GetLsNodeCoordinates(ctx context.Context, request *jagw.LsNodeCoordinatesRequest) (*jagw.LsNodeCoordinatesResponse, error) {
	logger := logrus.WithFields(logrus.Fields{"clientIp": getClientIp(ctx), "grpcFunction": "GetLsNodeCoordinates"})
	logger.Debug("Incoming request.")

	documentKeys := []string{}

	for _, lsNodeKey := range request.LsNodeKeys {
		documentKeys = append(documentKeys, fmt.Sprintf("%s_Coordinates", lsNodeKey))
	}
	
	documents := fetchDocuments(ctx, logger, documentKeys, class.LsNodeCoordinates)
	
	logger.Debug("Preparing response.")
	response := &jagw.LsNodeCoordinatesResponse{}

	for _, document := range documents {
		lsNode := convertLsNodeCoordinates(document)
		response.Coordinates = append(response.Coordinates, lsNode)
	}
	
	logger.Debug("Sending response.")
	return response, nil
}

func (s *requestServiceServer) GetLsLinks(ctx context.Context, request *jagw.TopologyRequest) (*jagw.LsLinkResponse, error) {
	logger := logrus.WithFields(logrus.Fields{"clientIp": getClientIp(ctx), "grpcFunction": "GetLsLinks"})
	logger.Debug("Incoming request.")

	documents := fetchDocuments(ctx, logger, request.Keys, class.LsLink)

	logger.Debug("Preparing response.")
	response := &jagw.LsLinkResponse{}

	for _, document := range documents {
		lsLink := convertLsLink(document, request.PropertyNames)
		response.LsLinks = append(response.LsLinks, lsLink)
	}
	
	logger.Debug("Sending response.")
	return response, nil
}

func (s *requestServiceServer) GetLsPrefixes(ctx context.Context, request *jagw.TopologyRequest) (*jagw.LsPrefixResponse, error) {
	logger := logrus.WithFields(logrus.Fields{"clientIp": getClientIp(ctx), "grpcFunction": "GetLsPrefixes"})
	logger.Debug("Incoming request.")

	documents := fetchDocuments(ctx, logger, request.Keys, class.LsPrefix)

	logger.Debug("Preparing response.")
	response := &jagw.LsPrefixResponse{}

	for _, document := range documents {
		lsPrefix := convertLsPrefix(document, request.PropertyNames)
		response.LsPrefixes = append(response.LsPrefixes, lsPrefix)
	}
	
	logger.Debug("Sending response.")
	return response, nil
}

func (s *requestServiceServer) GetLsSrv6Sids(ctx context.Context, request *jagw.TopologyRequest) (*jagw.LsSrv6SidResponse, error) {
	logger := logrus.WithFields(logrus.Fields{"clientIp": getClientIp(ctx), "grpcFunction": "GetLsSrv6Sids"})
	logger.Debug("Incoming request.")

	documents := fetchDocuments(ctx, logger, request.Keys, class.LsSrv6Sid)

	logger.Debug("Preparing response.")
	response := &jagw.LsSrv6SidResponse{}

	for _, document := range documents {
		lsSRv6SID := convertLsSrv6Sid(document, request.PropertyNames)
		response.LsSrv6Sids = append(response.LsSrv6Sids, lsSRv6SID)
	}
	
	logger.Debug("Sending response.")
	return response, nil
}

func (s *requestServiceServer) GetLsNodeEdges(ctx context.Context, request *jagw.TopologyRequest) (*jagw.LsNodeEdgeResponse, error) {
	logger := logrus.WithFields(logrus.Fields{"clientIp": getClientIp(ctx), "grpcFunction": "GetLsNodeEdges"})
	logger.Debug("Incoming request.")

	documents := fetchDocuments(ctx, logger, request.Keys, class.LsNodeEdge)

	logger.Debug("Preparing response.")
	response := &jagw.LsNodeEdgeResponse{}

	for _, document := range documents {
		lsNodeEdge := convertLsNodeEdge(document, request.PropertyNames)
		response.LsNodeEdges = append(response.LsNodeEdges, lsNodeEdge)
	}

	logger.Debug("Sending response.")
	return response, nil
}

func (s *requestServiceServer) GetTelemetryData(ctx context.Context, request *jagw.TelemetryRequest) (*jagw.TelemetryResponse, error) {
	logger := logrus.WithFields(logrus.Fields{"clientIp": getClientIp(ctx), "grpcFunction": "GetTelemetryData"})
	logger.Debug("Incoming request.")
	
	telemetryData := fetchTelemetryData(logger, request)
	response := &jagw.TelemetryResponse{TelemetryData: telemetryData}

	logger.Debug("Sending response.")
	return response, nil
}