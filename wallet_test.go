package main

import (
	"crypto/ecdsa"
	"crypto/md5"
	"crypto/rand"
	"testing"
)

func TestWalletCreationHasValidKeypair(t *testing.T) {
	wallet := CreateWallet()
	privKey1, pubKey1 := wallet.PrivateKey, wallet.PublicKey

	privKey2, pubKey2 := decode(privKey1, pubKey1)
	privateKey, publicKey := encode(privKey2, pubKey2)

	if wallet.PublicKey != publicKey {
		t.Fatalf("Public key in wallet creation does not encode and decode")
	}

	if wallet.PrivateKey != privateKey {
		t.Fatalf("Private key in wallet creation does not encode and decode")
	}
}

func TestWalletCanSignCorrectly(t *testing.T) {
	privateKey, publicKey := generateKeypair()
	h := md5.New()
	signHash := h.Sum(nil)

	privateKeyDecoded, publicKeyEncoded := decode(privateKey, publicKey)
	r, s, _ := ecdsa.Sign(rand.Reader, privateKeyDecoded, signHash)
	verifyStatus := ecdsa.Verify(publicKeyEncoded, signHash, r, s)

	if !verifyStatus {
		t.Fatal("Could not correctly verify signature")
	}
}
