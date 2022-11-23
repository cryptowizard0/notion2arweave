package operator

// INotionOperator
type INotionOperator interface {
	// Fetch
	// !MUST support recursive
	// @Pararm uuid,
	// @Return string, txId return by Arweave
	FetchPage(uuid string) (txId string, err error)
	FetchDatabase(uuid string) (txId string, err error)
	FetchImage(uuid string) (txId string, err error)
	// Level 2 func
	// FetchChildBlocks, get subblocks
	FetchChildBlocks(parentId string) (content string, err error)

	// Upload
	// @Pararm parentId, uuid of parent page
	// @Return uuid, new uuid
	// @Return content, json format page content
	UploadPage(parentId, arTxId string) (uuid, content string, err error)
	UploadDatabase(parentId, arTxId string) (uuid, content string, err error)
	UploadImage(arTxid string) (uuid, content string, err error)
}

// IContentAnalyzer
type IContentAnalyzer interface {
	// Saved2Created Converting stored content on Notion to created content, Json format
	Saved2Created(savedContent string) (createdContent string, err error)
}

// IArweaveOperator
type IArweaveOperator interface {
	// Load
	// @Return content, json format
	LoadPage() (content string, err error)
	LoadDatabase() (content string, err error)
	LoadImage() (content string, err error)

	// Save
	SavePage(content string) (txId string, err error)
	SaveDatabase(content string) (txId string, err error)
	SaveImage(content string) (txId string, err error)
}
