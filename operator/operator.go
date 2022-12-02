package operator

import (
	"fmt"

	"github.com/cryptowizard0/notion2arweave/utils"
	log "github.com/sirupsen/logrus"
)

// Operator inherite IOperator
type Operator struct {
	NotionOpter *NotionOperator
	ArOpter     *ArweaveOperator
}

func CreateOperator(notionAuth, priKey, payCurrency string) *Operator {
	return &Operator{
		NotionOpter: CreateNotionOperator(notionAuth),
		ArOpter:     CreateArweaveOperator(priKey, payCurrency),
	}
}

// SavePage2Ar fetch page from notion and upload to arweave
func (o *Operator) SavePage2Ar(uuid string) (arTxId string, err error) {
	log.WithField("uuid", uuid).Info("operator: save page to arweave")

	// get page content
	content, err := o.NotionOpter.FetchPage(uuid)
	if err != nil {
		log.WithField("uuid", uuid).Error("fetch page content error: ", err.Error())
		return "", err
	}

	// upload to arweave
	arTxId, err = o.ArOpter.SavePage(content)
	if err != nil {
		return "", err
	}
	log.WithField("txid", arTxId).Info("save page to arweave success")
	return
}

// LoadPageFromAr get page from arweave and upload to notion
func (o *Operator) LoadPageFromAr(parentId, arTxId string) (uuid string, err error) {
	log.WithField("arTxId", arTxId).WithField("parent", parentId).Info("operator: load page from arweave")

	// load content from arseeding
	content, err := o.ArOpter.LoadPage(arTxId)
	if err != nil {
		log.Error("get content from arweave error: ", err.Error())
		return "", err
	}
	if content == "" {
		log.WithField("arTxId", arTxId).Error("Can't find content within ArTx")
		return "", fmt.Errorf("can't find content within ArTx: %s", arTxId)
	}
	log.WithField("arTxId", arTxId).Debug(content)

	// convert content to upload format
	page, err := utils.Content2ArweavePage(content)
	if err != nil {
		return "", err
	}

	// upload page to notion
	uuid, err = o.NotionOpter.UploadPage(parentId, page)
	if err != nil {
		log.Error("upload page to notion failed! error:", err.Error())
		return "", err
	}
	log.WithField("uuid", uuid).Info("upload page to notion success")

	return uuid, nil
}
