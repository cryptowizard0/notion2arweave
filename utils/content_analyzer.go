package utils

import (
	"encoding/json"

	"github.com/cryptowizard0/notion2arweave/types"
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

	return &page, nil
}
