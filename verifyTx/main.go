package main

// decode txs base64 to txHash in tendermint

import (
	"encoding/base64"
	"fmt"

	"github.com/tendermint/tendermint/crypto/tmhash"
)

func main() {
	// txs base64
	txs := []string{
		"CocCCoICCiQvY29zbXdhc20ud2FzbS52MS5Nc2dFeGVjdXRlQ29udHJhY3QS2QEKK29yYWkxN3ZzZ3lmZmRnazUyNzlxeXV6dHZ6amd0MnE1a2x6aGpmNzIydzkSK29yYWkxOXA0M3kwdHFucjVxbGhmd254ZnQydTV1bnBoNXluNjB5N3R1dnUafXsid2l0aGRyYXciOnsiYXNzZXRfaW5mbyI6eyJuYXRpdmVfdG9rZW4iOnsiZGVub20iOiJpYmMvQTJFMkVFQzkwNTdBNEExQzJDMEE2QTRDNzhCMDIzOTExOERGNUYyNzg4MzBGNTBCNEE2QkREN0E2NjUwNkI3OCJ9fX19EgASZgpRCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohA7tDviTE5kBZ7gxAHfh6xzZd9iyJs3YK+OUz4+L3oozlEgQKAggBGL4BEhEKCwoEb3JhaRIDNTAwEJbBEBpAV4sY7oyq2wgMNNHZfbrJwoAS71ectOC9BIPHZrf+kVQMHHD7ri8Pjz4ICQqiF+FWv8SdjnaQUJ2mmAA1o56QXg==",
		"CoICCv8BCiQvY29zbXdhc20ud2FzbS52MS5Nc2dFeGVjdXRlQ29udHJhY3QS1gEKK29yYWkxa21qcmxkZ2ozd2FrZjRxbWV1ZHJjZWQwbTl5N3FoMG01YXNmMngSK29yYWkxbmQ0cjA1M2Uza2dlZGdsZDJ5bWVuOGw5eXJ3OHhwanlhYWw3ajUaensiaW5jcmVhc2VfYWxsb3dhbmNlIjp7ImFtb3VudCI6Ijk5OTk5OTk5OTk5OTk5OTk5OTk5OTk5OTk5OTk5OSIsInNwZW5kZXIiOiJvcmFpMXlubWQyY2VtcnloY3d0anEzYWRoY3dheXJtODlsMmNyNHR3czR2In19EmUKUApGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQIS0r5ZO0japSdfqZC7mT4u85eoP1xyHn3n/x2lFnWHlxIECgIIARgCEhEKCwoEb3JhaRIDNzM2EO79CBpAJRZ6ytb+q1iTUkasCMly7iF+twZYIHEbUtUMpJfabKBkq5GKPPuaezY5k48ivQknLdyg5lu6ojv21NXamaPSVA==",
	}

	for _, tx := range txs {
		// decode base64
		b, err := base64.StdEncoding.DecodeString(tx)
		if err != nil {
			panic(err)
		}
		// txBytes
		txBytes := tmhash.Sum(b)

		// txHash
		txHex := fmt.Sprintf("%x", txBytes)
		fmt.Println(txHex)
	}

}
