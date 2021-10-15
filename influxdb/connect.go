package influxdb

import (
	"fmt"
	"log"
	"os"
	"time"

	influx "github.com/influxdata/influxdb1-client/v2"
)

var InfluxClient *influx.Client

func InitializeInfluxClient() {
	log.Printf("InitializeInfluxClient - INFLUX_ADDRESS: %s", os.Getenv("INFLUX_ADDRESS"))
	log.Printf("InitializeInfluxClient - INFLUX_USER: %s", os.Getenv("INFLUX_USER"))
	log.Printf("InitializeInfluxClient - INFLUX_PASSWORD: %s", os.Getenv("INFLUX_PASSWORD"))
	log.Printf("InitializeInfluxClient - INFLUX_DB: %s", os.Getenv("INFLUX_DB"))

	client, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr:     fmt.Sprintf("http://%s", os.Getenv("INFLUX_ADDRESS")),
		Username: os.Getenv("INFLUX_USER"),
		Password: os.Getenv("INFLUX_PASSWORD"),
	})
	if err != nil {
		//TODO: Inform SR-App about unavailable
		log.Fatal("Error creating InfluxDB Client: ", err.Error())
	}

	_, _, err = client.Ping(time.Second * 10)

	if err != nil {
		log.Printf("Could not ping: %s", err)
	}

	InfluxClient = &client
}

func queryInflux(queryString string) *influx.Response {
	log.Printf("queryInflux - INFLUX_DB: %s", os.Getenv("INFLUX_DB"))

	query := influx.NewQuery(queryString, os.Getenv("INFLUX_DB"), "")
	response, err := (*InfluxClient).Query(query)

	if err != nil {
		log.Fatalf("Error querying InfluxDb: %v", err)
	}
	if response.Error() != nil {
		log.Fatalf("(Response-)Error querying InfluxDb: %v", response.Error())
	}
	return response
}
