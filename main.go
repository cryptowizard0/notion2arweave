package main

import (
	"fmt"

	"github.com/cryptowizard0/notion2arweave/service"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	// Read configs
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("read config failed: %s", err.Error()))
	}

	// Init log
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(log.DebugLevel)
}

func main() {
	fmt.Println("=======================")
	fmt.Println("Hello notion 2 arweave!")
	fmt.Println("=======================")

	service.StartServe()
}
