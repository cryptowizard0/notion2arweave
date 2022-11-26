package utils

import (
	"encoding/json"

	"github.com/cryptowizard0/notion2arweave/types"
	"github.com/dstotijn/go-notion"
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
		// case notion.DividerBlock:
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
