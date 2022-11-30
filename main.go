package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cryptowizard0/notion2arweave/operator"
	"github.com/cryptowizard0/notion2arweave/service"
	"github.com/go-resty/resty/v2"
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

	service.StartServe()

	//samplePageTest()
}

func te() {
	client := resty.New()
	// url := "https://arseed.web3infra.dev/pjpGPYj_tpxd_43Y7TogDDGZup3m-HQE9IuycpvcRXY"
	//url := "https://images.unsplash.com/photo-1505740420928-5e560c06d30e?ixlib=rb-4.0.3&q=80&fm=jpg&crop=entropy&cs=tinysrgb"
	url := `https://s3.us-west-2.amazonaws.com/secure.notion-static.com/3fadede7-3dad-4228-85f7-02d927444a1a/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20221128%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20221128T100420Z&X-Amz-Expires=3600&X-Amz-Signature=06917c0fdd04c7a69b5adbe6662c95deb215ce60671197f646f41dbff449a369&X-Amz-SignedHeaders=host&x-id=GetObject`
	resp, err := client.R().Get(url)
	if err != nil {
		fmt.Println("get err: ", err.Error())
	}

	filePath := "./image.jpg"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("create file failed: ", err)
		return
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	nn, err := write.Write(resp.Body())
	if err != nil {
		fmt.Println("create file failed: ", err)
		return
	}

	fmt.Println("write success ", nn)
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
