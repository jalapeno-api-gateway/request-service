package requestservice

import (
	"github.com/jalapeno-api-gateway/model/property"
	"github.com/jalapeno-api-gateway/request-service/influxdb"
	"google.golang.org/protobuf/proto"
)

func fetchTelemetryData(ipv4address string, propertyNames []string) *TelemetryData {
	response := &TelemetryData{Ipv4Address: proto.String(ipv4address)}
	for _, property := range propertyNames {
		response.setProperty(ipv4address, property)
	}
	return response
}

func (response *TelemetryData) setProperty(ipv4address string, propertyName string) {
	switch propertyName {
	case property.DataRate:
		response.setDataRate(ipv4address)
	}
}

func (response *TelemetryData) setDataRate(ipv4address string) {
	dataRate, err := influxdb.FetchDataRate(ipv4address)
	if err == nil {
		response.DataRate = proto.Int64(dataRate)
	}
}