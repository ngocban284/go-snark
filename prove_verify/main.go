package main

import (
	"os"

	prover "github.com/iden3/go-rapidsnark/prover"
	"github.com/iden3/go-rapidsnark/verifier"
)

func main() {

	provingKey, err := os.ReadFile("./circuit_0000.zkey")
	if err != nil {
		panic(err)
	}

	witness, err := os.ReadFile("./witness.wtns")
	if err != nil {
		panic(err)
	}

	vetificationKey, err := os.ReadFile("./verification_key.json")
	if err != nil {
		panic(err)
	}
	// use typeProof because prover.Groth16Prover returns a typeProof.ZKProof
	proof, err := prover.Groth16Prover(provingKey, witness)
	if err != nil {
		panic(err)
	}

	//verifier
	err = verifier.VerifyGroth16(*proof, vetificationKey)
	if err != nil {
		panic(err)
	}
	print("Proof verified!")
}
