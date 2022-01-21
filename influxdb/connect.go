package influxdb

import (
	// "fmt"
	"log"
	// "os"

	influx "github.com/influxdata/influxdb1-client/v2"
)

var InfluxClient *influx.Client

func InitializeInfluxClient() {
	client, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr:     "http://10.20.1.24:30308",
		Username: "root",
		Password: "jalapeno",
		// Addr:     fmt.Sprintf("http://%s", os.Getenv("INFLUX_ADDRESS")),
		// Username: os.Getenv("INFLUX_USER"),
		// Password: os.Getenv("INFLUX_PASSWORD"),
	})
	if err != nil {
		//TODO: Inform SR-App about unavailable
		log.Fatal("Error creating InfluxDB Client: ", err.Error())
	}
	InfluxClient = &client
}

func queryInflux(queryString string) *influx.Response {
	query := influx.NewQuery(queryString, "mdt_db", "")
	// query := influx.NewQuery(queryString, os.Getenv("INFLUX_DB"), "")
	response, err := (*InfluxClient).Query(query)

	if err != nil {
		log.Fatalf("Error querying InfluxDb: %v", err)
	}
	if response.Error() != nil {
		log.Fatalf("Error querying InfluxDb: %v", response.Error())
	}
	return response
}
