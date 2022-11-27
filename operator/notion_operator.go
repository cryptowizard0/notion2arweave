package operator

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dstotijn/go-notion"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/cryptowizard0/notion2arweave/types"
	"github.com/cryptowizard0/notion2arweave/utils"
	"github.com/go-resty/resty/v2"
)

// NotionOperator inherite INotionOperator
type NotionOperator struct {
	authToken    string
	client       *resty.Client
	notionClient *notion.Client
}

// Implementation of Notion interaction
// Inheritance INotionOperator
//
// CreateNotionOperator
func CreateNotionOperator(auth string) *NotionOperator {
	client := resty.New()
	client.SetHeader("Accept", "application/json").
		SetHeader("Notion-Version", viper.GetString("notion.version")).
		SetAuthToken(auth).
		SetBaseURL(viper.GetString("notion.base_url"))

	return &NotionOperator{
		authToken:    auth,
		client:       client,
		notionClient: notion.NewClient(auth),
	}
}

// Fetch page from notion, upload to arweave4
// @Pararm uuid, page uuid
// @Return txId, txId return by Arweave
func (n *NotionOperator) FetchPage(uuid string) (content string, err error) {
	log.WithField("uuid", uuid).Info("notion operator: fetch page")

	// 1. get page info
	strPageInfo, err := n.fetchPageInfo(uuid)
	if err != nil {
		log.Error("fetch page info error:", err.Error())
		return "", err
	}

	// 2. get child blocks
	strPageContent, err := n.fetchPageContent(uuid, "")
	if err != nil {
		log.Error("fetch page info error:", err.Error())
		return "", err
	}

	// 3. make full content
	content = fmt.Sprintf(`{"page_info":%s,"page_content":%s}`, strPageInfo, strPageContent)
	// log.WithField("page id", uuid).Debug(content)

	return
}

// UploadPage uploading page content to notion
// @Pararm parentId, parent page uuid, where new page to be loaded
// @Pararm content, page content ,need convert to upload format
// @Return uuid, uuid of new page
// @Return content, new page content
func (n *NotionOperator) UploadPage(parentId string, page *types.ArweavePage) (uuid string, err error) {
	log.WithField("parent", parentId).Info("notion operator: upload page")

	pageProp, ok := page.PageInfo.Properties.(notion.PageProperties)
	if !ok {
		return "", fmt.Errorf("convert page preperites error")
	}

	children := page.PageContent.Results

	newPageParams := notion.CreatePageParams{
		ParentType: notion.ParentTypePage,
		ParentID:   parentId,
		Title:      pageProp.Title.Title,
		Children:   nil,
		Icon:       page.PageInfo.Icon,
		Cover:      page.PageInfo.Cover,
	}

	newPage, err := n.notionClient.CreatePage(context.Background(), newPageParams)
	if err != nil {
		return "", err
	}

	// upload content blocks, Max 100 per request
	for i := 0; (i * 100) < len(children); i++ {
		starindex := i * 100
		endindex := i*100 + 100
		if endindex > len(children) {
			endindex = len(children)
		}
		_, err = n.notionClient.AppendBlockChildren(context.Background(), newPage.ID, children[starindex:endindex])
		if err != nil {
			return "", err
		}
	}

	return newPage.ID, nil
}

// ========================================================================

// fetchPageInfo
func (n *NotionOperator) fetchPageInfo(uuid string) (content string, err error) {
	url := fmt.Sprintf("/v1/pages/%s", uuid)
	resp, err := n.client.R().Get(url)
	if err != nil {
		log.Error("get request error: ", err.Error())
		return "", err
	}
	if resp.StatusCode() != http.StatusOK {
		utils.LogResp_Error(resp)
		return "", fmt.Errorf(resp.String())
	}

	return string(resp.Body()), nil
}

func (n *NotionOperator) fetchPageInfoByNotionSdk(uuid string) (content string, err error) {
	page, err := n.notionClient.FindPageByID(context.Background(), uuid)
	if err != nil {
		log.Error("get page error: ", err.Error())
		return "", err
	}

	jsonPage, err := json.Marshal(page)
	if err != nil {
		return "", err
	}

	return string(jsonPage), nil
}

// fetchPageContent
func (n *NotionOperator) fetchPageContent(uuid, startCursor string) (content string, err error) {
	var url string
	if startCursor == "" {
		url = fmt.Sprintf("/v1/blocks/%s/children", uuid)
	} else {
		url = fmt.Sprintf("/v1/blocks/%s/children?start_cursor=%s", uuid, startCursor)
	}

	resp, err := n.client.R().Get(url)
	if err != nil {
		log.Error("get request error:", err.Error())
		return "", err
	}
	if resp.StatusCode() != http.StatusOK {
		utils.LogResp_Error(resp)
		return "", fmt.Errorf(resp.String())
	}

	type hasMore struct {
		HasMore    bool    `json:"has_more"`
		NextCursor *string `json:"next_cursor"`
	}
	var morePage hasMore
	err = json.Unmarshal(resp.Body(), &morePage)
	if err != nil {
		return "", err
	}
	fullContent := string(resp.Body())

	if morePage.HasMore {
		log.WithField("id", morePage.NextCursor).Info("more page")
		moreContent, err := n.fetchPageContent(uuid, *morePage.NextCursor)
		if err != nil {
			return "", err
		}
		fullContent, err = utils.MergeChildBlocks(string(resp.Body()), moreContent)
		if err != nil {
			return "", err
		}
	}

	return fullContent, nil
}
