package rsa

import (
	"testing"
)

func TestKeyGeneration(t *testing.T) {
	key := NewKeyPair()

	if key == nil {
		t.FailNow()
	}

	if key.Private.N == nil {
		t.FailNow()
	}

	if key.Private.D == nil {
		t.FailNow()
	}

	if key.Public.E == nil {
		t.FailNow()
	}

	if key.Public.N == nil {
		t.FailNow()
	}
}

func TestSign(t *testing.T) {
	key := NewKeyPair()
	msg := "sign this"
	sig := Sign(key, msg)

	if sig == nil {
		t.FailNow()
	}

}

func TestFullFlow(t *testing.T) {
	key := NewKeyPair()
	msg := "sign this"
	sig := Sign(key, msg)
	verificationRes := Verify(key, msg, sig)
	if verificationRes == false {
		t.FailNow()
	}
}
