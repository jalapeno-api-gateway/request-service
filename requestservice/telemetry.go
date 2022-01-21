package requestservice

import (
	// "encoding/json"
	// "fmt"
	"log"

	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"github.com/jalapeno-api-gateway/request-service/influxdb"
	"google.golang.org/protobuf/proto"
)

func fetchTelemetryData(request *jagw.TelemetryRequest) []string {
	jsonArray, err := influxdb.Fetch(request)
	if err != nil {
		log.Printf("%v\n", err)
	} 
	return jsonArray
}

func fetchMeasurements() []*jagw.Measurement {
	measurementNames, err := influxdb.FetchMeasurements()
	if err != nil {
		log.Printf("%v\n", err)
	}

	measurements := []*jagw.Measurement{}
	for _, name := range measurementNames {
		m := jagw.Measurement{
			Name: proto.String(name),
			TimestampLatestMeasurement: proto.Int64(fetchLatestTimestamp(name)),
		}
		measurements = append(measurements, &m)
	}

	return measurements
}

func fetchMeasurementColumns(measurement string) []*jagw.MeasurementColumn {
	columns, err := influxdb.FetchColumns(measurement)
	if err != nil {
		log.Printf("%v\n", err)
	}

	measurementcolumns := []*jagw.MeasurementColumn{}
	for i := 0; i < len(columns[0]); i++ {
		measurementcolumns = append(measurementcolumns, &jagw.MeasurementColumn{
			Name: &columns[0][i],
			Type: &columns[1][i],
			InfluxType: &columns[2][i],
			LastValueStringyfied: &columns[3][i],
		})
	}
	return measurementcolumns
}

func fetchLatestTimestamp(measurement string) int64 {
	latestMeasurement, err := influxdb.FetchTimestampOfLatestMeasurement(measurement)
	if err != nil {
		log.Printf("%v\n", err)
		return 0
	}
	return latestMeasurement
}