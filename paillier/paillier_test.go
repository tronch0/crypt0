package paillier

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	keys, err := GenerateKey(rand.Reader)
	if err != nil {
		t.FailNow()
	}

	n := new(big.Int).SetInt64(15)

	c, err := Encrypt(&keys.PublicKey, n.Bytes())
	if err != nil {
		t.FailNow()
	}

	d, err := Decrypt(keys, c)
	if err != nil {
		t.FailNow()
	}

	res := new(big.Int).SetBytes(d)

	if res.Cmp(new(big.Int).SetInt64(15)) != 0 {
		t.FailNow()
	}
}

func TestAdd(t *testing.T) {
	keys, err := GenerateKey(rand.Reader)
	if err != nil {
		t.FailNow()
	}

	n := new(big.Int).SetInt64(15)

	cipher, err := Encrypt(&keys.PublicKey, n.Bytes())
	if err != nil {
		t.FailNow()
	}

	addResEnc, err := Add(&keys.PublicKey, cipher, new(big.Int).SetInt64(10).Bytes())
	if err != nil {
		t.FailNow()
	}

	addRes, err := Decrypt(keys, addResEnc)
	if err != nil {
		t.FailNow()
	}

	res := new(big.Int).SetBytes(addRes)

	if res.Cmp(new(big.Int).SetInt64(25)) != 0 {
		t.FailNow()
	}
}

func TestMul(t *testing.T) {
	keys, err := GenerateKey(rand.Reader)
	if err != nil {
		t.FailNow()
	}

	n := new(big.Int).SetInt64(15)

	cipher, err := Encrypt(&keys.PublicKey, n.Bytes())
	if err != nil {
		t.FailNow()
	}

	mulResEnc := Mul(&keys.PublicKey, cipher, new(big.Int).SetInt64(10).Bytes())

	mulRes, err := Decrypt(keys, mulResEnc)
	if err != nil {
		t.FailNow()
	}

	res := new(big.Int).SetBytes(mulRes)

	if res.Cmp(new(big.Int).SetInt64(150)) != 0 {
		t.FailNow()
	}
}
