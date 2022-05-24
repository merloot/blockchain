package main

import (
	"merloot/blockchain/internal/blockchain"
	"merloot/blockchain/internal/cli"
)

func main() {

	bc := blockchain.NewBlockchain()

	defer bc.DB.Close()

	cli := cli.CLI{BC: bc}
	cli.Run()
}
