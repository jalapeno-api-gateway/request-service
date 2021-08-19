package requestservice

import (
	"gitlab.ost.ch/ins/jalapeno-api/request-service/influxdb"
	"gitlab.ost.ch/ins/jalapeno-api/request-service/redis"
)

func convertToGrpcDataRate(dataRate influxdb.DataRate) DataRate {
	return DataRate{
		Ipv4Address: dataRate.Ipv4Address,
		DataRate: dataRate.DataRate,
	}
}

func convertToGrpcLsNode(nodeDocument redis.LsNodeDocument) LsNode {
	return LsNode{
		Key: nodeDocument.Key,
		Name: nodeDocument.Name,
		Asn: nodeDocument.Asn,
		RouterIp: nodeDocument.Router_ip,
	}
}

func convertToGrpcLsLink(linkDocument redis.LsLinkDocument) LsLink {
	return LsLink{
		Key: linkDocument.Key,
		RouterIp: linkDocument.Router_ip,
		PeerIp: linkDocument.Peer_ip,
		LocalLinkIp: linkDocument.LocalLink_ip,
		RemoteLinkIp: linkDocument.RemoteLink_ip,
		IgpMetric: int32(linkDocument.Igp_metric),
	}
}