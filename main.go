package main

import (
	"fmt"

	"github.com/cryptowizard0/notion2arweave/operator"
	"github.com/cryptowizard0/notion2arweave/types"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Hello notion 2 arweave!")

	queryUseOPerator()
}

func queryUseOPerator() {
	opt := operator.CreateNotionOperator(types.Notion_Auth)
	txId, err := opt.FetchPage("c904d90c9abf4de68f7520786193d4c0")
	if err != nil {
		fmt.Println(err.Error())
	}

	log.WithField("txId", txId).Info("fetch_page OK!")
}
