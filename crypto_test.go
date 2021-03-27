package main

import (
	"testing"
)

func TestHashReturnsValidHash(t *testing.T) {
	b := makeGenesis()
	h := makeFromBlock(b)
	hashedData := hash(h)
	if len(hashedData) < 64 {
		t.Fatalf("Hash has length (%v) but needs (%v)",
			len(hashedData),
			64)
	}
}

func TestSameInputReturnsSameHash(t *testing.T) {
	b := makeGenesis()
	h := makeFromBlock(b)

	hashedData := hash(h)
	hashedDataDup := hash(h)

	if hashedData != hashedDataDup {
		t.Fatalf("Expected (%v), got instead (%v)",
			hashedData,
			hashedDataDup)
	}
}
