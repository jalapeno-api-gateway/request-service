package requestservice

import (
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"github.com/jalapeno-api-gateway/request-service/influxdb"
	"github.com/sirupsen/logrus"
)

func fetchTelemetryData(logger *logrus.Entry, request *jagw.TelemetryRequest) []string {
	jsonArray := influxdb.Fetch(logger, request)
	return jsonArray
}