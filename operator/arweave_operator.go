package operator

import (
	"github.com/cryptowizard0/notion2arweave/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/everFinance/arseeding/sdk"
	"github.com/everFinance/arseeding/sdk/schema"
	"github.com/everFinance/goether"
)

type ArweaveOperator struct {
	ArseedSdk    *sdk.SDK
	ArseedClient *sdk.ArSeedCli
	PayCurrency  string
}

func CreateArweaveOperator(priKey, payCurrency string) *ArweaveOperator {
	eccSigner, err := goether.NewSigner(priKey)
	if err != nil {
		log.Error("create signer failed! Error:", err.Error())
		return nil
	}
	arseedSdk, err := sdk.NewSDK(viper.GetString("arweave.arseed_url"), viper.GetString("arweave.everpay_url"), eccSigner)
	if err != nil {
		log.Error("create arseed sdk failed! Error:", err.Error())
		return nil
	}
	client := sdk.New(viper.GetString("arweave.arseed_url"))
	return &ArweaveOperator{
		ArseedSdk:    arseedSdk,
		PayCurrency:  payCurrency,
		ArseedClient: client,
	}
}

// SavePage upload to arweave using arseeding
// @Param content, page content
// @Return txId, return by Arweave
func (a *ArweaveOperator) SavePage(content string) (txId string, err error) {
	log.Info("arweave operator: save_page")

	tags := utils.MakeTags("page", "TODO: sign a message", content)

	_, txId, err = a.ArseedSdk.SendDataAndPay([]byte(content), a.PayCurrency, &schema.OptionItem{Tags: tags}, false)
	if err != nil {
		return "", err
	}

	return
}

// LoadPage load a page content from arweave using arseeding
// @Param arTxId, tx id on arweave
// @Return content, return "" if content tag not found
func (a *ArweaveOperator) LoadPage(arTxId string) (content string, err error) {
	log.WithField("txid", arTxId).Info("arweave operator: load_page")

	item, err := a.ArseedClient.GetItemMeta(arTxId)
	if err != nil {
		return "", err
	}
	content = utils.GetTagValue("content", item.Tags)

	return
}
