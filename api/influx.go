package api

import (
	influx "github.com/influxdata/influxdb/client/v2"
	"github.com/spf13/viper"
)

var client influx.Client

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
