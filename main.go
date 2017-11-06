package main

import (
	"log"

	"github.com/raptorbox/raptor-stream/server"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigType("json")
	viper.SetConfigName("stream")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s", err)
		return
	}

	server.Start(viper.GetString("port"))
}
