package influxdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

const (
	InterfaceNameIdentifier = "interface_name"
	IpAddressIdentifier = "ip_information/ip_address"
	DataRateIdentifier = "data_rates/output_data_rate"
	PacketsSentIdentifier = "interface_statistics/full_interface_stats/packets_sent"
	PacketsReceivedIdentifier = "interface_statistics/full_interface_stats/packets_received"
	StateIdentifier = "state"
	LastStateTransitionTimeIdentifier = "last_state_transition_time"
)

func FetchDataRate(ipv4address string) (int64, error) {
	queryString := fmt.Sprintf("select last(\"%s\") FROM \"Cisco-IOS-XR-pfi-im-cmd-oper:interfaces/interface-xr/interface\" WHERE \"ip_information/ip_address\" = '%s'", DataRateIdentifier, ipv4address)
	response := queryInflux(queryString)

	if len(response.Results[0].Series) == 0 {
		return 0, errors.New("no data rate found for this ipv4address")
	}

	rawDataRate, err := response.Results[0].Series[0].Values[0][1].(json.Number).Int64()
	if err != nil {
		log.Fatalf("Could not convert data rate to Int64")
	}
	return rawDataRate, nil
}
