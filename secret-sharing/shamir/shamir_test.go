package shamir

import (
	"github.com/tronch0/crypt0/field"
	"math/big"
	"testing"
)

func TestSplit(t *testing.T) {
	s := new(big.Int).SetInt64(5)
	secret := field.New(s, new(big.Int).SetInt64(117))
	shares, _ := Split(secret, 5, 3)

	if len(shares) != 5 {
		t.FailNow()
	}

	for i := 0; i < len(shares); i++ {
		if shares[i].X == nil {
			t.FailNow()
		}

		if shares[i].Y == nil {
			t.FailNow()
		}
	}
}

func TestRecover(t *testing.T) {
	y1, _ := new(big.Int).SetString("1943205921343591736349625943162149943061", 10)
	y2, _ := new(big.Int).SetString("5720132587727048682831669377789156280081", 10)
	y3, _ := new(big.Int).SetString("11330779999150370839446130303881019011065", 10)

	shares := []Point{
		Point{
			X: field.New(new(big.Int).SetInt64(1), new(big.Int).SetInt64(117)),
			Y: field.New(y1, new(big.Int).SetInt64(117)),
		},
		Point{
			X: field.New(new(big.Int).SetInt64(2), new(big.Int).SetInt64(117)),
			Y: field.New(y2, new(big.Int).SetInt64(117)),
		},
		Point{
			X: field.New(new(big.Int).SetInt64(3), new(big.Int).SetInt64(117)),
			Y: field.New(y3, new(big.Int).SetInt64(117)),
		},
	}

	expectedSecret := field.New(new(big.Int).SetInt64(5), new(big.Int).SetInt64(117))
	secret := Recover(shares)

	if secret.Cmp(expectedSecret) != 0 {
		t.FailNow()
	}
}

func TestFullFlow(t *testing.T) {
	s := new(big.Int).SetInt64(44535)
	order := new(big.Int).SetInt64(17)

	secret := field.New(s, order)
	shares, _ := Split(secret, 5, 3)
	sharesSubset := shares[:3]

	recoverSecret := Recover(sharesSubset)

	if recoverSecret.Cmp(secret) != 0 {
		t.FailNow()
	}
}

func TestComputePolynomial(t *testing.T) {
	// y = 1 + 2*2 + 2*2**2 = 13
	coefs := []*field.Element{
		field.New(new(big.Int).SetInt64(1), new(big.Int).SetInt64(117)),
		field.New(new(big.Int).SetInt64(2), new(big.Int).SetInt64(117)),
		field.New(new(big.Int).SetInt64(2), new(big.Int).SetInt64(117)),
	}
	x := field.New(new(big.Int).SetInt64(2), new(big.Int).SetInt64(117))
	y := computePolynomial(coefs, x)

	expectedY := field.New(new(big.Int).SetInt64(13), new(big.Int).SetInt64(117))

	if expectedY.Cmp(y) != 0 {
		t.FailNow()
	}
}
