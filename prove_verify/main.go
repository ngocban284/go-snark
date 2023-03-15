package main

import (
	"os"

	prover "github.com/iden3/go-rapidsnark/prover"
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

	// use typeProof because prover.Groth16Prover returns a typeProof.ZKProof
	proof, err := prover.Groth16Prover(provingKey, witness)
	if err != nil {
		panic(err)
	}

	// print proof . proof is a typeProof.ZKProof
	print("proof:", proof.Proof)
	print("\n")
	print("proof.A:", proof.Proof.A)
	print("\n")
	print("proof.B:", proof.Proof.B)
	print("\n")
	print("proof.C:", proof.Proof.C)
	print("\n")
	print("publicSignals:", proof.PubSignals)
	print("\n")
}
