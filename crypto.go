package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"strconv"
)

type HashInput struct {
	timestamp  int64
	lastHash   string
	difficulty int
	nonce      int
	data       Data
}

func makeFromBlock(b Block) HashInput {
	return HashInput{
		timestamp:  b.ProtocolState.Timestamp,
		lastHash:   b.ProtocolState.LastHash,
		difficulty: b.ProtocolState.Difficulty,
		nonce:      b.ProtocolState.Nonce,
		data:       b.Data,
	}
}

func toString(h HashInput) []byte {
	return []byte(strconv.Itoa(int(h.timestamp)) + " " +
		h.lastHash + " " +
		strconv.Itoa(h.difficulty) + " " +
		strconv.Itoa(h.nonce) + " " +
		string(h.data))
}

func hexHash(h HashInput) string {
	data := toString(h)
	hasher := sha256.New()
	hasher.Write(data)
	return (hex.EncodeToString(hasher.Sum(nil)))
}

func binHash(h HashInput) string {
	data := hexHash(h)
	return stringToBin(data)
}

func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}

func generateKeypair() (*ecdsa.PrivateKey, ecdsa.PublicKey, error) {
	pubkeyCurve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
	return privateKey, privateKey.PublicKey, err

}

// https://stackoverflow.com/a/41315404/4160498
func encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncoded), string(pemEncodedPub)
}

// https://stackoverflow.com/a/41315404/4160498
func decode(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return privateKey, publicKey
}
