package influxdb

import (
	"fmt"
	"os"

	influx "github.com/influxdata/influxdb1-client/v2"
	"github.com/sirupsen/logrus"
	// "github.com/jalapeno-api-gateway/jagw-core/logging"
)

var InfluxClient *influx.Client

func InitializeInfluxClient() {
	influxAddress := os.Getenv("INFLUX_ADDRESS")
	influxUser := os.Getenv("INFLUX_USER")
	influxPassword := os.Getenv("INFLUX_PASSWORD")

	logrus.WithFields(logrus.Fields{"influxAddress": influxAddress, "influxUser": influxUser}).Debug("Initializing Influx client.")

	client, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr:     fmt.Sprintf("http://%s", influxAddress),
		Username: influxUser,
		Password: influxPassword,
	})
	if err != nil {
		// TODO: Inform SR-App about unavailable
		logrus.WithError(err).Panic("Failed to create InfluxDB client.")
	}
	InfluxClient = &client
}

func queryInflux(queryString string) *influx.Response {
	logger := logrus.WithField("queryString", queryString)
	logger.Debug("Querying InfluxDB.")

	query := influx.NewQuery(queryString, os.Getenv("INFLUX_DB"), "")
	response, err := (*InfluxClient).Query(query)

	if err != nil {
		logger.WithError(err).Panic("Failed to query InfluxDB.")
	}
	if response.Error() != nil {
		logger.WithError(response.Error()).Panic("Failed to query InfluxDB.")
	}
	
	return response
}
