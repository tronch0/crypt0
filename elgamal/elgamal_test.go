package elgamal

import (
	"fmt"
	"math/big"
	"testing"
)

func TestMul(t *testing.T) {
	// Generate key pair
	pp, err := generatePublicParam(256)
	priv_a, err := GenerateKeyPair(pp.P, pp.G)
	if err != nil {
		panic(err)
	}

	// Homomorphic addition
	msg_a := big.NewInt(7)
	msg_b := big.NewInt(10)

	c1_a, c2_a, err := Encrypt(&priv_a.PublicKey, msg_a)
	if err != nil {
		panic(err)
	}

	c1_b, c2_b, err := Encrypt(&priv_a.PublicKey, msg_b)
	if err != nil {
		panic(err)
	}

	a3, b3, err := HomomorphicMul(&priv_a.PublicKey, c1_a, c2_a, c1_b, c2_b)
	if err != nil {
		panic(err)
	}

	msg3, err := Decrypt(priv_a, a3, b3)
	if err != nil {
		panic(err)
	}

	if new(big.Int).SetInt64(70).Cmp(msg3) != 0 {
		t.Fatalf("expected 70, got %d", msg3)
	}

	fmt.Println(msg_a, "*", msg_b, "=", msg3)
}
