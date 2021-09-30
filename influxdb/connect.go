package influxdb

import (
	"fmt"
	"log"
	"os"

	influx "github.com/influxdata/influxdb1-client/v2"
)

var InfluxClient *influx.Client

const INFLUX_DB_PORT = 30308

func InitializeInfluxClient() {
	client, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr:     fmt.Sprintf("http://%s:%d", os.Getenv("JALAPENO_SERVER"), INFLUX_DB_PORT),
		Username: os.Getenv("INFLUX_USER"),
		Password: os.Getenv("INFLUX_PASSWORD"),
	})
	if err != nil {
		//TODO: Inform SR-App about unavailable
		log.Fatal("Error creating InfluxDB Client: ", err.Error())
	}
	InfluxClient = &client
}

func queryInflux(queryString string) *influx.Response {
	query := influx.NewQuery(queryString, os.Getenv("INFLUX_DB"), "")
	response, err := (*InfluxClient).Query(query)

	if err != nil {
		log.Fatalf("Error querying InfluxDb: %v", err)
	}
	if response.Error() != nil {
		log.Fatalf("Error querying InfluxDb: %v", response.Error())
	}
	return response
}
