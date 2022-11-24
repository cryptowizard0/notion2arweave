package operator

import (
	"fmt"

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

func (o *Operator) SavePage2Ar(uuid string) (arTxId string, err error) {
	log.WithField("uuid", uuid).Info("operator: save pate to arweave")

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

	return
}

func (o *Operator) LoadPageFromAr(parentId, arTxId string) (uuid string, err error) {
	log.WithField("arTxId", arTxId).WithField("parent", parentId).Info("operator: load page from arweave")

	// load content from arseeding
	srcContent, err := o.ArOpter.LoadPage(arTxId)
	if err != nil {
		log.Error("get content from arweave error: ", err.Error())
		return "", err
	}
	if srcContent == "" {
		log.WithField("arTxId", arTxId).Error("Can't find content within ArTx")
		return "", fmt.Errorf("can't find content within ArTx: %s", arTxId)
	}

	// upload page to notion
	uuid, err = o.NotionOpter.UploadPage(parentId, srcContent)
	if err != nil {
		log.Error("upload page to notion failed! error:", err.Error())
	}

	return
}
