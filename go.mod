module go/zksnark

go 1.19

require (
	github.com/ethereum/go-ethereum v1.10.26 // indirect
	github.com/iden3/go-iden3-crypto v0.0.13 // indirect
	github.com/iden3/go-rapidsnark/prover v0.0.9 // indirect
	github.com/iden3/go-rapidsnark/types v0.0.2 // indirect
	github.com/iden3/go-rapidsnark/verifier v0.0.3 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
)


// Copy the same replace dep from https://github.com/bnb-chain/go-sdk/blob/master/go.mod
replace (
    github.com/cosmos/cosmos-sdk => github.com/bnb-chain/bnc-cosmos-sdk v0.25.4-0.20221221115251-f9e69ff1b273
    github.com/tendermint/go-amino => github.com/bnb-chain/bnc-go-amino v0.14.1-binance.2
    github.com/tendermint/iavl => github.com/bnb-chain/bnc-tendermint-iavl v0.12.0-binance.4
    github.com/tendermint/tendermint => github.com/bnb-chain/bnc-tendermint v0.32.3-binance.3.0.20221109023026-379ddbab19d1
    github.com/zondax/ledger-cosmos-go => github.com/bnb-chain/ledger-cosmos-go v0.9.9-binance.3
    github.com/zondax/ledger-go => github.com/bnb-chain/ledger-go v0.9.1
    golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20190823183015-45b1026d81ae
)