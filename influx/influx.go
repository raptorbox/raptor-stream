package influx

import (
	"fmt"
	"time"

	influx "github.com/influxdata/influxdb/client/v2"
	"github.com/raptorbox/raptor-sdk-go/models"
	"github.com/spf13/viper"
)

var client influx.Client

//Record to store
type Record struct {
	Name   string
	Tags   map[string]string
	Values []interface{}
	Fields map[string]interface{}
	Time   time.Time
}

//Client return a influxdb client instance
func Client() error {
	if client != nil {
		return nil
	}
	var err error
	client, err = influx.NewHTTPClient(influx.HTTPConfig{
		Addr:     viper.GetString("influxdb.host"),
		Username: viper.GetString("influxdb.username"),
		Password: viper.GetString("influxdb.password"),
	})
	return err
}

// Write a batch of records
func Write(r []*models.Record) error {

	if len(r) == 0 {
		return nil
	}

	err := Client()
	if err != nil {
		return err
	}

	bp, err := influx.NewBatchPoints(influx.BatchPointsConfig{
		Database:  viper.GetString("influxdb.database"),
		Precision: viper.GetString("influxdb.precision"),
	})
	if err != nil {
		return err
	}

	for i := 0; i < len(r); i++ {

		record := r[i]

		if record.GetStream() == nil {
			return fmt.Errorf("Stream is empty")
		}

		point, err1 := influx.NewPoint(
			record.DeviceID+"-"+record.GetStream().Name,
			map[string]string{
				"streamId": record.StreamID,
				"deviceId": record.DeviceID,
				"userId":   record.UserID,
			},
			record.Channels,
			time.Unix(record.Timestamp, 0),
		)
		if err1 != nil {
			return err1
		}

		bp.AddPoint(point)
	}

	err = client.Write(bp)
	return err
}
