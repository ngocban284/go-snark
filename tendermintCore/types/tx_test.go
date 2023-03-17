package types

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto/tmhash"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	ctest "github.com/tendermint/tendermint/libs/test"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func makeTxs(cnt, size int) Txs {
	txs := make(Txs, cnt)
	for i := 0; i < cnt; i++ {
		txs[i] = tmrand.Bytes(size)
	}
	return txs
}

func randInt(low, high int) int {
	off := tmrand.Int() % (high - low)
	return low + off
}

func TestTxIndex(t *testing.T) {
	for i := 0; i < 20; i++ {
		txs := makeTxs(15, 60)
		for j := 0; j < len(txs); j++ {
			tx := txs[j]
			idx := txs.Index(tx)
			assert.Equal(t, j, idx)
		}
		assert.Equal(t, -1, txs.Index(nil))
		assert.Equal(t, -1, txs.Index(Tx("foodnwkf")))
	}
}

func TestTxIndexByHash(t *testing.T) {
	for i := 0; i < 20; i++ {
		txs := makeTxs(15, 60)
		for j := 0; j < len(txs); j++ {
			tx := txs[j]
			idx := txs.IndexByHash(tx.Hash())
			assert.Equal(t, j, idx)
		}
		assert.Equal(t, -1, txs.IndexByHash(nil))
		assert.Equal(t, -1, txs.IndexByHash(Tx("foodnwkf").Hash()))
	}
}

// base64 => dex => Tx
func b64(s string) Tx {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

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
	return txHexs
}

func txHashToBytes(txHashs []string) Txs {
	var txs Txs
	for _, txHash := range txHashs {
		txBytes, err := hex.DecodeString(txHash)
		if err != nil {
			panic(err)
		}
		txs = append(txs, txBytes)
	}
	return txs
}

func TestValidTxProof(t *testing.T) {

	txsHex := []string{
		"CtoCCtcCCiQvY29zbXdhc20ud2FzbS52MS5Nc2dFeGVjdXRlQ29udHJhY3QSrgIKK29yYWkxN3h1eW0wcGVwcmphbmVycmNqbGZwMG1oNzN6bG14dzllYXhmdWYSK29yYWkxbTZxNWs1bnIyZWg4cTByZHJmNTd3cjdwaGs3dXZscGc3bXdmdjUawwF7InByb3ZpZGVfbGlxdWlkaXR5Ijp7ImFzc2V0cyI6W3siaW5mbyI6eyJ0b2tlbiI6eyJjb250cmFjdF9hZGRyIjoib3JhaTFsdXMwZjByaHg4czAzZ2RsbHgybjZ2aGttZjA1MzZkdjU3d2ZnZSJ9fSwiYW1vdW50IjoiMjQ1NTA0MiJ9LHsiaW5mbyI6eyJuYXRpdmVfdG9rZW4iOnsiZGVub20iOiJvcmFpIn19LCJhbW91bnQiOiI1NzgwIn1dfX0qDAoEb3JhaRIENTc4MBJnClEKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiECEiE8R0b4E5z/VQhzn9pE6nKkfSQCJoMll2IIo8jKShgSBAoCCAEYugcSEgoMCgRvcmFpEgQxNjE1ENftIBpAibKp2DM/qdczdVX2Bti74nSY5E5MoI98j15PN8OzET0wEy+1dt3PIK+qO37CaHpjAo+udQYxp18TO26jSY8JRw==",
		"CtQFCpIDCiQvY29zbXdhc20ud2FzbS52MS5Nc2dFeGVjdXRlQ29udHJhY3QS6QIKK29yYWkxc2VucjBlOXpmbDRtd3B4YWhheXV2dWNhYWNtcTJxeTlkbWo0bG0SK29yYWkxMGxkZ3p1ZWQ2empwMG1rcXdzdjJtdXgzbWw1MGw5N2M3NHg4c2cajAJ7InNlbmQiOnsiYW1vdW50IjoiMTQwOTk5MDAwMDAwIiwiY29udHJhY3QiOiJvcmFpMTR3eTh4bmRobnZqbXg2emwyODY2eHF2czdmcXd2MmFyaGhycXE5IiwibXNnIjoiZXlKamIyNTJaWEowWDNKbGRtVnljMlVpT25zaVpuSnZiU0k2ZXlKdVlYUnBkbVZmZEc5clpXNGlPbnNpWkdWdWIyMGlPaUpwWW1NdlF6UTFPRUkwUTBNMFJqVTFPREV6T0RoQ09VRkRRalF3TnpjMFJrUkdRa05GUkVNM04wRTNSamREUkVaQ01URXlRalEyT1RjNU5FRkdPRFpETkVFMk9TSjlmWDE5In19CrwCCikvaWJjLmFwcGxpY2F0aW9ucy50cmFuc2Zlci52MS5Nc2dUcmFuc2ZlchKOAgoIdHJhbnNmZXISCmNoYW5uZWwtMjAaYApEaWJjL0M0NThCNENDNEY1NTgxMzg4QjlBQ0I0MDc3NEZERkJDRURDNzdBN0Y3Q0RGQjExMkI0Njk3OTRBRjg2QzRBNjkSGDE0MDk5OTAwMDAwMDAwMDAwMDAwMDAwMCIrb3JhaTFzZW5yMGU5emZsNG13cHhhaGF5dXZ1Y2FhY21xMnF5OWRtajRsbSosb3JhaWIxc2VucjBlOXpmbDRtd3B4YWhheXV2dWNhYWNtcTJxeTk2Nmtld2M4gLq3pLj+yqYXQi9vcmFpYjB4NzRDODIxMjliM0NEMTBkNGNFRjUwN2ZkMTM3ZkY4NDI5QzY1REQ3YxJnClEKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiEDqrsYkjbYPhEEEDdODVA0QZHH38HdmjZv2IuyuYhpngMSBAoCCAEYwnwSEgoMCgRvcmFpEgQ1MDAxEICJehpAbGnGdb7lQxAVPY5VQnDVkJvGvy277zI8TRvqI0jlH4xc82AArPgDPuuCfviOt/9oVCm5H91d6ynDUrNJDiX6qQ==",
	}

	cases := []struct {
		txs Txs
	}{
		{txHashToBytes(b64ToHex(txsHex...))},
	}

	for h, tc := range cases {
		txs := tc.txs
		root := txs.Hash()
		// make sure valid proof for every tx
		for i := range txs {
			tx := []byte(txs[i])
			proof := txs.Proof(i)
			assert.EqualValues(t, i, proof.Proof.Index, "%d: %d", h, i)
			assert.EqualValues(t, len(txs), proof.Proof.Total, "%d: %d", h, i)
			assert.EqualValues(t, root, proof.RootHash, "%d: %d", h, i)
			assert.EqualValues(t, tx, proof.Data, "%d: %d", h, i)
			assert.EqualValues(t, txs[i].Hash(), proof.Leaf(), "%d: %d", h, i)
			assert.Nil(t, proof.Validate(root), "%d: %d", h, i)
			assert.NotNil(t, proof.Validate([]byte("foobar")), "%d: %d", h, i)

			// read-write must also work
			var (
				p2  TxProof
				pb2 tmproto.TxProof
			)
			pbProof := proof.ToProto()
			bin, err := pbProof.Marshal()
			require.NoError(t, err)

			err = pb2.Unmarshal(bin)
			require.NoError(t, err)

			p2, err = TxProofFromProto(pb2)
			if assert.Nil(t, err, "%d: %d: %+v", h, i, err) {
				assert.Nil(t, p2.Validate(root), "%d: %d", h, i)
			}
		}
	}
}

func TestTxProofUnchangable(t *testing.T) {
	// run the other test a bunch...
	for i := 0; i < 40; i++ {
		testTxProofUnchangable(t)
	}
}

func testTxProofUnchangable(t *testing.T) {
	// make some proof
	txs := makeTxs(randInt(2, 100), randInt(16, 128))
	root := txs.Hash()
	i := randInt(0, len(txs)-1)
	proof := txs.Proof(i)

	// make sure it is valid to start with
	assert.Nil(t, proof.Validate(root))
	pbProof := proof.ToProto()
	bin, err := pbProof.Marshal()
	require.NoError(t, err)

	// try mutating the data and make sure nothing breaks
	for j := 0; j < 500; j++ {
		bad := ctest.MutateByteSlice(bin)
		if !bytes.Equal(bad, bin) {
			assertBadProof(t, root, bad, proof)
		}
	}
}

// This makes sure that the proof doesn't deserialize into something valid.
func assertBadProof(t *testing.T, root []byte, bad []byte, good TxProof) {

	var (
		proof   TxProof
		pbProof tmproto.TxProof
	)
	err := pbProof.Unmarshal(bad)
	if err == nil {
		proof, err = TxProofFromProto(pbProof)
		if err == nil {
			err = proof.Validate(root)
			if err == nil {
				// XXX Fix simple merkle proofs so the following is *not* OK.
				// This can happen if we have a slightly different total (where the
				// path ends up the same). If it is something else, we have a real
				// problem.
				assert.NotEqual(t, proof.Proof.Total, good.Proof.Total, "bad: %#v\ngood: %#v", proof, good)
			}
		}
	}
}
