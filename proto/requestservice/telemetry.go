package requestservice

import (
	"github.com/jalapeno-api-gateway/jagw-core/model/property"
	"github.com/jalapeno-api-gateway/request-service/influxdb"
	"google.golang.org/protobuf/proto"
)

func fetchTelemetryData(ipv4address string, propertyNames []string) *TelemetryData {
	response := &TelemetryData{Ipv4Address: proto.String(ipv4address)}

	if len(propertyNames) == 0 { // If no propertyNames were specified by the SR-App, return all
		propertyNames = append(property.AllPhysicalInterfaceProperties, property.AllLoopbackInterfaceProperties...)
	}

	for _, property := range propertyNames {
		response.setProperty(ipv4address, property)
	}
	return response
}

func (response *TelemetryData) setProperty(ipv4address string, propertyName string) {
	switch propertyName {
		case property.DataRate: response.setDataRate(ipv4address)
		case property.PacketsSent: response.setPacketsSent(ipv4address)
		case property.PacketsReceived: response.setPacketsReceived(ipv4address)
		case property.State: response.setState(ipv4address)
		case property.LastStateTransitionTime: response.setLastStateTransitionTime(ipv4address)
	}
}

func (response *TelemetryData) setDataRate(ipv4address string) {
	value, err := influxdb.FetchOutputDataRate(ipv4address)
	if err == nil {
		response.DataRate = proto.Int64(value)
	}
}

func (response *TelemetryData) setPacketsSent(ipv4address string) {
	value, err := influxdb.FetchPacketsSent(ipv4address)
	if err == nil {
		response.PacketsSent = proto.Int64(value)
	}
}

func (response *TelemetryData) setPacketsReceived(ipv4address string) {
	value, err := influxdb.FetchPacketsReceived(ipv4address)
	if err == nil {
		response.PacketsReceived = proto.Int64(value)
	}
}

func (response *TelemetryData) setState(ipv4address string) {
	value, err := influxdb.FetchState(ipv4address)
	if err == nil {
		response.State = proto.String(value)
	}
}

func (response *TelemetryData) setLastStateTransitionTime(ipv4address string) {
	value, err := influxdb.FetchLastStateTransitionTime(ipv4address)
	if err == nil {
		response.LastStateTransitionTime = proto.Int64(value)
	}
}