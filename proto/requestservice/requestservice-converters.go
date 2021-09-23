package requestservice

import (
	"github.com/jalapeno-api-gateway/request-service/influxdb"
	"github.com/jalapeno-api-gateway/request-service/redis"
	"google.golang.org/protobuf/proto"
)

// Property Names
const (
	DataRate = "DataRate"
	Name = "Name"
	Asn = "Asn"
	RouterIp = "RouterIp"
	PeerIp = "PeerIp"
	LocalLinkIp = "LocalLinkIp"
	RemoteLinkIp = "RemoteLinkIp"
	IgpMetric = "IgpMetric"
)

var allLsNodeProperties = []string{
	Name,
	Asn,
	RouterIp,
}

var allLsLinkProperties = []string{
	RouterIp,
	PeerIp,
	LocalLinkIp,
	RemoteLinkIp,
	IgpMetric,
}

func convertLsNode(document redis.LsNodeDocument, propertyNames []string) *LsNode {
	lsNode := LsNode{Key: document.Key}

	if len(propertyNames) == 0 {
		propertyNames = allLsNodeProperties
	}

	for _, property := range propertyNames {
		switch property {
			case Name: lsNode.Name = proto.String(document.Name)
			case Asn: lsNode.Asn = proto.Int32(document.Asn)
			case RouterIp: lsNode.RouterIp = proto.String(document.Router_ip)
		}
	}
	
	return &lsNode
}


func convertLsLink(document redis.LsLinkDocument, propertyNames []string) *LsLink {
	lsLink := LsLink{Key: document.Key}

	if len(propertyNames) == 0 {
		propertyNames = allLsLinkProperties
	}

	for _, property := range propertyNames {
		switch property {
			case RouterIp: lsLink.RouterIp = proto.String(document.Router_ip)
			case PeerIp: lsLink.PeerIp = proto.String(document.Peer_ip)
			case LocalLinkIp: lsLink.LocalLinkIp = proto.String(document.LocalLink_ip)
			case RemoteLinkIp: lsLink.RemoteLinkIp = proto.String(document.RemoteLink_ip)
			case IgpMetric: lsLink.IgpMetric = proto.Int32(int32(document.Igp_metric))
		}
	}

	return &lsLink
}

func fetchTelemetryData(ipv4address string, propertyNames []string) *TelemetryData {
	response := &TelemetryData{Ipv4Address: ipv4address}
	for _, property := range propertyNames {
		response.setProperty(ipv4address, property)
	}
	return response
}

func (response *TelemetryData) setProperty(ipv4address string, property string) {
	switch property {
	case DataRate:
		response.setDataRate(ipv4address)
	}
}

func (response *TelemetryData) setDataRate(ipv4address string) {
	dataRate, err := influxdb.FetchDataRate(ipv4address)
	if err == nil {
		response.DataRate = proto.Int64(dataRate)
	}
}