package requestservice

import (
	"github.com/jalapeno-api-gateway/model/property"
	"github.com/jalapeno-api-gateway/model/topology"
	"google.golang.org/protobuf/proto"
)

func convertLSNode(doc interface{}, propertyNames []string) *LSNode {
	document := doc.(topology.LSNode)
	lsNode := LSNode{Key: proto.String(document.Key)}

	if len(propertyNames) == 0 {
		propertyNames = property.AllLSNodeProperties
	}

	for _, propertyName := range propertyNames {
		switch propertyName {
			case property.Key: lsNode.Key = proto.String(document.Key)
			case property.ID: lsNode.ID = proto.String(document.ID)
			case property.RouterHash: lsNode.RouterHash = proto.String(document.RouterHash)
			case property.DomainID: lsNode.DomainID = proto.Int64(document.DomainID)
			case property.RouterIP: lsNode.RouterIP = proto.String(document.RouterIP)
			case property.PeerHash: lsNode.PeerHash = proto.String(document.PeerHash)
			case property.PeerIP: lsNode.PeerIP = proto.String(document.PeerIP)
			case property.PeerASN: lsNode.PeerASN = proto.Int32(document.PeerASN)
			case property.Timestamp: lsNode.Timestamp = proto.String(document.Timestamp)
			case property.IGPRouterID: lsNode.IGPRouterID = proto.String(document.IGPRouterID)
			case property.ASN: lsNode.ASN = proto.Uint32(document.ASN)
			case property.MTID: lsNode.MTID = convertMTIDSlice(document.MTID)
			case property.AreaID: lsNode.AreaID = proto.String(document.AreaID)
			case property.Protocol: lsNode.Protocol = proto.String(document.Protocol)
			case property.ProtocolID: lsNode.ProtocolID = proto.Uint32(uint32(document.ProtocolID))
			case property.Name: lsNode.Name = proto.String(document.Name)
			case property.IsPrepolicy: lsNode.IsPrepolicy = proto.Bool(document.IsPrepolicy)
			case property.IsAdjRIBIn: lsNode.IsAdjRIBIn = proto.Bool(document.IsAdjRIBIn)
		}
	}
	
	return &lsNode
}

func convertLSLink(doc interface{}, propertyNames []string) *LSLink {
	document := doc.(topology.LSLink)
	lsLink := LSLink{Key: proto.String(document.Key)}

	if len(propertyNames) == 0 {
		propertyNames = property.AllLSLinkProperties
	}

	for _, propertyName := range propertyNames {
		switch propertyName {
			case property.Key: lsLink.Key = proto.String(document.Key)
			case property.ID: lsLink.ID = proto.String(document.ID)
			case property.RouterHash: lsLink.RouterHash = proto.String(document.RouterHash)
			case property.RouterIP: lsLink.RouterIP = proto.String(document.RouterIP)
			case property.DomainID: lsLink.DomainID = proto.Int64(document.DomainID)
			case property.PeerHash: lsLink.PeerHash = proto.String(document.PeerHash)
			case property.PeerIP: lsLink.PeerIP = proto.String(document.PeerIP)
			case property.PeerASN: lsLink.PeerASN = proto.Int32(document.PeerASN)
			case property.Timestamp: lsLink.Timestamp = proto.String(document.Timestamp)
			case property.IGPRouterID: lsLink.IGPRouterID = proto.String(document.IGPRouterID)
			case property.Protocol: lsLink.Protocol = proto.String(document.Protocol)
			case property.AreaID: lsLink.AreaID = proto.String(document.AreaID)
			case property.Nexthop: lsLink.Nexthop = proto.String(document.Nexthop)
			case property.MTID: lsLink.MTID = convertMTID(document.MTID)
			case property.LocalLinkIP: lsLink.LocalLinkIP = proto.String(document.LocalLinkIP)
			case property.RemoteLinkIP: lsLink.RemoteLinkIP = proto.String(document.RemoteLinkIP)
			case property.IGPMetric: lsLink.IGPMetric = proto.Uint32(document.IGPMetric)
			case property.RemoteNodeHash: lsLink.RemoteNodeHash = proto.String(document.RemoteNodeHash)
			case property.LocalNodeHash: lsLink.LocalNodeHash = proto.String(document.LocalNodeHash)
			case property.RemoteIGPRouterID: lsLink.RemoteIGPRouterID = proto.String(document.RemoteIGPRouterID)
		}
	}

	return &lsLink
}

func convertLSPrefix(doc interface{}, propertyNames []string) *LSPrefix {
	document := doc.(*topology.LSPrefix)
	lsPrefix := LSPrefix{Key: proto.String(document.Key)}

	if len(propertyNames) == 0 {
		propertyNames = property.AllLSPrefixProperties
	}

	for _, propertyName := range propertyNames {
		switch propertyName {
			case property.Key: lsPrefix.Key = proto.String(document.Key)
			case property.ID: lsPrefix.ID = proto.String(document.ID)
			case property.RouterHash: lsPrefix.RouterHash = proto.String(document.RouterHash)
			case property.RouterIP: lsPrefix.RouterIP = proto.String(document.RouterIP)
			case property.DomainID: lsPrefix.DomainID = proto.Int64(document.DomainID)
			case property.PeerHash: lsPrefix.PeerHash = proto.String(document.PeerHash)
			case property.PeerIP: lsPrefix.PeerIP = proto.String(document.PeerIP)
			case property.PeerASN: lsPrefix.PeerASN = proto.Int32(document.PeerASN)
			case property.Timestamp: lsPrefix.Timestamp = proto.String(document.Timestamp)
			case property.IGPRouterID: lsPrefix.IGPRouterID = proto.String(document.IGPRouterID)
			case property.Protocol: lsPrefix.Protocol = proto.String(document.Protocol)
			case property.AreaID: lsPrefix.AreaID = proto.String(document.AreaID)
			case property.Nexthop: lsPrefix.Nexthop = proto.String(document.Nexthop)
			case property.LocalNodeHash: lsPrefix.LocalNodeHash = proto.String(document.LocalNodeHash)
			case property.MTID: lsPrefix.MTID = convertMTID(document.MTID)
			case property.Prefix: lsPrefix.Prefix = proto.String(document.Prefix)
			case property.PrefixLen: lsPrefix.PrefixLen = proto.Int32(document.PrefixLen)
			case property.PrefixMetric: lsPrefix.PrefixMetric = proto.Uint32(document.PrefixMetric)
			case property.IsPrepolicy: lsPrefix.IsPrepolicy = proto.Bool(document.IsPrepolicy)
			case property.IsAdjRIBIn: lsPrefix.IsAdjRIBIn = proto.Bool(document.IsAdjRIBIn)
		}
	}

	return &lsPrefix
}

func convertLSSRv6SID(doc interface{}, propertyNames []string) *LSSRv6SID {
	document := doc.(*topology.LSSRv6SID)
	lsSRv6SID := LSSRv6SID{Key: proto.String(document.Key)}

	if len(propertyNames) == 0 {
		propertyNames = property.AllLSLinkProperties
	}

	for _, propertyName := range propertyNames {
		switch propertyName {
			case property.Key: lsSRv6SID.Key = proto.String(document.Key)
			case property.ID: lsSRv6SID.ID = proto.String(document.ID)
			case property.RouterHash: lsSRv6SID.RouterHash = proto.String(document.RouterHash)
			case property.RouterIP: lsSRv6SID.RouterIP = proto.String(document.RouterIP)
			case property.DomainID: lsSRv6SID.DomainID = proto.Int64(document.DomainID)
			case property.PeerHash: lsSRv6SID.PeerHash = proto.String(document.PeerHash)
			case property.PeerIP: lsSRv6SID.PeerIP = proto.String(document.PeerIP)
			case property.PeerASN: lsSRv6SID.PeerASN = proto.Int32(document.PeerASN)
			case property.Timestamp: lsSRv6SID.Timestamp = proto.String(document.Timestamp)
			case property.IGPRouterID: lsSRv6SID.IGPRouterID = proto.String(document.IGPRouterID)
			case property.LocalNodeASN: lsSRv6SID.LocalNodeASN = proto.Uint32(document.LocalNodeASN)
			case property.Protocol: lsSRv6SID.Protocol = proto.String(document.Protocol)
			case property.Nexthop: lsSRv6SID.Nexthop = proto.String(document.Nexthop)
			case property.LocalNodeHash: lsSRv6SID.LocalNodeHash = proto.String(document.LocalNodeHash)
			case property.MTID: lsSRv6SID.MTID = convertMTID(document.MTID)
			case property.IGPFlags: lsSRv6SID.IGPFlags = proto.Uint32(uint32(document.IGPFlags))
			case property.IsPrepolicy: lsSRv6SID.IsPrepolicy = proto.Bool(document.IsPrepolicy)
			case property.IsAdjRIBIn: lsSRv6SID.IsAdjRIBIn = proto.Bool(document.IsAdjRIBIn)
			case property.SRv6SID: lsSRv6SID.SRv6SID = proto.String(document.SRv6SID)
		}
	}

	return &lsSRv6SID
}

func convertMTIDSlice(documents []*topology.MultiTopologyIdentifier) []*MultiTopologyIdentifier {
	mtids := []*MultiTopologyIdentifier{}
	for _, doc := range documents {
		mtids = append(mtids, convertMTID(doc))
	}
	return mtids
}

func convertMTID(doc *topology.MultiTopologyIdentifier) *MultiTopologyIdentifier {
	return &MultiTopologyIdentifier{
		OFlag: proto.Bool(doc.OFlag),
		AFlag: proto.Bool(doc.AFlag),
		MTID: proto.Uint32(uint32(doc.MTID)),
	}
}