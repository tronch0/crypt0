package additive

import (
	"encoding/hex"
	"github.com/tronch0/crypt0/field"
	"math/big"
	"testing"
)

func TestSplit(t *testing.T) {
	secretB, _ := hex.DecodeString("secret")
	s := new(big.Int).SetBytes(secretB)
	order := new(big.Int).SetInt64(500000)

	secret := field.New(s, order)
	shares := Split(secret, 5)

	if len(shares) != 5 {
		t.FailNow()
	}
}

func TestRecover(t *testing.T) {
	order := new(big.Int).SetInt64(50)
	expectedSecret := field.New(new(big.Int).SetInt64(30), order)

	shares := []*field.Element{
		field.New(new(big.Int).SetInt64(40), order),
		field.New(new(big.Int).SetInt64(5), order),
		field.New(new(big.Int).SetInt64(36), order),
		field.New(new(big.Int).SetInt64(49), order),
	}

	secret := Recover(shares)

	if secret.Cmp(expectedSecret) != 0 {
		t.FailNow()
	}
}

func TestFullFlow(t *testing.T) {
	secretB, _ := hex.DecodeString("secret")
	s := new(big.Int).SetBytes(secretB)
	order := new(big.Int).SetInt64(500000)

	secret := field.New(s, order)
	shares := Split(secret, 5)

	if len(shares) != 5 {
		t.FailNow()
	}

	recoveredSecret := Recover(shares)

	if secret.Cmp(recoveredSecret) != 0 {
		t.FailNow()
	}
}
