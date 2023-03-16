package main_test

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/light"
	"github.com/tendermint/tendermint/proto/tendermint/version"
	"github.com/tendermint/tendermint/types"
)

func First[T, U any](val T, _ U) T {
	return val
}

func TestAbc(t *testing.T) {
	const (
		chainID    = "Oraichain"
		lastHeight = 100
	)

	var (
		header = types.SignedHeader{
			Header: &types.Header{
				Version: version.Consensus{
					Block: 11,
				},
				ChainID: chainID,
				Height:  lastHeight,
				Time:    First(time.Parse(time.RFC3339, "2021-02-24T04:24:31Z")),
				LastBlockID: types.BlockID{
					Hash: First(hex.DecodeString("5A62978F71F09351A417D47E8BCF92955527AFF62D3BF351EA9E4A820A7BAE48")),
					PartSetHeader: types.PartSetHeader{
						Total: 1,
						Hash:  First(hex.DecodeString("25912DE4DABE4CFDDA28630DC30BEEA8B0FAF300CC7B8B752A808871953293D7")),
					},
				},
				LastCommitHash:     First(hex.DecodeString("9BAA37C9ABE1D05AEAD485F044C41945264F80EB7D8C9F1F68D5904644DE751B")),
				DataHash:           First(hex.DecodeString("E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855")),
				ValidatorsHash:     First(hex.DecodeString("9BDB7697219072BB7FE38DDAF09A6AEFA5BBC93A88103889A675AF70C54B7689")),
				NextValidatorsHash: First(hex.DecodeString("9BDB7697219072BB7FE38DDAF09A6AEFA5BBC93A88103889A675AF70C54B7689")),
				ConsensusHash:      First(hex.DecodeString("048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F")),
				AppHash:            First(hex.DecodeString("A965E16DB87F80A0D183EC62AE0C196B6555AFA149093CC57C54B30E6E98028E")),
				LastResultsHash:    First(hex.DecodeString("E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855")),
				EvidenceHash:       First(hex.DecodeString("E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855")),
				ProposerAddress:    First(hex.DecodeString("AD1A1DB19D951A2DBDF710D73711872D85E7D50D")),
			},
			Commit: &types.Commit{
				Height: lastHeight,
				Round:  0,
				BlockID: types.BlockID{
					Hash: First(hex.DecodeString("9858874C995832A8B61FA216241289FA4B5C1672F060E58147D78DE9AB3169EA")),
					PartSetHeader: types.PartSetHeader{
						Total: 1,
						Hash:  First(hex.DecodeString("703FC54DB43BA857AE1EE662B1FD588ACDFF5B2F4739303E3327A9E3416D17AA")),
					},
				},
				Signatures: []types.CommitSig{
					{
						BlockIDFlag:      2,
						ValidatorAddress: First(hex.DecodeString("AD1A1DB19D951A2DBDF710D73711872D85E7D50D")),
						Timestamp:        First(time.Parse(time.RFC3339, "2021-02-24T04:24:36Z")),
						Signature:        First(base64.StdEncoding.DecodeString("i8p0orWHgJmlwbToSa+3jdcFHBy0z/1YIi0DpUyE93bLsMQAhd2fjoIXCUEYilM9/QdoVjpjjkMc0U1FmuwdDQ==")),
					},
					{
						BlockIDFlag:      3,
						ValidatorAddress: First(hex.DecodeString("1D33EA191E17141E9B4EE5407E541FA6CE952E26")),
						Timestamp:        First(time.Parse(time.RFC3339, "2021-02-24T04:24:36Z")),
						Signature:        First(base64.StdEncoding.DecodeString("d4LStsxKku7nbTCVySU/66noQYzfP+Q+fZYUnUlRfWcQc/OpF2RmTEkasg5ysVGHY7El7q/V7m6Dyhe0RXNqBw==")),
					},
					{
						BlockIDFlag:      3,
						ValidatorAddress: First(hex.DecodeString("1E4689A0633BA2F472115C36869A1D1B6D044A56")),
						Timestamp:        First(time.Parse(time.RFC3339, "2021-02-24T04:24:36Z")),
						Signature:        First(base64.StdEncoding.DecodeString("6M9KLXzoEUrMLfD1wVnCFPVal4FaGgv/ULfK685Ibx7SQm/gwOh1YS7lf+mLq8dlBdwduw+YwVAaqX3EaxvmBA==")),
					},
					{
						BlockIDFlag:      3,
						ValidatorAddress: First(hex.DecodeString("76A56EEF3236312DF28486A22E1E655ADDFC1985")),
						Timestamp:        First(time.Parse(time.RFC3339, "2021-02-24T04:24:36Z")),
						Signature:        First(base64.StdEncoding.DecodeString("r/dq31ooniMziO3KYFk3S7OuK3Q4Jvlb5X5vq8xK5j3Msv6JUFvJx95u6YDzy7NYXHDVcN32+tvZK8DpD9oEAg==")),
					},
					{
						BlockIDFlag:      3,
						ValidatorAddress: First(hex.DecodeString("8C4BF7429E8BD82587D733ED6C0D3798130F21D0")),
						Timestamp:        First(time.Parse(time.RFC3339, "2021-02-24T04:24:36Z")),
						Signature:        First(base64.StdEncoding.DecodeString("ohgh4o5xz7do+BuDwdT2WfnhFVowj9zen25xnpk9Z311GiuUwT9r2tZ431Rnccm7+WTtvhuMK8PdTDeucUO3Dg==")),
					},
					{
						BlockIDFlag:      3,
						ValidatorAddress: First(hex.DecodeString("A4828E69BDE108561732CAA60C4CD12529A7694D")),
						Timestamp:        First(time.Parse(time.RFC3339, "2021-02-24T04:24:36Z")),
						Signature:        First(base64.StdEncoding.DecodeString("B1ti2RIr1SEDzl0HQQDv4+K1wZBEXOBfq5NZ14BA5NmY6gjhiYHcGTjnHV94whqyy5uvH5epMlWolbCYOkPiAA==")),
					},
					{
						BlockIDFlag:      3,
						ValidatorAddress: First(hex.DecodeString("BD5F59329BB61B62A7A852BEF1E57C99438D0C86")),
						Timestamp:        First(time.Parse(time.RFC3339, "2021-02-24T04:24:36Z")),
						Signature:        First(base64.StdEncoding.DecodeString("B+e7sxLhqbDDc9igSZ0h7gI1CQ2W1EY3JXx0QsE1j1fyWuMyTZW9fZDbGXp1Oy2uCalsMbaZjdBpt+HgC6A3Cg==")),
					},
					{
						BlockIDFlag:      3,
						ValidatorAddress: First(hex.DecodeString("E77AD30C31BB8F1AF3C7E00FEF07C77F4A5021BB")),
						Timestamp:        First(time.Parse(time.RFC3339, "2021-02-24T04:24:36Z")),
						Signature:        First(base64.StdEncoding.DecodeString("ALCJLmkHf22NVGxAaRQWFFK1s9UAjCmwm85J7xlt85a71bTlnzsbqv3c7/j36URO/NkA8oTBzraKVPotxRj8Dg==")),
					},
				},
			},
		}

		newHeader = types.SignedHeader{
			Header: &types.Header{
				Version: version.Consensus{
					Block: 11,
				},
				ChainID: chainID,
				Height:  101,
				Time:    First(time.Parse(time.RFC3339, "2021-02-24T04:24:36Z")),
				LastBlockID: types.BlockID{
					Hash: First(hex.DecodeString("9858874C995832A8B61FA216241289FA4B5C1672F060E58147D78DE9AB3169EA")),
					PartSetHeader: types.PartSetHeader{
						Total: 1,
						Hash:  First(hex.DecodeString("703FC54DB43BA857AE1EE662B1FD588ACDFF5B2F4739303E3327A9E3416D17AA")),
					},
				},
				LastCommitHash:     First(hex.DecodeString("697EA21D0AA7B017441B99EBFDC3B31C922F658AA852732B037C923248560B9D")),
				DataHash:           First(hex.DecodeString("E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855")),
				ValidatorsHash:     First(hex.DecodeString("9BDB7697219072BB7FE38DDAF09A6AEFA5BBC93A88103889A675AF70C54B7689")),
				NextValidatorsHash: First(hex.DecodeString("9BDB7697219072BB7FE38DDAF09A6AEFA5BBC93A88103889A675AF70C54B7689")),
				ConsensusHash:      First(hex.DecodeString("048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F")),
				AppHash:            First(hex.DecodeString("08DA86F6DF6361D020568AC2EFD1B7C9C04103023D8738B331820EA870184730")),
				LastResultsHash:    First(hex.DecodeString("E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855")),
				EvidenceHash:       First(hex.DecodeString("E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855")),
				ProposerAddress:    First(hex.DecodeString("AD1A1DB19D951A2DBDF710D73711872D85E7D50D")),
			},
			Commit: &types.Commit{
				Height: 101,
				Round:  0,
				BlockID: types.BlockID{
					Hash: First(hex.DecodeString("681F63366F82A4E288DD119D679409980A1A17BBA61AB182A8C0725BD60DEC6C")),
					PartSetHeader: types.PartSetHeader{
						Total: 1,
						Hash:  First(hex.DecodeString("BE06C2B8183AA4749084784B97957AE1115F300B310BCEDD4F00A47D28DE0F5C")),
					},
				},
				Signatures: []types.CommitSig{
					{
						BlockIDFlag:      2,
						ValidatorAddress: First(hex.DecodeString("AD1A1DB19D951A2DBDF710D73711872D85E7D50D")),
						Timestamp:        First(time.Parse(time.RFC3339, "2021-02-24T04:26:01Z")),
						Signature:        First(base64.StdEncoding.DecodeString("Mr+FoDTGyi3fhiHbkoIVUe2V+YSxu7Ornq8T/nPEjwm6oDdS8ydFmkrGeGU8ezWcueCsar0GQTgvJUFQ5GkpAg==")),
					},
					{
						BlockIDFlag:      3,
						ValidatorAddress: First(hex.DecodeString("1D33EA191E17141E9B4EE5407E541FA6CE952E26")),
						Timestamp:        First(time.Parse(time.RFC3339, "2021-02-24T04:26:01Z")),
						Signature:        First(base64.StdEncoding.DecodeString("6wg3bnv5jciDMRiwhPlfaPBQhQrYegc7F8DSScboAg7rj/U6SCqyR+x913LJKil4BB5ypOG5zTq1nmjRG0s9Aw==")),
					},
					{
						BlockIDFlag:      3,
						ValidatorAddress: First(hex.DecodeString("1E4689A0633BA2F472115C36869A1D1B6D044A56")),
						Timestamp:        First(time.Parse(time.RFC3339, "2021-02-24T04:26:01Z")),
						Signature:        First(base64.StdEncoding.DecodeString("kPu0l/BVC15UmMzlJ+osaW5K9TJETkEi1TGynUxgSbEYiLcnJXK9LBpPABDvIgNXg6OiCS4KcbCw5C3ZaypBBg==")),
					},
					{
						BlockIDFlag:      3,
						ValidatorAddress: First(hex.DecodeString("76A56EEF3236312DF28486A22E1E655ADDFC1985")),
						Timestamp:        First(time.Parse(time.RFC3339, "2021-02-24T04:26:01Z")),
						Signature:        First(base64.StdEncoding.DecodeString("zqrrs5LTlwKz8wTqrs5mHum23G4RD4cNIGGwxhP5WL8C4EFYAC7i/fBMgBYTo8/drBpgPXLwVTDflgI/TqwLBg==")),
					},
					{
						BlockIDFlag:      3,
						ValidatorAddress: First(hex.DecodeString("8C4BF7429E8BD82587D733ED6C0D3798130F21D0")),
						Timestamp:        First(time.Parse(time.RFC3339, "2021-02-24T04:26:01Z")),
						Signature:        First(base64.StdEncoding.DecodeString("YRe7zFO/8hC1/Wp3SQKdoFIhe9vHamlhp+8D74mjVprZWsjfMblcFBgTDtEtJZHW1dYjeCrSooIv7z6ZID+jBA==")),
					},
					{
						BlockIDFlag:      3,
						ValidatorAddress: First(hex.DecodeString("A4828E69BDE108561732CAA60C4CD12529A7694D")),
						Timestamp:        First(time.Parse(time.RFC3339, "2021-02-24T04:26:01Z")),
						Signature:        First(base64.StdEncoding.DecodeString("g5OsEC1qPUStRBWik17dHumzPZKyOmxBDcMS5XS0TtfZlLgmHKm1fEKe/vNks3xdiMLjX0lHAMGpVHrMHeddCg==")),
					},
					{
						BlockIDFlag:      3,
						ValidatorAddress: First(hex.DecodeString("BD5F59329BB61B62A7A852BEF1E57C99438D0C86")),
						Timestamp:        First(time.Parse(time.RFC3339, "2021-02-24T04:26:01Z")),
						Signature:        First(base64.StdEncoding.DecodeString("wIoTx42MrCJ6lZ9v86Q2K7g5eBMXFjG6UN38TneKaswPG7zMN//uUSHBw9bPk5pyvDkzGee7h3ZUDDaUZWRqBQ==")),
					},
					{
						BlockIDFlag:      3,
						ValidatorAddress: First(hex.DecodeString("E77AD30C31BB8F1AF3C7E00FEF07C77F4A5021BB")),
						Timestamp:        First(time.Parse(time.RFC3339, "2021-02-24T04:26:01Z")),
						Signature:        First(base64.StdEncoding.DecodeString("cMCVRaRa/ypo1DPIGYeC23u25/MCqCHut/Capv0pZHbZJ08ZqGoQMrxTebDtPYx6KR5iulR7AXB3vc0GpdB0BQ==")),
					},
				},
			},
		}

		// 20, 30, 40, 50 - the first 3 don't have 2/3, the last 3 do!
		vals = types.ValidatorSet{
			Validators: []*types.Validator{
				{
					Address: First(hex.DecodeString("oraivalcons145dpmvvaj5dzm00hzrtnwyv89kz704gdk4xrl3")),
					// PubKey:  First(base64.StdEncoding.DecodeString("jLunee++7+9tO0vVIBG59POGwkbShGiOWbtggTZMjMM=")),
					PubKey: ed25519.PubKey(First(base64.StdEncoding.DecodeString("jLunee++7+9tO0vVIBG59POGwkbShGiOWbtggTZMjMM="))),
				},
				{
					Address: First(hex.DecodeString("oraivalcons1r5e75xg7zu2pax6wu4q8u4ql5m8f2t3xe27c6p")),
					PubKey:  ed25519.PubKey(First(base64.StdEncoding.DecodeString("l9PY/oC7El5N7BmHIhn2Rw1n+BBSKxwPKAjnh6JCQbc="))),
				},
				{
					Address: First(hex.DecodeString("oraivalcons1rergngrr8w30gus3tsmgdxsardksgjjkzs2kgh")),
					PubKey:  ed25519.PubKey(First(base64.StdEncoding.DecodeString("w8cGC01n/3SDiUgCTq8aFQgAp5lsjlOqOIhsm/s4jOs="))),
				},
				{
					Address: First(hex.DecodeString("oraivalcons1w6jkamejxccjmu5ys63zu8n9ttwlcxv9wsms6g")),
					PubKey:  ed25519.PubKey(First(base64.StdEncoding.DecodeString("iHea1XlBnUpjaE5wDBBD9XJ+I9lQj7YUMr1wJYxiSpk="))),
				},
				{
					Address: First(hex.DecodeString("oraivalcons1339lws5730vztp7hx0kkcrfhnqfs7gws69zg6m")),
					PubKey:  ed25519.PubKey(First(base64.StdEncoding.DecodeString("Fnu5TVF9wY/Z3lHlt0rTZ6Q6NCnNToKnspzMdEGEJO8="))),
				},
				{
					Address: First(hex.DecodeString("oraivalcons15jpgu6dauyy9v9eje2nqcnx3y556w62d7jlhlu")),
					PubKey:  ed25519.PubKey(First(base64.StdEncoding.DecodeString("Og7coQxbSm4cMWbgpLqJrNPTbZi3TZBqr2CT9rPrz+E="))),
				},
				{
					Address: First(hex.DecodeString("oraivalcons1h404jv5mkcdk9fag22l0retun9pc6ryxmspf9m")),
					PubKey:  ed25519.PubKey(First(base64.StdEncoding.DecodeString("0uDyc0WNK6VW98XHCYCXgyetK863YIyP31pikPp8jiU="))),
				},
			},
		}
	)

	fmt.Println(header)
	fmt.Println(newHeader)

	err := light.VerifyAdjacent(&header, &newHeader, &vals, 1*time.Hour, First(time.Parse(time.RFC3339, "2021-02-24T05:24:36Z")), maxClockDrift)

	fmt.Println(err)
}
