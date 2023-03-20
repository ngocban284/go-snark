package main

import (
	"fmt"
	"log"

	"github.com/chenzhijie/go-web3"
)

func main() {

	// change to your rpc provider
	var rpcProvider = "https://data-seed-prebsc-2-s1.binance.org:8545"
	web3, err := web3.NewWeb3(rpcProvider)
	if err != nil {
		panic(err)
	}
	blockNumber, err := web3.Eth.GetBlockNumber()
	if err != nil {
		panic(err)
	}
	fmt.Println("Current block number: ", blockNumber)

	//
	chainId, err := web3.Eth.ChainID()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current chainId: ", chainId)
}
