package main

import (
	// "bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	// "testing"
	types "github.com/tendermint/tendermint/types"
	// "github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"
	// merkle "github.com/tendermint/tendermint/crypto/merkle"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

// synthetic txs to Txs
func b64ToHex(txs ...string) []string {
	var txHexs []string
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
		txHexs = append(txHexs, txHex)
	}
	// fmt.Println(txHexs)
	return txHexs
}

func txHashToBytes(txHashs []string) types.Txs {
	var txs types.Txs
	for _, txHash := range txHashs {
		txBytes, err := hex.DecodeString(txHash)
		// fmt.Println("txByte", hex.EncodeToString(txBytes))
		if err != nil {
			panic(err)
		}
		txs = append(txs, txBytes)
		// fmt.Println(txs)
	}
	return txs
}

func main() {

	txsHex := []string{
		"CtoCCtcCCiQvY29zbXdhc20ud2FzbS52MS5Nc2dFeGVjdXRlQ29udHJhY3QSrgIKK29yYWkxN3h1eW0wcGVwcmphbmVycmNqbGZwMG1oNzN6bG14dzllYXhmdWYSK29yYWkxbTZxNWs1bnIyZWg4cTByZHJmNTd3cjdwaGs3dXZscGc3bXdmdjUawwF7InByb3ZpZGVfbGlxdWlkaXR5Ijp7ImFzc2V0cyI6W3siaW5mbyI6eyJ0b2tlbiI6eyJjb250cmFjdF9hZGRyIjoib3JhaTFsdXMwZjByaHg4czAzZ2RsbHgybjZ2aGttZjA1MzZkdjU3d2ZnZSJ9fSwiYW1vdW50IjoiMjQ1NTA0MiJ9LHsiaW5mbyI6eyJuYXRpdmVfdG9rZW4iOnsiZGVub20iOiJvcmFpIn19LCJhbW91bnQiOiI1NzgwIn1dfX0qDAoEb3JhaRIENTc4MBJnClEKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiECEiE8R0b4E5z/VQhzn9pE6nKkfSQCJoMll2IIo8jKShgSBAoCCAEYugcSEgoMCgRvcmFpEgQxNjE1ENftIBpAibKp2DM/qdczdVX2Bti74nSY5E5MoI98j15PN8OzET0wEy+1dt3PIK+qO37CaHpjAo+udQYxp18TO26jSY8JRw==",
		"CtQFCpIDCiQvY29zbXdhc20ud2FzbS52MS5Nc2dFeGVjdXRlQ29udHJhY3QS6QIKK29yYWkxc2VucjBlOXpmbDRtd3B4YWhheXV2dWNhYWNtcTJxeTlkbWo0bG0SK29yYWkxMGxkZ3p1ZWQ2empwMG1rcXdzdjJtdXgzbWw1MGw5N2M3NHg4c2cajAJ7InNlbmQiOnsiYW1vdW50IjoiMTQwOTk5MDAwMDAwIiwiY29udHJhY3QiOiJvcmFpMTR3eTh4bmRobnZqbXg2emwyODY2eHF2czdmcXd2MmFyaGhycXE5IiwibXNnIjoiZXlKamIyNTJaWEowWDNKbGRtVnljMlVpT25zaVpuSnZiU0k2ZXlKdVlYUnBkbVZmZEc5clpXNGlPbnNpWkdWdWIyMGlPaUpwWW1NdlF6UTFPRUkwUTBNMFJqVTFPREV6T0RoQ09VRkRRalF3TnpjMFJrUkdRa05GUkVNM04wRTNSamREUkVaQ01URXlRalEyT1RjNU5FRkdPRFpETkVFMk9TSjlmWDE5In19CrwCCikvaWJjLmFwcGxpY2F0aW9ucy50cmFuc2Zlci52MS5Nc2dUcmFuc2ZlchKOAgoIdHJhbnNmZXISCmNoYW5uZWwtMjAaYApEaWJjL0M0NThCNENDNEY1NTgxMzg4QjlBQ0I0MDc3NEZERkJDRURDNzdBN0Y3Q0RGQjExMkI0Njk3OTRBRjg2QzRBNjkSGDE0MDk5OTAwMDAwMDAwMDAwMDAwMDAwMCIrb3JhaTFzZW5yMGU5emZsNG13cHhhaGF5dXZ1Y2FhY21xMnF5OWRtajRsbSosb3JhaWIxc2VucjBlOXpmbDRtd3B4YWhheXV2dWNhYWNtcTJxeTk2Nmtld2M4gLq3pLj+yqYXQi9vcmFpYjB4NzRDODIxMjliM0NEMTBkNGNFRjUwN2ZkMTM3ZkY4NDI5QzY1REQ3YxJnClEKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiEDqrsYkjbYPhEEEDdODVA0QZHH38HdmjZv2IuyuYhpngMSBAoCCAEYwnwSEgoMCgRvcmFpEgQ1MDAxEICJehpAbGnGdb7lQxAVPY5VQnDVkJvGvy277zI8TRvqI0jlH4xc82AArPgDPuuCfviOt/9oVCm5H91d6ynDUrNJDiX6qQ==",
	}

	txs := txHashToBytes(b64ToHex(txsHex...))

	txBzs := make([][]byte, len(txs))
	for i := 0; i < len(txs); i++ {
		txBzs[i] = txs[i].Hash()

	}
	zero := make([]byte, 1)
	zero[0] = 0x00
	one := make([]byte, 1)
	one[0] = 0x01

	leaf1 := append(zero, txBzs[0]...)
	leaf1_hash := tmhash.Sum(leaf1)

	leaf2 := append(zero, txBzs[1]...)
	leaf2_hash := tmhash.Sum(leaf2)

	root1 := append(one, leaf1_hash...)
	root2 := append(root1, leaf2_hash...)
	root_hash := tmhash.Sum(root2)

	root := txs.Hash()
	fmt.Println("root_hash", hex.EncodeToString(root_hash))
	fmt.Println("root", hex.EncodeToString(root))
	fmt.Println(txs)

	l := len(txs)
	bzs := make([][]byte, l)
	for i := 0; i < l; i++ {
		bzs[i] = txs[i].Hash()
	}

	fmt.Println(hex.EncodeToString(leaf1_hash))
	// make sure valid proof for every tx
	for i := range txs {
		tx := []byte(txs[i])
		proof := txs.Proof(i)
		fmt.Println("----------------------------------")
		fmt.Println(tx)
		fmt.Println(hex.EncodeToString(proof.Proof.Aunts[0]))
		fmt.Println("index", proof.Proof.Index)
		fmt.Println("Total", proof.Proof.Total)
		fmt.Println("RootHash", proof.RootHash)
		fmt.Println("leaf", hex.EncodeToString(proof.Leaf()))
		fmt.Println("validateRoot", proof.Validate(root))
		fmt.Println("notnil", proof.Validate([]byte("foobar")))
		// assert.EqualValues(t, i, proof.Proof.Index)
		// assert.EqualValues(t, len(txs), proof.Proof.Total)
		// assert.EqualValues(t, root, proof.RootHash)
		// assert.EqualValues(t, tx, proof.Data)
		// assert.EqualValues(t, txs[i].Hash(), proof.Leaf())
		// assert.Nil(t, proof.Validate(root))
		// assert.NotNil(t, proof.Validate([]byte("foobar")))

		// read-write must also work
		var (
			p2  types.TxProof
			pb2 tmproto.TxProof
		)
		pbProof := proof.ToProto()
		bin, _ := pbProof.Marshal()

		err := pb2.Unmarshal(bin)
		fmt.Println(err)
		p2, _ = types.TxProofFromProto(pb2)
		fmt.Println(p2)
	}

}
