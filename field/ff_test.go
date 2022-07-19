package field

import (
	"crypto/rand"
	"math/big"
	"testing"
)

var (
	prime, _ = rand.Prime(rand.Reader, 1024)
)

func TestCreation(t *testing.T) {
	n := big.NewInt(5)
	order := big.NewInt(9)
	fe := New(n, order)

	if fe.n.Cmp(big.NewInt(5)) != 0 {
		t.Fatal("field element creation assert failed")
	}

	if fe.order.Cmp(big.NewInt(9)) != 0 {
		t.Fatal("field element creation assert failed")
	}
}

func TestCmp(t *testing.T) {
	n1 := big.NewInt(5)
	order1 := big.NewInt(9)
	fe1 := New(n1, order1)

	n2 := big.NewInt(5)
	order2 := big.NewInt(9)
	fe2 := New(n2, order2)

	if fe1.Cmp(fe2) != 0 {
		t.Fatal("field element cmp assert failed")
	}

	n2 = big.NewInt(6)
	fe2 = New(n2, order2)

	if (fe1.Cmp(fe2) < 0) == false {
		t.Fatal("field element cmp assert failed")
	}

	n2 = big.NewInt(4)
	fe2 = New(n2, order2)

	if (fe1.Cmp(fe2) > 0) == false {
		t.Fatal("field element cmp assert failed")
	}

}

func TestCmpIncorrectOrder(t *testing.T) {
	n1 := big.NewInt(5)
	order1 := big.NewInt(9)
	fe1 := New(n1, order1)

	n2 := big.NewInt(6)
	order2 := big.NewInt(7)
	fe2 := New(n2, order2)

	defer func() { recover() }()

	fe1.Cmp(fe2)

	// Never reaches here if panics (expected panic).
	t.Errorf("did not panic")
}

func TestAdd1(t *testing.T) {
	n1 := big.NewInt(2)
	order1 := big.NewInt(9)
	fe1 := New(n1, order1)

	n2 := big.NewInt(3)
	order2 := big.NewInt(9)
	fe2 := New(n2, order2)

	res := fe1.Add(fe2)

	expectedN := big.NewInt(5)
	expectedOrder := big.NewInt(9)
	expectedFe := New(expectedN, expectedOrder)

	if res.Cmp(expectedFe) != 0 {
		t.Fatal("field element cmp assert failed")
	}
}

func TestAdd2(t *testing.T) {
	n1 := big.NewInt(44)
	order1 := big.NewInt(57)
	fe1 := New(n1, order1)

	n2 := big.NewInt(33)
	order2 := big.NewInt(57)
	fe2 := New(n2, order2)

	res := fe1.Add(fe2)

	expectedN := big.NewInt(20)
	expectedOrder := big.NewInt(57)
	expectedFe := New(expectedN, expectedOrder)

	if res.Cmp(expectedFe) != 0 {
		t.Fatal("field element cmp assert failed")
	}
}

func TestSub1(t *testing.T) {
	n1 := big.NewInt(5)
	order1 := big.NewInt(9)
	fe1 := New(n1, order1)

	n2 := big.NewInt(7)
	order2 := big.NewInt(9)
	fe2 := New(n2, order2)

	res := fe1.Sub(fe2)

	expectedN := big.NewInt(7)
	expectedOrder := big.NewInt(9)
	expectedFe := New(expectedN, expectedOrder)

	if res.Cmp(expectedFe) != 0 {
		t.Fatal("field element cmp assert failed")
	}
}

func TestSub2(t *testing.T) {
	n1 := big.NewInt(9)
	order1 := big.NewInt(57)
	fe1 := New(n1, order1)

	n2 := big.NewInt(29)
	order2 := big.NewInt(57)
	fe2 := New(n2, order2)

	res := fe1.Sub(fe2)

	expectedN := big.NewInt(37)
	expectedOrder := big.NewInt(57)
	expectedFe := New(expectedN, expectedOrder)

	if res.Cmp(expectedFe) != 0 {
		t.Fatal("field element cmp assert failed")
	}
}

func TestMul1(t *testing.T) {
	n1 := big.NewInt(95)
	order1 := big.NewInt(97)
	fe1 := New(n1, order1)

	n2 := big.NewInt(45)
	order2 := big.NewInt(97)
	fe2 := New(n2, order2)

	n3 := big.NewInt(31)
	order3 := big.NewInt(97)
	fe3 := New(n3, order3)

	res := fe1.Mul(fe2)

	res = res.Mul(fe3)

	expectedN := big.NewInt(23)
	expectedOrder := big.NewInt(97)
	expectedFe := New(expectedN, expectedOrder)

	if res.Cmp(expectedFe) != 0 {
		t.Fatal("field element cmp assert failed")
	}
}

func TestMul2(t *testing.T) {
	n1 := big.NewInt(17)
	order1 := big.NewInt(97)
	fe1 := New(n1, order1)

	n2 := big.NewInt(13)
	order2 := big.NewInt(97)
	fe2 := New(n2, order2)

	n3 := big.NewInt(19)
	order3 := big.NewInt(97)
	fe3 := New(n3, order3)

	n4 := big.NewInt(44)
	order4 := big.NewInt(97)
	fe4 := New(n4, order4)

	res := fe1.Mul(fe2)

	res = res.Mul(fe3)

	res = res.Mul(fe4)

	expectedN := big.NewInt(68)
	expectedOrder := big.NewInt(97)
	expectedFe := New(expectedN, expectedOrder)

	if res.Cmp(expectedFe) != 0 {
		t.Fatal("field element cmp assert failed")
	}
}

func TestMulScalar1(t *testing.T) {
	n1 := big.NewInt(17)
	order1 := big.NewInt(97)
	fe1 := New(n1, order1)

	coef := big.NewInt(551)

	res := fe1.ScalarMul(coef)

	expectedN := big.NewInt(55)
	expectedOrder := big.NewInt(97)
	expectedFe := New(expectedN, expectedOrder)
	if res.Cmp(expectedFe) != 0 {
		t.Fatal("field element cmp assert failed")
	}
}

func TestMulScalar2(t *testing.T) {
	n1 := big.NewInt(17)
	order1 := big.NewInt(97)
	fe1 := New(n1, order1)

	coef := big.NewInt(2)

	res := fe1.ScalarMul(coef)

	expectedFe := fe1.Add(fe1)

	if res.Cmp(expectedFe) != 0 {
		t.Fatal("field element cmp assert failed")
	}
}

func TestExpo1(t *testing.T) {
	n1 := big.NewInt(17)
	order1 := big.NewInt(97)
	fe1 := New(n1, order1)

	res := fe1.Expo(big.NewInt(43))

	expectedN := big.NewInt(87)
	expectedOrder := big.NewInt(97)
	expectedFe := New(expectedN, expectedOrder)

	if res.Cmp(expectedFe) != 0 {
		t.Fatal("field element cmp assert failed")
	}
}

func TestExpo2(t *testing.T) {
	n1 := big.NewInt(17)
	order1 := big.NewInt(31)
	fe1 := New(n1, order1)

	res := fe1.Expo(big.NewInt(3))

	expectedN := big.NewInt(15)
	expectedOrder := big.NewInt(31)
	expectedFe := New(expectedN, expectedOrder)

	if res.Cmp(expectedFe) != 0 {
		t.Fatal("field element cmp assert failed")
	}
}

func TestExpo3(t *testing.T) {
	n1 := big.NewInt(71)
	order1 := big.NewInt(97)
	fe1 := New(n1, order1)

	res := fe1.Expo(big.NewInt(-55))

	expectedN := big.NewInt(59)
	expectedOrder := big.NewInt(97)
	expectedFe := New(expectedN, expectedOrder)

	if res.Cmp(expectedFe) != 0 {
		t.Fatal("field element cmp assert failed")
	}
}

func TestDiv1(t *testing.T) {
	n1 := big.NewInt(95)
	order1 := big.NewInt(97)
	fe1 := New(n1, order1)

	n2 := big.NewInt(45)
	order2 := big.NewInt(97)
	fe2 := New(n2, order2)

	res := fe1.Div(fe2)

	expectedN := big.NewInt(56)
	expectedOrder := big.NewInt(97)
	expectedFe := New(expectedN, expectedOrder)

	if res.Cmp(expectedFe) != 0 {
		t.Fatal("field element cmp assert failed")
	}
}

func TestDiv2(t *testing.T) {
	n1 := big.NewInt(77)
	order1 := big.NewInt(97)
	fe1 := New(n1, order1)

	n2 := big.NewInt(50)
	order2 := big.NewInt(97)
	fe2 := New(n2, order2)

	res := fe1.Div(fe2)

	expectedN := big.NewInt(19)
	expectedOrder := big.NewInt(97)
	expectedFe := New(expectedN, expectedOrder)

	if res.Cmp(expectedFe) != 0 {
		t.Fatal("field element cmp assert failed")
	}
}

func TestClone(t *testing.T) {
	// Create a FE
	n1 := big.NewInt(2)
	order1 := big.NewInt(9)
	fe1 := New(n1, order1)

	// Clone the FE
	fe2 := Clone(fe1)

	// Add to the first FE
	n3 := big.NewInt(5)
	order3 := big.NewInt(9)
	fe3 := New(n3, order3)
	fe1 = fe1.Add(fe3)

	// Check the two FE for expected results
	expectedFe1 := New(big.NewInt(7), big.NewInt(9))
	expectedFe2 := New(big.NewInt(2), big.NewInt(9))

	if fe1.Cmp(expectedFe1) != 0 {
		t.Fatal("field element cmp assert failed")
	}

	if fe2.Cmp(expectedFe2) != 0 {
		t.Fatal("field element cmp assert failed")
	}
}

func TestGetNum(t *testing.T) {
	// Create a FE
	n1 := big.NewInt(2)
	order1 := big.NewInt(9)
	fe1 := New(n1, order1)

	res := fe1.GetNum()

	expectedRes := big.NewInt(2)

	if res.Cmp(expectedRes) != 0 {
		t.Fatal("get n cmp assert failed")
	}

	// Create a FE
	n2 := big.NewInt(9)
	order2 := big.NewInt(7)
	fe2 := New(n2, order2)

	res2 := fe2.GetNum()

	expectedRes2 := big.NewInt(2)

	if res2.Cmp(expectedRes2) != 0 {
		t.Fatal("get n cmp assert failed")
	}

}
