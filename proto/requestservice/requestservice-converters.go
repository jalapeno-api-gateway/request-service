package requestservice

import (
	"gitlab.ost.ch/ins/jalapeno-api/request-service/influxdb"
	"gitlab.ost.ch/ins/jalapeno-api/request-service/redis"
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
			case Name: lsNode.Name = document.Name
			case Asn: lsNode.Asn = document.Asn
			case RouterIp: lsNode.RouterIp = document.Router_ip
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
			case RouterIp: lsLink.RouterIp = document.Router_ip
			case PeerIp: lsLink.PeerIp = document.Peer_ip
			case LocalLinkIp: lsLink.LocalLinkIp = document.LocalLink_ip
			case RemoteLinkIp: lsLink.RemoteLinkIp = document.RemoteLink_ip
			case IgpMetric: lsLink.IgpMetric = int32(document.Igp_metric)
		}
	}

	return &lsLink
}

func (lsLink *LsLink) setLsLinkProperty(document redis.LsLinkDocument, property string) {
	
}

func fetchTelemetryResponse(ipv4address string, propertyNames []string) *TelemetryResponse {
	response := &TelemetryResponse{Ipv4Address: ipv4address}
	for _, property := range propertyNames {
		response.setProperty(ipv4address, property)
	}
	return response
}

func (response *TelemetryResponse) setProperty(ipv4address string, property string) {
	switch property {
	case DataRate:
		response.setDataRate(ipv4address)
	}
}

func (response *TelemetryResponse) setDataRate(ipv4address string) {
	dataRate, err := influxdb.FetchDataRate(ipv4address)
	if err == nil {
		response.DataRate = dataRate
	}
}