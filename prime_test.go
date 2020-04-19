package jprimego

import (
	"testing"
)

func TestFastGeneratePrime(t *testing.T) {
	p := FastGeneratePrime(1024)
	isPrime := p.ProbablyPrime(40)
	if !isPrime {
		t.Error("FastGeneratePrime generate non prime number")
	}
}

func TestSign1Gen(t *testing.T) {
	a1 := sign1Gen(1)
	if a1 != -1 {
		t.Error("sign1Gen(1) should be -1")
	}
	a2 := sign1Gen(2)
	if a2 != 1 {
		t.Error("sign1Gen(2) should be 1")
	}
	a3 := sign1Gen(3)
	if a3 != -1 {
		t.Error("sign1Gen(3) should be -1")
	}
	a4 := sign1Gen(4)
	if a4 != 1 {
		t.Error("sign1Gen(4) should be 1")
	}
}

func TestSign2Gen(t *testing.T) {
	a1 := sign2Gen(1)
	if a1 != -1 {
		t.Error("sign2Gen(1) should be -1")
	}
	a2 := sign2Gen(2)
	if a2 != 1 {
		t.Error("sign2Gen(2) should be 1")
	}
	a3 := sign2Gen(3)
	if a3 != 1 {
		t.Error("sign2Gen(3) should be 1")
	}
	a4 := sign2Gen(4)
	if a4 != -1 {
		t.Error("sign2Gen(4) should be -1")
	}
}
