package operator

import (
	"encoding/json"
	"fmt"

	"github.com/cryptowizard0/go-notion"
	"github.com/cryptowizard0/notion2arweave/types"
	"github.com/cryptowizard0/notion2arweave/utils"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/everFinance/arseeding/sdk"
	"github.com/everFinance/arseeding/sdk/schema"
	"github.com/everFinance/goether"
)

type ArweaveOperator struct {
	ArseedSdk   *sdk.SDK
	PayCurrency string
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

	return &ArweaveOperator{
		ArseedSdk:   arseedSdk,
		PayCurrency: payCurrency,
	}
}

// SavePage upload page to arweave using arseeding
// @Param content, page content
// @Return txId, return by Arweave
func (a *ArweaveOperator) SavePage(content string) (txId string, err error) {
	log.Info("arweave operator: save page to arweave")

	page, err := a.filterChildBlocks(content)
	if err != nil {
		return "", err
	}
	bContent, err := json.Marshal(page)
	if err != nil {
		return "", err
	}

	// bContent := []byte(content)
	tags := utils.MakeTags("page", "TODO: sign a message", string(bContent))

	_, txId, err = a.ArseedSdk.SendDataAndPay(bContent, a.PayCurrency, &schema.OptionItem{Tags: tags}, false)
	if err != nil {
		return "", err
	}

	return
}

// LoadPage load a page content from arweave using arseeding
// @Param arTxId, tx id on arweave
// @Return content, return "" if content tag not found
func (a *ArweaveOperator) LoadPage(arTxId string) (content string, err error) {
	log.WithField("txid", arTxId).Info("arweave operator: load page from arweave")

	item, err := a.ArseedSdk.Cli.GetItemMeta(arTxId)
	if err != nil {
		return "", err
	}
	content = utils.GetTagValue("content", item.Tags)

	return
}

//=================================================
func (a *ArweaveOperator) filterChildBlocks(srcContent string) (*types.ArweavePage, error) {
	var page types.ArweavePage
	err := json.Unmarshal([]byte(srcContent), &page)
	if err != nil {
		return nil, err
	}
	log.Info("Blocks count: ", len(page.PageContent.Results))
	var tmpBlocks []notion.Block

	for _, block := range page.PageContent.Results {
		dtoblock, ok := block.(notion.BlockDTO)
		if !ok {
			log.Errorf("input block is not notion.BlockDTO: %#v", block)
			return nil, fmt.Errorf("Convert to BlockDTO failed")
		}

		if dtoblock.Image != nil {
			imageBlock := dtoblock.Image

			log.Infof("image block: %#v", imageBlock)
			var url string
			if imageBlock.Type == notion.FileTypeExternal {
				url = imageBlock.External.URL
			} else {
				url = imageBlock.File.URL
			}

			arTxId, err := a.saveImage(url)
			if err != nil {
				log.Error("save image to arweave failed! error: ", err)
				tmpBlocks = append(tmpBlocks, block)
				continue
			}
			if imageBlock.Type == notion.FileTypeExternal {
				imageBlock.External.URL = "https://arseed.web3infra.dev/" + arTxId
			} else {
				imageBlock.File.URL = "https://arseed.web3infra.dev/" + arTxId
			}

		}
		tmpBlocks = append(tmpBlocks, dtoblock)
	}
	page.PageContent.Results = tmpBlocks

	return &page, nil
}

func (a *ArweaveOperator) saveImage(url string) (arTxId string, err error) {
	client := resty.New()
	resp, err := client.R().Get(url)
	if err != nil {
		return "", err
	}
	image := resp.Body()

	tags := utils.MakeImageTags("TODO: sign a message")

	_, arTxId, err = a.ArseedSdk.SendDataAndPay(image, a.PayCurrency, &schema.OptionItem{Tags: tags}, false)
	if err != nil {
		return "", err
	}

	return arTxId, nil
}
