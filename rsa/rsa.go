package rsa

import (
	"crypto/rand"
	"github.com/tronch0/crypt0/bigint"
	"math/big"
)

type KeyPair struct {
	Public  *PublicKey
	Private *PrivateKey
}
type PrivateKey struct {
	D *big.Int
	N *big.Int
}
type PublicKey struct {
	E *big.Int
	N *big.Int
}

func NewKeyPair() *KeyPair {
	p, _ := rand.Prime(rand.Reader, 256)
	q, _ := rand.Prime(rand.Reader, 256)
	n := new(big.Int).Mul(p, q)

	one, _ := new(big.Int).SetString("1", 10)
	a := new(big.Int).Sub(p, one)
	b := new(big.Int).Sub(q, one)
	phi := new(big.Int).Mul(a, b)
	e := new(big.Int).SetInt64(65537)
	d := new(big.Int).ModInverse(e, phi)

	return &KeyPair{
		Private: &PrivateKey{N: n, D: d},
		Public:  &PublicKey{N: n, E: e},
	}
}

func Sign(keyPair *KeyPair, msg string) *big.Int {
	hashedMsgNum := bigint.HashStringToBigInt(msg)
	sig := new(big.Int).Exp(hashedMsgNum, keyPair.Private.D, keyPair.Private.N)

	return sig
}

func Verify(keyPair *KeyPair, msg string, sig *big.Int) bool {
	m := bigint.HashStringToBigInt(msg)
	m1 := new(big.Int).Exp(sig, keyPair.Public.E, keyPair.Public.N)

	return m.Cmp(m1) == 0
}
