package main

import (
	"fmt"

	"crypto/ed25519"
)

func main() {
	// genertate key pair
	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	// sign message
	msg := []byte("hello world")
	sig := ed25519.Sign(priv, msg)
	// verify signature
	if !ed25519.Verify(pub, msg, sig) {
		panic("invalid signature")
	}
	fmt.Println("Signature verified!")
}
