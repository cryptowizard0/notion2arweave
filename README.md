# notion2arweave
Arweave2notion can store the content of notion to arweave and  restore the pages to notion.

## Compile and run

```sh
$ go tidy

$ go build

$ ./notion2arweave
```

## Config
```toml
appname = "notion2arweave"
version = "0.1.0"

[arweave]
    # Your Metamask private key, used to pay for uploading arweave, using everpay
	pk = "xxxxxxxxx" 
	everpay_url = "https://api.everpay.io"
	arseed_url = "https://arseed.web3infra.dev"

[notion]
	api_auth = "secret_xxxxx" # your notion secret key
	base_url = "https://api.notion.com"
	version = "2022-06-28"

[service]
	port = "2333" # service port
```

### 👉 Key references
- arweave: https://www.arweave.org/
- everpay: https://everpay.io/
- arseeding: https://web3infra.dev/
- go-notion fork: https://github.com/cryptowizard0/go-notion 
# Restful API
**Save notion page to arweave**
```
GET: /v1/page/save/:uuid

```
- **uuid:** uuid of the notion page.

**Load a page from arweave and add it to notion**
```
GET: /v1/page/load/:parent/:artxid
```
- **parent:** uuid of the notion page to load new page.
- **artxid:** transaction id of arweave for storing page content.


# Supported notion block types
- Paragraph
- Heading1
- Heading2
- Heading3
- NumberedListItem
- BulletedListItem
- ToDo
- Toggle
- Callout
- Divider
- Quote
- Video
- Image

## TODO
- [ ] Download sub page
- [ ] Support more block types on notion
- [ ] Store TableView content 
- [ ] Front-end website
- [ ] User wallet signature and payments