package influxdb

import (
	"fmt"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"github.com/sirupsen/logrus"
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

func FetchMeasurements() ([]string, error) {
	queryString := "show measurements"
	response := queryInflux(queryString)

	measurements := []string{}
	for _, value := range response.Results[0].Series[0].Values {
		measurements = append(measurements, fmt.Sprintf("%v", value[0]))
	}

	return measurements, nil
}

func FetchColumns(measurement string) ([][]string, error) {
	tagQueryString := fmt.Sprintf("show tag keys from \"%s\"", measurement)
	fieldQueryString := fmt.Sprintf("show field keys from \"%s\"", measurement)

	tagResponse := queryInflux(tagQueryString)
	fieldResponse := queryInflux(fieldQueryString)

	columns := [][]string{{},{},{},{}}
	for _, value := range tagResponse.Results[0].Series[0].Values {
		columns[0] = append(columns[0], fmt.Sprintf("%v", value[0]))
		columns[1] = append(columns[1], "string")
		columns[2] = append(columns[2], "tag")
	}
	for _, value := range fieldResponse.Results[0].Series[0].Values {
		columns[0] = append(columns[0], fmt.Sprintf("%v", value[0]))
		columns[1] = append(columns[1], fmt.Sprintf("%v", value[1]))
		columns[2] = append(columns[2], "field")
	}

	latestMeasurementQueryString := fmt.Sprintf("select * from \"%s\" limit 1", measurement)
	latestMeasurementResponse := queryInflux(latestMeasurementQueryString)
	latestMeasurements := latestMeasurementResponse.Results[0].Series[0]

	columns[3] = make([]string, len(columns[0]))

	for i, property := range columns[0] {
		for j, measurementProperty := range latestMeasurements.Columns {
			if property == measurementProperty {
				if latestMeasurements.Values[0][j] == nil {
					columns[3][i] = ""
				} else if columns[1][i] == "string" {
						columns[3][i] = fmt.Sprintf("\"%v\"", latestMeasurements.Values[0][j])
				} else {
					if latestMeasurements.Values[0][j] == nil {
						columns[3][i] = fmt.Sprintf("%v", latestMeasurements.Values[0][j])
					}
				}
				break
			}
		}
	}

	return columns, nil
}

func FetchTimestampOfLatestMeasurement(logger *logrus.Entry, measurement string) int64 {
	fieldQueryString := fmt.Sprintf("show field keys from \"%s\"", measurement)
	fieldResponse := queryInflux(fieldQueryString)

	// For the request we have to provide at least one field name.
	// InfluxDb mandates at least one field per measurement, therefore we can be sure that the first value in the fieldResponse is set.
	firstField := fieldResponse.Results[0].Series[0].Values[0][0]
	queryString := fmt.Sprintf("select last(\"%s\") from \"%s\"", firstField, measurement)
	response := queryInflux(queryString)

	timestampString := fmt.Sprintf("%v", response.Results[0].Series[0].Values[0][0])
	timestampUnix, err := time.Parse("2006-01-02T15:04:05Z", timestampString)
	if err != nil {
		logger.WithError(err).Panic("Failed to parse timestamp.")
	}

	return timestampUnix.Unix()
}

func Fetch(logger *logrus.Entry, request *jagw.TelemetryRequest) []string {
	selection := formatSelection(request.Properties)
	filters := formatFilters(request)
	queryString := fmt.Sprintf("select %s FROM \"%s\" WHERE %s", selection, *request.SensorPath, filters)

	response := queryInflux(queryString)

	if len(response.Results[0].Series) == 0 {
		return []string{}
	}

	return createJSONArray(response)
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

/*
Response looks something like this:
{
  "Results": [
    {
      "statement_id": 0,
      "Series": [
        {
          "name": "Cisco-IOS-XR-pfi-im-cmd-oper:interfaces/interface-xr/interface",
          "columns": [
            "time",
            "data_rates/output_data_rate"
          ],
          "values": [
            [
              "2022-01-13T11:22:06.721Z",
              1
            ],
            [
              "2022-01-13T11:22:06.727Z",
              1
            ],
            [
              "2022-01-13T11:22:06.739Z",
              1
            ],
            [
              "2022-01-13T11:22:06.758Z",
              53
            ]
          ]
        }
      ],
      "Messages": null
    }
  ]
}
*/
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
