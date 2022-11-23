package operator

import (
	arTypes "github.com/everFinance/goar/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/everFinance/arseeding/sdk"
	"github.com/everFinance/arseeding/sdk/schema"
	"github.com/everFinance/goether"
)

var BaseTag = []arTypes.Tag{
	{Name: "Content-Type", Value: "application/json"},
	{Name: "App-Name", Value: "notion2arweawve"},
	{Name: "App-Version", Value: "0.1.0"},
}

type ArweaveOperator struct {
	ArseedSdk   sdk.SDK
	PayCurrency string
}

func CreateArweaveOperator(priKey, payCurrency string) (*ArweaveOperator, error) {
	eccSigner, err := goether.NewSigner(priKey)
	if err != nil {
		return nil, err
	}
	sdk, err := sdk.NewSDK(viper.GetString("arweave.arseed_url"), viper.GetString("arweave.everpay_url"), eccSigner)
	if err != nil {
		return nil, err
	}
	return &ArweaveOperator{
		ArseedSdk:   *sdk,
		PayCurrency: payCurrency,
	}, nil
}

// SavePage upload to arweave using arseeding
// @Pararm content, page content
// @Return txId, return by Arweave
func (a *ArweaveOperator) SavePage(content string) (txId string, err error) {
	log.Info("save_page")

	tags := []arTypes.Tag{}
	tags = append(tags, BaseTag...)
	tags = append(tags, []arTypes.Tag{
		{Name: "type", Value: "page"},
		{Name: "sign", Value: "sign a message"},
		{Name: "content", Value: content},
	}...)

	_, txId, err = a.ArseedSdk.SendDataAndPay([]byte(content), a.PayCurrency, &schema.OptionItem{Tags: tags}, false)
	if err != nil {
		return "", err
	}

	return
}
