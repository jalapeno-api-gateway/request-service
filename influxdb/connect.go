package influxdb

import (
	"log"
	"os"

	influx "github.com/influxdata/influxdb1-client/v2"
)

var InfluxClient *influx.Client

func InitializeInfluxClient() {
	client, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr:     os.Getenv("INFLUX_DB_URL"),
		Username: os.Getenv("INFLUX_USER"),
		Password: os.Getenv("INFLUX_PASSWORD"),
	})
	if err != nil {
		//TODO: Inform SR-App about unavailable
		log.Fatal("Error creating InfluxDB Client: ", err.Error())
	}
	InfluxClient = &client
}

func queryInflux(client influx.Client, queryString string) *influx.Response {
	query := influx.NewQuery(queryString, os.Getenv("INFLUX_DB"), "")
	response, err := client.Query(query)

	if err != nil {
		log.Fatalf("Error querying InfluxDb: %v", err)
	}
	if response.Error() != nil {
		log.Fatalf("Error querying InfluxDb: %v", response.Error())
	}
	return response
}
