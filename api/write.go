package api

import (
	"fmt"
	"time"

	influx "github.com/influxdata/influxdb/client/v2"
	"github.com/raptorbox/raptor-sdk-go/models"
	"github.com/spf13/viper"
)

// WriteRecord a batch of records
func WriteRecord(r []*models.Record) error {

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

		fields := record.Channels

		if record.UserID != "" {
			fields["userId"] = record.UserID
		}
		fields["streamId"] = record.StreamID
		fields["deviceId"] = record.DeviceID

		point, err1 := influx.NewPoint(
			record.DeviceID+"-"+record.GetStream().Name,
			map[string]string{},
			fields,
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
