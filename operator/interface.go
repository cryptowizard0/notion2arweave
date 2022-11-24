package operator

// INotionOperator
type INotionOperator interface {
	// Fetch
	// !MUST support recursive
	// @Pararm uuid,
	// @Return string, txId return by Arweave
	FetchPage(uuid string) (content string, err error)
	FetchDatabase(uuid string) (content string, err error)
	FetchImage(uuid string) (content string, err error)
	// Level 2 func
	// FetchChildBlocks, get subblocks
	FetchChildBlocks(parentId string) (content string, err error)

	// Upload content to notion
	// @Pararm parentId, uuid of parent page
	// @Return uuid, new uuid
	// @Return content, json format page content
	UploadPage(parentId, content string) (uuid string, err error)
	UploadDatabase(parentId, content string) (uuid string, err error)
	UploadImage(content string) (uuid string, err error)
}

// IArweaveOperator
type IArweaveOperator interface {
	// Load
	// @Return content, json format
	LoadPage(arTxId string) (content string, err error)
	LoadDatabase(arTxId string) (content string, err error)
	LoadImage(arTxId string) (content string, err error)

	// Save
	SavePage(content string) (txId string, err error)
	SaveDatabase(content string) (txId string, err error)
	SaveImage(content string) (txId string, err error)
}

// IContentAnalyzer
type IContentAnalyzer interface {
	// Covert2UploadContent Converting content from arweave to upload format
	Covert2UploadContent(savedContent string) (createdContent string, err error)
}

type IOperator interface {
	// Save2Ar fetch content from notion and upload to arweave
	// !MUST support recursive
	// @Pararm uuid,
	// @Return string, txId return by Arweave
	SavePage2Ar(uuid string) (arTxId string, err error)
	SaveDatabase2Ar(uuid string) (arTxId string, err error)
	SaveImage2Ar(uuid string) (arTxId string, err error)

	// LoadFromAr get content from arweave and upload to notion
	// @Pararm parentId, uuid of parent page
	// @Return uuid, new uuid
	// @Return content, json format page content
	LoadPageFromAr(parentId, arTxId string) (uuid string, err error)
	LoadDatabaseFromAr(parentId, arTxId string) (uuid string, err error)
	LoadImageFromAr(parentId, arTxId string) (uuid string, err error)
}
