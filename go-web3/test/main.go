package main

import (
	"github.com/chenzhijie/go-web3"
)

func main() {

	// change to your rpc provider
	var rpcProvider = "https://data-seed-prebsc-2-s1.binance.org:8545"
	web3, err := web3.NewWeb3(rpcProvider)
	if err != nil {
		panic(err)
	}

}
