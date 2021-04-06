package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"reflect"
	"testing"
)

func TestHashReturnsValidHash(t *testing.T) {
	b := makeGenesis()
	h := makeFromBlock(b)
	hashedData := hexHash(h)
	if len(hashedData) < 64 {
		t.Fatalf("Hash has length (%v) but needs (%v)",
			len(hashedData),
			64)
	}
}

func TestSameInputReturnsSameHash(t *testing.T) {
	b := makeGenesis()
	h := makeFromBlock(b)

	hashedData := hexHash(h)
	hashedDataDup := hexHash(h)

	if hashedData != hashedDataDup {
		t.Fatalf("Expected (%v), got instead (%v)",
			hashedData,
			hashedDataDup)
	}
}

func TestEncodeAndDecodeKeys(t *testing.T) {
	pubkeyCurve := elliptic.P256()
	privateKey, _ := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)

	publicKey := &privateKey.PublicKey
	encodedPrivateKey, encodedPublicKey := encode(privateKey, publicKey)

	newPrivateKey, newPublicKey := decode(encodedPrivateKey, encodedPublicKey)

	if !reflect.DeepEqual(privateKey, newPrivateKey) {
		t.Fatal("Private keys do not match.")
	}
	if !reflect.DeepEqual(publicKey, newPublicKey) {
		t.Fatal("Public keys do not match.")
	}
}

func TestEllipticSignature(t *testing.T) {
	pubkeyCurve := elliptic.P256()
	privateKey, _ := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
	publicKey := &privateKey.PublicKey

	h := md5.New()
	signHash := h.Sum(nil)

	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, signHash)

	verifyStatus := ecdsa.Verify(publicKey, signHash, r, s)
	if !verifyStatus {
		t.Fatal("Could not correctly verify signature")
	}
}

func TestEncodeAndDecodeEllipticSignature(t *testing.T) {
	pubkeyCurve := elliptic.P256()
	privateKey, _ := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)

	publicKey := &privateKey.PublicKey
	encodedPrivateKey, encodedPublicKey := encode(privateKey, publicKey)

	newPrivateKey, _ := decode(encodedPrivateKey, encodedPublicKey)

	h := md5.New()
	signHash := h.Sum(nil)

	r, s, _ := ecdsa.Sign(rand.Reader, newPrivateKey, signHash)

	verifyStatus := ecdsa.Verify(publicKey, signHash, r, s)
	if !verifyStatus {
		t.Fatal("Could not correctly verify signature")
	}
}
