package requestservice

import (
	"github.com/jalapeno-api-gateway/jagw-core/model/property"
	"github.com/jalapeno-api-gateway/jagw-core/model/topology"
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"google.golang.org/protobuf/proto"
)

func convertLSNode(doc interface{}, propertyNames []string) *jagw.LsNode {
	document := doc.(topology.LSNode)
	lsNode := jagw.LsNode{Key: proto.String(document.Key)}

	if len(propertyNames) == 0 {
		propertyNames = property.AllLSNodeProperties
	}

	for _, propertyName := range propertyNames {
		switch propertyName {
			case property.Key: lsNode.Key = proto.String(document.Key)
			case property.ID: lsNode.Id = proto.String(document.ID)
			case property.RouterHash: lsNode.RouterHash = proto.String(document.RouterHash)
			case property.DomainID: lsNode.DomainId = proto.Int64(document.DomainID)
			case property.RouterIP: lsNode.RouterIp = proto.String(document.RouterIP)
			case property.PeerHash: lsNode.PeerHash = proto.String(document.PeerHash)
			case property.PeerIP: lsNode.PeerIp = proto.String(document.PeerIP)
			case property.PeerASN: lsNode.PeerAsn = proto.Int32(document.PeerASN)
			case property.Timestamp: lsNode.Timestamp = proto.String(document.Timestamp)
			case property.IGPRouterID: lsNode.IgpRouterId = proto.String(document.IGPRouterID)
			case property.ASN: lsNode.Asn = proto.Uint32(document.ASN)
			case property.MTID: lsNode.Mtid = convertMTIDSlice(document.MTID)
			case property.AreaID: lsNode.AreaId = proto.String(document.AreaID)
			case property.Protocol: lsNode.Protocol = proto.String(document.Protocol)
			case property.ProtocolID: lsNode.ProtocolId = proto.Uint32(uint32(document.ProtocolID))
			case property.Name: lsNode.Name = proto.String(document.Name)
			case property.IsPrepolicy: lsNode.IsPrepolicy = proto.Bool(document.IsPrepolicy)
			case property.IsAdjRIBIn: lsNode.IsAdjRibIn = proto.Bool(document.IsAdjRIBIn)
		}
	}
	
	return &lsNode
}

func convertLSLink(doc interface{}, propertyNames []string) *jagw.LsLink {
	document := doc.(topology.LSLink)
	lsLink := jagw.LsLink{Key: proto.String(document.Key)}

	if len(propertyNames) == 0 {
		propertyNames = property.AllLSLinkProperties
	}

	for _, propertyName := range propertyNames {
		switch propertyName {
			case property.Key: lsLink.Key = proto.String(document.Key)
			case property.ID: lsLink.Id = proto.String(document.ID)
			case property.RouterHash: lsLink.RouterHash = proto.String(document.RouterHash)
			case property.RouterIP: lsLink.RouterIp = proto.String(document.RouterIP)
			case property.DomainID: lsLink.DomainId = proto.Int64(document.DomainID)
			case property.PeerHash: lsLink.PeerHash = proto.String(document.PeerHash)
			case property.PeerIP: lsLink.PeerIp = proto.String(document.PeerIP)
			case property.PeerASN: lsLink.PeerAsn = proto.Int32(document.PeerASN)
			case property.Timestamp: lsLink.Timestamp = proto.String(document.Timestamp)
			case property.IGPRouterID: lsLink.IgpRouterId = proto.String(document.IGPRouterID)
			case property.Protocol: lsLink.Protocol = proto.String(document.Protocol)
			case property.AreaID: lsLink.AreaId = proto.String(document.AreaID)
			case property.Nexthop: lsLink.Nexthop = proto.String(document.Nexthop)
			case property.MTID: lsLink.Mtid = convertMTID(document.MTID)
			case property.LocalLinkIP: lsLink.LocalLinkIp = proto.String(document.LocalLinkIP)
			case property.RemoteLinkIP: lsLink.RemoteLinkIp = proto.String(document.RemoteLinkIP)
			case property.IGPMetric: lsLink.IgpMetric = proto.Uint32(document.IGPMetric)
			case property.RemoteNodeHash: lsLink.RemoteNodeHash = proto.String(document.RemoteNodeHash)
			case property.LocalNodeHash: lsLink.LocalNodeHash = proto.String(document.LocalNodeHash)
			case property.RemoteIGPRouterID: lsLink.RemoteIgpRouterId = proto.String(document.RemoteIGPRouterID)
		}
	}

	return &lsLink
}

func convertLSPrefix(doc interface{}, propertyNames []string) *jagw.LsPrefix {
	document := doc.(*topology.LSPrefix)
	lsPrefix := jagw.LsPrefix{Key: proto.String(document.Key)}

	if len(propertyNames) == 0 {
		propertyNames = property.AllLSPrefixProperties
	}

	for _, propertyName := range propertyNames {
		switch propertyName {
			case property.Key: lsPrefix.Key = proto.String(document.Key)
			case property.ID: lsPrefix.Id = proto.String(document.ID)
			case property.RouterHash: lsPrefix.RouterHash = proto.String(document.RouterHash)
			case property.RouterIP: lsPrefix.RouterIp = proto.String(document.RouterIP)
			case property.DomainID: lsPrefix.DomainId = proto.Int64(document.DomainID)
			case property.PeerHash: lsPrefix.PeerHash = proto.String(document.PeerHash)
			case property.PeerIP: lsPrefix.PeerIp = proto.String(document.PeerIP)
			case property.PeerASN: lsPrefix.PeerAsn = proto.Int32(document.PeerASN)
			case property.Timestamp: lsPrefix.Timestamp = proto.String(document.Timestamp)
			case property.IGPRouterID: lsPrefix.IgpRouterId = proto.String(document.IGPRouterID)
			case property.Protocol: lsPrefix.Protocol = proto.String(document.Protocol)
			case property.AreaID: lsPrefix.AreaId = proto.String(document.AreaID)
			case property.Nexthop: lsPrefix.Nexthop = proto.String(document.Nexthop)
			case property.LocalNodeHash: lsPrefix.LocalNodeHash = proto.String(document.LocalNodeHash)
			case property.MTID: lsPrefix.Mtid = convertMTID(document.MTID)
			case property.Prefix: lsPrefix.Prefix = proto.String(document.Prefix)
			case property.PrefixLen: lsPrefix.PrefixLen = proto.Int32(document.PrefixLen)
			case property.PrefixMetric: lsPrefix.PrefixMetric = proto.Uint32(document.PrefixMetric)
			case property.IsPrepolicy: lsPrefix.IsPrepolicy = proto.Bool(document.IsPrepolicy)
			case property.IsAdjRIBIn: lsPrefix.IsAdjRibIn = proto.Bool(document.IsAdjRIBIn)
		}
	}

	return &lsPrefix
}

func convertLSSRv6SID(doc interface{}, propertyNames []string) *jagw.LsSrv6Sid {
	document := doc.(*topology.LSSRv6SID)
	lsSRv6SID := jagw.LsSrv6Sid{Key: proto.String(document.Key)}

	if len(propertyNames) == 0 {
		propertyNames = property.AllLSLinkProperties
	}

	for _, propertyName := range propertyNames {
		switch propertyName {
			case property.Key: lsSRv6SID.Key = proto.String(document.Key)
			case property.ID: lsSRv6SID.Id = proto.String(document.ID)
			case property.RouterHash: lsSRv6SID.RouterHash = proto.String(document.RouterHash)
			case property.RouterIP: lsSRv6SID.RouterIp = proto.String(document.RouterIP)
			case property.DomainID: lsSRv6SID.DomainId = proto.Int64(document.DomainID)
			case property.PeerHash: lsSRv6SID.PeerHash = proto.String(document.PeerHash)
			case property.PeerIP: lsSRv6SID.PeerIp = proto.String(document.PeerIP)
			case property.PeerASN: lsSRv6SID.PeerAsn = proto.Int32(document.PeerASN)
			case property.Timestamp: lsSRv6SID.Timestamp = proto.String(document.Timestamp)
			case property.IGPRouterID: lsSRv6SID.IgpRouterId = proto.String(document.IGPRouterID)
			case property.LocalNodeASN: lsSRv6SID.LocalNodeAsn = proto.Uint32(document.LocalNodeASN)
			case property.Protocol: lsSRv6SID.Protocol = proto.String(document.Protocol)
			case property.Nexthop: lsSRv6SID.Nexthop = proto.String(document.Nexthop)
			case property.LocalNodeHash: lsSRv6SID.LocalNodeHash = proto.String(document.LocalNodeHash)
			case property.MTID: lsSRv6SID.Mtid = convertMTID(document.MTID)
			case property.IGPFlags: lsSRv6SID.IgpFlags = proto.Uint32(uint32(document.IGPFlags))
			case property.IsPrepolicy: lsSRv6SID.IsPrepolicy = proto.Bool(document.IsPrepolicy)
			case property.IsAdjRIBIn: lsSRv6SID.IsAdjRibIn = proto.Bool(document.IsAdjRIBIn)
			case property.SRv6SID: lsSRv6SID.Srv6Sid = proto.String(document.SRv6SID)
		}
	}

	return &lsSRv6SID
}

func convertLSNodeEdge(doc interface{}, propertyNames []string) *jagw.LsNodeEdge {
	document := doc.(*topology.LSNodeEdge)
	lsNodeEdge := jagw.LsNodeEdge{Key: proto.String(document.Key)}

	if len(propertyNames) == 0 {
		propertyNames = property.AllLSLinkProperties
	}

	for _, propertyName := range propertyNames {
		switch propertyName {
			case property.Key: lsNodeEdge.Key = proto.String(document.Key)
			case property.ID: lsNodeEdge.Id = proto.String(document.ID)
			case property.From: lsNodeEdge.From = proto.String(document.From)
			case property.To: lsNodeEdge.To = proto.String(document.To)
			case property.Link: lsNodeEdge.Link = proto.String(document.Link)
		}
	}

	return &lsNodeEdge
}

func convertMTIDSlice(documents []*topology.MultiTopologyIdentifier) []*jagw.MultiTopologyIdentifier {
	mtids := []*jagw.MultiTopologyIdentifier{}
	for _, doc := range documents {
		mtids = append(mtids, convertMTID(doc))
	}
	return mtids
}

func convertMTID(doc *topology.MultiTopologyIdentifier) *jagw.MultiTopologyIdentifier {
	return &jagw.MultiTopologyIdentifier{
		OFlag: proto.Bool(doc.OFlag),
		AFlag: proto.Bool(doc.AFlag),
		Mtid: proto.Uint32(uint32(doc.MTID)),
	}
}