package requestservice

import (
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"github.com/jalapeno-api-gateway/request-service/influxdb"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func fetchTelemetryData(logger *logrus.Entry, request *jagw.TelemetryRequest) []string {
	jsonArray := influxdb.Fetch(logger, request)
	return jsonArray
}

func fetchMeasurements(logger *logrus.Entry) []*jagw.Measurement {
	measurementNames, err := influxdb.FetchMeasurements()
	if err != nil {
		logger.WithError(err).Panic("Failed to fetch measurements.")
	}

	measurements := []*jagw.Measurement{}
	for _, name := range measurementNames {
		m := jagw.Measurement{
			Name: proto.String(name),
			TimestampLatestMeasurement: proto.Int64(fetchLatestTimestamp(logger, name)),
		}
		measurements = append(measurements, &m)
	}

	return measurements
}

func fetchMeasurementColumns(logger *logrus.Entry, measurement string) []*jagw.MeasurementColumn {
	columns, err := influxdb.FetchColumns(measurement)
	if err != nil {
		logger.WithField("measurement", measurement).WithError(err).Panic("Failed to fetch measurement columns.")
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

func fetchLatestTimestamp(logger *logrus.Entry, measurement string) int64 {
	return influxdb.FetchTimestampOfLatestMeasurement(logger, measurement)
}
