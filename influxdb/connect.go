package influxdb

import (
	"fmt"
	"log"
	"os"

	influx "github.com/influxdata/influxdb1-client/v2"
)

func ConnectToInfluxDb() influx.Client {
	client, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr:     os.Getenv("INFLUX_DB_URL"),
		Username: os.Getenv("INFLUX_USER"),
		Password: os.Getenv("INFLUX_PASSWORD"),
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	return client
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