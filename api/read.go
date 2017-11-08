package api

import (
	"fmt"

	influx "github.com/influxdata/influxdb/client/v2"
	"github.com/spf13/viper"
)

//RecordQuery filter for reading data
type RecordQuery struct {
	DeviceID string
	StreamID string
}

//ToMap collect filter values
func (b *RecordQuery) ToMap() map[string]string {
	m := make(map[string]string)
	if b.DeviceID != "" {
		m["deviceId"] = b.DeviceID
	}
	if b.StreamID != "" {
		m["streamId"] = b.StreamID
	}
	return m
}

//Read values
func Read(baseFilter RecordQuery) (*influx.Response, error) {

	filter := ""
	i := 0
	for col, val := range baseFilter.ToMap() {
		if i > 0 {
			filter += ", "
		}
		filter += fmt.Sprintf("%s = '%s'", col, val)
		i++
	}

	query := fmt.Sprintf("select * from `%s` where %s",
		baseFilter.DeviceID+"-"+baseFilter.StreamID,
		filter)

	q := influx.Query{
		Command:  query,
		Database: viper.GetString("influxdb.database"),
	}

	return client.Query(q)
}
