package influxdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"github.com/iancoleman/strcase"
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

func Fetch(request *jagw.TelemetryRequest) ([]string, error) {
	selection := formatSelection(request.Properties)
	filters := formatFilters(request)
	log.Printf("Selection: %s", selection)
	log.Printf("SensorPath: %s", *request.SensorPath)
	log.Printf("Filters: %s", filters)
	queryString := fmt.Sprintf("select %s FROM \"%s\" WHERE %s", selection, *request.SensorPath, filters)

	response := queryInflux(queryString)

	if len(response.Results[0].Series) == 0 {
		return []string{}, errors.New("error 1")
	}

	s, _ := json.MarshalIndent(response, "", "  ")
	fmt.Printf("%s\n\n", string(s))

	return createJSONArray(response), nil
}

func formatSelection(properties []string) string {
	var b strings.Builder
	for i, property := range properties {
		b.Reset()
		if property == "*" {
			b.WriteString(property)
		} else if strings.Contains(property, "*") {
			// TODO This is not allowed, return error or something
		} else {
			b.WriteString("\"")
			b.WriteString(property)
			b.WriteString("\"")
		}
		properties[i] = b.String()
	}
	return strings.Join(properties, ", ")
}

func formatFilters(request *jagw.TelemetryRequest) string {
	var b strings.Builder
	formatStringFilters(&b, request.StringFilters)
	
	if request.RangeFilter == nil {
		trimmed := removeTrailingCharacters(b.String(), 5) // remove trailing instance of " AND "
		return trimmed + " limit 1"
	} else {
		formatRangeFilter(&b, request.RangeFilter)
		return b.String()
	}
}

func formatStringFilters(b *strings.Builder, stringFilters []*jagw.StringFilter) {
	for _, stringFilter := range stringFilters {
		b.WriteString("\"")
		b.WriteString(*stringFilter.Property)
		b.WriteString("\"")
		switch *stringFilter.Operator {
			case jagw.StringOperator_EQUAL: b.WriteString(" = ")
			case jagw.StringOperator_NOT_EQUAL: b.WriteString(" != ")
		}
		b.WriteString("'")
		b.WriteString(*stringFilter.Value)
		b.WriteString("'")
		b.WriteString(" AND ")
	}
}

func formatRangeFilter(b *strings.Builder, rangeFilter *jagw.RangeFilter) {
	b.WriteString("time >= ")
	fmt.Fprintf(b, "%d", *rangeFilter.EarliestTimestamp)
	b.WriteString(" AND ")
	b.WriteString("time <= ")
	if rangeFilter.LatestTimestamp == nil {
		b.WriteString("now()")
	} else {
		fmt.Fprintf(b, "%d", *rangeFilter.LatestTimestamp)
	}
}

func createJSONArray(response *client.Response) []string {
	series := response.Results[0].Series[0]

	formattedPropertyNames := make([]string, len(series.Columns))
	for i, property := range series.Columns {
		formattedPropertyNames[i] = formatPropertyName(property)
	}

	jsonArray := make([]string, len(series.Values))
	for i := 0; i < len(jsonArray); i++ {
		jsonArray[i] = createSingleJSON(formattedPropertyNames, series.Values[i])
	}

	return jsonArray
}

func createSingleJSON(formattedPropertyNames []string, values []interface{}) string {
	var b strings.Builder
	b.WriteString("{")
	
	for i := 0; i < len(formattedPropertyNames); i++ {		
		if values[i] != nil {
			b.WriteString("\"")
			b.WriteString(formattedPropertyNames[i])
			b.WriteString("\"")
			b.WriteString(": ")
			switch values[i].(type) {
				case string: fmt.Fprintf(&b, "\"%v\"", values[i])
				default: fmt.Fprintf(&b, "%v", values[i])
			}
			b.WriteString(", ")
		} else {
			// TODO Handle case "No value for this property"
		}
	}
	
	trimmed := removeTrailingCharacters(b.String(), 2) // Remove trailing ", "
	log.Printf("%v}\n", trimmed)
	return trimmed + "}"
}

func removeTrailingCharacters(s string, numberOfCharacters int) string {
	if len(s) > 0 {
		s = s[:len(s) - numberOfCharacters]
	}
	return s
}

/*
Converts a string in the format of:
   "data_rates/output_data_rate"
to:
   "DataRates_OutputDataRate"
*/
func formatPropertyName(propertyName string) string {
	names := strings.Split(propertyName, "/")
	var b strings.Builder
	
	lastIndex := len(names) -1
	for i, name := range names {
		b.WriteString(strcase.ToCamel(name))
		if i < lastIndex {
			b.WriteString("_")
		}
	}

	return b.String()
}
