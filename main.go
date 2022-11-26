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

	// samplePageTest()
	imgTest()
}

func samplePageTest() {
	opt := operator.CreateOperator(
		viper.GetString("notion.api_auth"),
		viper.GetString("arweave.pk"),
		"USDC")
	if opt == nil {
		fmt.Println("Error, create operator failed!")
		return
	}

	arTxId, err := opt.SavePage2Ar("8f7937d345e84645b3b8580dc138e7d2")
	if err != nil {
		fmt.Println("Error, SavePage2Ar! ", err.Error())
	}
	log.Info("Save 2 ar success: ", arTxId)

	uuid, err := opt.LoadPageFromAr("c904d90c9abf4de68f7520786193d4c0", arTxId)
	if err != nil {
		fmt.Println("Error, upload failed! ", err.Error())
		return
	}

	log.WithField("uuid", uuid).Info("Success!")
}

func imgTest() {
	opt := operator.CreateOperator(
		viper.GetString("notion.api_auth"),
		viper.GetString("arweave.pk"),
		"USDC")
	if opt == nil {
		fmt.Println("Error, create operator failed!")
		return
	}

	arTxId, err := opt.SavePage2Ar("2316fe9dade64ffbb5aa45b46f069dbf")
	if err != nil {
		fmt.Println("Error, SavePage2Ar! ", err.Error())
	}
	log.Info("Save 2 ar success: ", arTxId)

	uuid, err := opt.LoadPageFromAr("c904d90c9abf4de68f7520786193d4c0", arTxId)
	if err != nil {
		fmt.Println("Error, upload failed! ", err.Error())
		return
	}

	log.WithField("uuid", uuid).Info("Success!")
}
