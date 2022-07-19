package additive

import (
	"github.com/tronch0/crypt0/bigint"
	"github.com/tronch0/crypt0/field"
	"math/big"
)

func Split(secret *field.Element, numOfShares int) []*field.Element {
	fieldOrder := secret.GetOrder()
	res := []*field.Element{}
	sumOfShares := field.New(new(big.Int).SetInt64(0), fieldOrder)

	for i := 0; i < numOfShares-1; i++ {
		r := bigint.GetRandom()
		share := field.New(r, fieldOrder)
		res = append(res, share)
		sumOfShares = sumOfShares.Add(share)
	}

	res = append(res, secret.Sub(sumOfShares))

	return res
}

func Recover(shares []*field.Element) *field.Element {
	if len(shares) <= 2 {
		panic("error on number of shares")
	}
	zero := new(big.Int).SetInt64(0)
	fieldOrder := shares[0].GetOrder()
	res := field.New(zero, fieldOrder)

	for i := 0; i < len(shares); i++ {
		res = res.Add(shares[i])
	}

	return res
}
