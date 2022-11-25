package types

import "github.com/dstotijn/go-notion"

type ArweavePage struct {
	PageInfo    notion.Page                  `json:"page_info"`
	PageContent notion.BlockChildrenResponse `json:"page_content"`
}
