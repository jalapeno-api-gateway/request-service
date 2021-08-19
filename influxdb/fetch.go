package influxdb

import (
	"encoding/json"
	"fmt"
	"log"

	influx "github.com/influxdata/influxdb1-client/v2"
)

func FetchDataRates(client influx.Client, ipv4addresses []string) []DataRate {
	var dataRates []DataRate

	for _, ip := range ipv4addresses {
		queryString := fmt.Sprintf("select last(\"data_rates/output_data_rate\") FROM \"Cisco-IOS-XR-pfi-im-cmd-oper:interfaces/interface-xr/interface\" WHERE \"ip_information/ip_address\" = '%s'", ip)
		response := queryInflux(client, queryString)

		rawDataRate, err := response.Results[0].Series[0].Values[0][1].(json.Number).Int64()
		if err != nil {
			log.Fatalf("Could not convert DataRate to Int64")
		}
		dataRate := DataRate{Ipv4Address: ip, DataRate: rawDataRate}
		dataRates = append(dataRates, dataRate)
	}
	return dataRates
}
