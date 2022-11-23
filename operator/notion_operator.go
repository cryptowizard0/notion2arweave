package operator

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/cryptowizard0/notion2arweave/utils"
	"github.com/go-resty/resty/v2"
)

// NotionOperator inherite INotionOperator
type NotionOperator struct {
	authToken string
	client    *resty.Client
}

// Implementation of Notion interaction
// Inheritance INotionOperator
//
// CreateNotionOperator
func CreateNotionOperator(auth string) *NotionOperator {
	client := resty.New()
	client.SetHeader("Accept", "application/json").
		SetHeader("Notion-Version", "2022-06-28").
		SetAuthToken(auth).
		SetBaseURL(viper.GetString("notion.base_url"))

	return &NotionOperator{
		authToken: auth,
		client:    client,
	}
}

// Fetch page from notion, upload to arweave4
// @Pararm uuid, page uuid
// @Return txId, txId return by Arweave
func (n *NotionOperator) FetchPage(uuid string) (txId string, err error) {
	log.WithField("uuid", uuid).Info("notion operator: fetch page")

	// 1. get page info
	strPageInfo, err := n.fetchPageInfo(uuid)
	if err != nil {
		log.Error("fetch page info error:", err.Error())
		return "", err
	}

	// 2. get child blocks
	strPageContent, err := n.fetchPageContent(uuid)
	if err != nil {
		log.Error("fetch page info error:", err.Error())
		return "", err
	}

	// 3. make full content
	fullContent := fmt.Sprintf(`{"page_info":%s,"page_content":%s}`, strPageInfo, strPageContent)
	log.Debug(fullContent)

	// 4. upload arweave
	arOpt, err := CreateArweaveOperator(viper.GetString("arweave.pk"), "USDC")
	if err != nil {
		log.Error("create arweave operator error:", err.Error())
		return "", err
	}
	txId, err = arOpt.SavePage(fullContent)
	if err != nil {
		return "", err
	}

	return
}

// UploadPage uploading page content get from arweave to notion
// @Pararm parentId, parent page uuid, where new page to be loaded
// @Pararm arTxId, txId return by Arweave
// @Return uuid, uuid of new page
// @Return content, new page content
func (n *NotionOperator) UploadPage(parentId, arTxId string) (uuid, content string, err error) {
	log.WithField("arTxId", arTxId).WithField("parent", parentId).Info("notion operator: upload page")

	// 1. load content from ar
	arOpt, err := CreateArweaveOperator(viper.GetString("arweave.pk"), "USDC")
	if err != nil {
		log.Error("create arweave operator error:", err.Error())
		return "", "", err
	}
	srcContent, err := arOpt.LoadPage(arTxId)
	if err != nil {
		log.Error("get content from arweave error: ", err.Error())
		return "", "", err
	}
	if srcContent == "" {
		log.WithField("arTxId", arTxId).Error("Can't find content within ArTx")
		return "", "", fmt.Errorf("can't find content within ArTx: %s", arTxId)
	}

	// 2. convert content
	// 3. upload ar

	return "", srcContent, nil
}

func (n *NotionOperator) fetchPageInfo(uuid string) (content string, err error) {
	url := fmt.Sprintf("/v1/pages/%s", uuid)
	resp, err := n.client.R().Get(url)
	if err != nil {
		log.Error("get request error:", err.Error())
		return "", err
	}
	if resp.StatusCode() != http.StatusOK {
		utils.LogResp_Error(resp)
		return "", fmt.Errorf(resp.String())
	}

	return string(resp.Body()), nil
}

func (n *NotionOperator) fetchPageContent(uuid string) (content string, err error) {
	url := fmt.Sprintf("/v1/blocks/%s/children", uuid)
	resp, err := n.client.R().Get(url)
	if err != nil {
		log.Error("get request error:", err.Error())
		return "", err
	}
	if resp.StatusCode() != http.StatusOK {
		utils.LogResp_Error(resp)
		return "", fmt.Errorf(resp.String())
	}

	return string(resp.Body()), nil
}
