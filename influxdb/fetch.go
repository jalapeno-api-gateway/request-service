package influxdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

const (
	InterfaceNameIdentifier           = "interface_name"
	IpAddressIdentifier               = "ip_information/ip_address"
	DataRateIdentifier                = "data_rates/output_data_rate"
	PacketsSentIdentifier             = "interface_statistics/full_interface_stats/packets_sent"
	PacketsReceivedIdentifier         = "interface_statistics/full_interface_stats/packets_received"
	StateIdentifier                   = "state"
	LastStateTransitionTimeIdentifier = "last_state_transition_time"
)

func FetchOutputDataRate(hostname string, linkId int32) (int64, error) {
	queryString := fmt.Sprintf("select last(\"%s\") FROM \"Cisco-IOS-XR-pfi-im-cmd-oper:interfaces/interface-xr/interface\" WHERE \"source\" = '%s' AND \"if_index\" = %d AND time < now() AND time > now() - 5m", DataRateIdentifier, hostname, linkId)
	return fetchInt64Value(queryString)
}

func FetchPacketsSent(hostname string, linkId int32) (int64, error) {
	queryString := fmt.Sprintf("select last(\"%s\") FROM \"Cisco-IOS-XR-pfi-im-cmd-oper:interfaces/interface-xr/interface\" WHERE \"source\" = '%s' AND \"if_index\" = %d AND time < now() AND time > now() - 5m", PacketsSentIdentifier, hostname, linkId)
	return fetchInt64Value(queryString)
}

func FetchPacketsReceived(hostname string, linkId int32) (int64, error) {
	queryString := fmt.Sprintf("select last(\"%s\") FROM \"Cisco-IOS-XR-pfi-im-cmd-oper:interfaces/interface-xr/interface\" WHERE \"source\" = '%s' AND \"if_index\" = %d AND time < now() AND time > now() - 5m", PacketsReceivedIdentifier, hostname, linkId)
	return fetchInt64Value(queryString)
}

func FetchState(hostname string, linkId int32) (string, error) {
	queryString := fmt.Sprintf("select last(\"%s\") FROM \"Cisco-IOS-XR-pfi-im-cmd-oper:interfaces/interface-xr/interface\" WHERE \"source\" = '%s' AND \"if_index\" = %d AND time < now() AND time > now() - 5m", StateIdentifier, hostname, linkId)
	return fetchStringValue(queryString)
}

func FetchLastStateTransitionTime(hostname string, linkId int32) (int64, error) {
	queryString := fmt.Sprintf("select last(\"%s\") FROM \"Cisco-IOS-XR-pfi-im-cmd-oper:interfaces/interface-xr/interface\" WHERE \"source\" = '%s' AND \"if_index\" = %d AND time < now() AND time > now() - 5m", LastStateTransitionTimeIdentifier, hostname, linkId)
	return fetchInt64Value(queryString)
}

func fetchStringValue(queryString string) (string, error) {
	response := queryInflux(queryString)
	
	if len(response.Results[0].Series) == 0 {
		return "", errors.New("no int64 value found for this ipv4address")
	}

	value := response.Results[0].Series[0].Values[0][1].(string)
	return value, nil
}

func fetchInt64Value(queryString string) (int64, error) {
	response := queryInflux(queryString)

	if len(response.Results[0].Series) == 0 {
		return 0, errors.New("no int64 value found for this ipv4address")
	}

	value, err := response.Results[0].Series[0].Values[0][1].(json.Number).Int64()
	if err != nil {
		log.Fatalf("Could not convert packets received to Int64")
	}
	return value, nil
}