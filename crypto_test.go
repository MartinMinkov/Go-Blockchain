package main

import (
	"testing"
)

func TestHashReturnsValidHash(t *testing.T) {
	b := makeGenesis()
	h := makeFromBlock(b)
	hashed_data := hash(h)
	if len(hashed_data) < 64 {
		t.Fatalf("Hash has length (%v) but needs (%v)", len(hashed_data), 64)
	}
}

func TestSameInputReturnsSameHash(t *testing.T) {
	b := makeGenesis()
	h := makeFromBlock(b)

	hashed_data := hash(h)
	hashed_data_dup := hash(h)

	if hashed_data != hashed_data_dup {
		t.Fatalf("Expected (%v), got instead (%v)", hashed_data, hashed_data_dup)
	}
}
