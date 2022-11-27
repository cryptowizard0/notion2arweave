package utils

import (
	"encoding/json"
	"fmt"

	"github.com/cryptowizard0/notion2arweave/types"
	"github.com/dstotijn/go-notion"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

// TODO: Can't support sub page, image, database now.
//
// Content2ArweavePage converting content from arweave to ArweavePage struct
func Content2ArweavePage(srcContent string) (*types.ArweavePage, error) {
	var page types.ArweavePage
	err := json.Unmarshal([]byte(srcContent), &page)
	if err != nil {
		return nil, err
	}
	log.Info("Blocks count: ", len(page.PageContent.Results))
	var tmpBlocks []notion.Block

	for _, block := range page.PageContent.Results {
		switch block.(type) {
		// supported block types
		case *notion.ParagraphBlock,
			*notion.Heading1Block,
			*notion.Heading2Block,
			*notion.Heading3Block,
			*notion.BulletedListItemBlock,
			*notion.NumberedListItemBlock,
			*notion.ToDoBlock,
			*notion.ToggleBlock,
			*notion.CalloutBlock,
			*notion.DividerBlock,
			*notion.QuoteBlock:
			tmpBlocks = append(tmpBlocks, block)

		// case notion.ChildPageBlock:
		// case notion.ChildDatabaseBlock:
		// case notion.CodeBlock:
		// case notion.EmbedBlock:
		// case notion.ImageBlock:
		// case notion.AudioBlock:
		// case notion.VideoBlock:
		// case notion.FileBlock:
		// case notion.PDFBlock:
		// case notion.BookmarkBlock:
		// case notion.EquationBlock:
		// case notion.TableOfContentsBlock:
		// case notion.BreadcrumbBlock:
		// case notion.ColumnListBlock:
		// case notion.ColumnBlock:
		// case notion.TableBlock:
		// case notion.TableRowBlock:
		// case notion.LinkPreviewBlock:
		// case notion.LinkToPageBlock:
		// case notion.SyncedBlock:
		// case notion.TemplateBlock:
		default:
			// do nothing
		}
	}
	page.PageContent.Results = tmpBlocks

	return &page, nil
}

// MergeChildBlocks
// content format is notion.BlockChildrenResponse
func MergeChildBlocks(content1, content2 string) (merged string, err error) {
	value1 := gjson.Get(content1, "results")
	value2 := gjson.Get(content2, "results")

	array1 := value1.Array()
	array2 := value2.Array()
	array1 = append(array1, array2...)
	strArr := "["
	for i, block := range array1 {
		strArr += block.String()
		if i < len(array1)-1 {
			strArr += ","
		}
	}
	strArr += "]"

	merged = fmt.Sprintf("{\"object\":\"%s\",\"results\":%s,\"type\":\"%s\",\"block\":\"%s\"}",
		gjson.Get(content1, "object").String(),
		strArr,
		gjson.Get(content1, "type").String(),
		gjson.Get(content1, "obblockject").String())

	log.Debug("merged: ", merged)

	return merged, nil
}
