package main

import (
	"fmt"

	"github.com/cryptowizard0/notion2arweave/operator"
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
	fmt.Println("Hello notion 2 arweave!")

	queryUseOPerator()
}

func queryUseOPerator() {
	opt := operator.CreateNotionOperator(viper.GetString("notion.api_auth"))
	txId, err := opt.FetchPage("c904d90c9abf4de68f7520786193d4c0")
	if err != nil {
		fmt.Println(err.Error())
	}

	log.WithField("txId", txId).Info("fetch_page OK!")
}
