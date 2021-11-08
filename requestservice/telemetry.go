package requestservice

import (
	"log"

	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"github.com/jalapeno-api-gateway/request-service/influxdb"
)

func fetchTelemetryData(request *jagw.TelemetryRequest) []string {
	jsonArray, err := influxdb.Fetch(request)
	if err != nil {
		log.Printf("%v\n", err)
	} 
	return jsonArray
}