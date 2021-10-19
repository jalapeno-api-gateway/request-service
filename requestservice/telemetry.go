package requestservice

import (
	"github.com/jalapeno-api-gateway/jagw-core/model/property"
	"github.com/jalapeno-api-gateway/request-service/influxdb"
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"google.golang.org/protobuf/proto"
)

func fetchTelemetryData(interfaceId *jagw.InterfaceIdentifier, propertyNames []string) *jagw.TelemetryData {
	response := &jagw.TelemetryData{InterfaceId: interfaceId}

	if len(propertyNames) == 0 { // If no propertyNames were specified by the SR-App, return all
		propertyNames = append(property.AllPhysicalInterfaceProperties, property.AllLoopbackInterfaceProperties...)
	}

	for _, property := range propertyNames {
		setProperty(response, interfaceId, property)
	}
	return response
}

func setProperty(response *jagw.TelemetryData, interfaceId *jagw.InterfaceIdentifier, propertyName string) {
	switch propertyName {
		case property.DataRate: setDataRate(response, interfaceId)
		case property.PacketsSent: setPacketsSent(response, interfaceId)
		case property.PacketsReceived: setPacketsReceived(response, interfaceId)
		case property.State: setState(response, interfaceId)
		case property.LastStateTransitionTime: setLastStateTransitionTime(response, interfaceId)
	}
}

func setDataRate(response *jagw.TelemetryData, interfaceId *jagw.InterfaceIdentifier) {
	value, err := influxdb.FetchOutputDataRate(*interfaceId.Hostname, *interfaceId.LinkId)
	if err == nil {
		response.DataRate = proto.Int64(value)
	}
}

func setPacketsSent(response *jagw.TelemetryData, interfaceId *jagw.InterfaceIdentifier) {
	value, err := influxdb.FetchPacketsSent(*interfaceId.Hostname, *interfaceId.LinkId)
	if err == nil {
		response.PacketsSent = proto.Int64(value)
	}
}

func setPacketsReceived(response *jagw.TelemetryData, interfaceId *jagw.InterfaceIdentifier) {
	value, err := influxdb.FetchPacketsReceived(*interfaceId.Hostname, *interfaceId.LinkId)
	if err == nil {
		response.PacketsReceived = proto.Int64(value)
	}
}

func setState(response *jagw.TelemetryData, interfaceId *jagw.InterfaceIdentifier) {
	value, err := influxdb.FetchState(*interfaceId.Hostname, *interfaceId.LinkId)
	if err == nil {
		response.State = proto.String(value)
	}
}

func setLastStateTransitionTime(response *jagw.TelemetryData, interfaceId *jagw.InterfaceIdentifier) {
	value, err := influxdb.FetchLastStateTransitionTime(*interfaceId.Hostname, *interfaceId.LinkId)
	if err == nil {
		response.LastStateTransitionTime = proto.Int64(value)
	}
}