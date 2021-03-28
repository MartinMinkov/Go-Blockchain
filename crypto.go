package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

type hash_input struct {
	timestamp  int64
	lastHash   string
	difficulty int
	nonce      int
	data       Data
}

func makeFromBlock(b Block) hash_input {
	return hash_input{
		timestamp:  b.ProtocolState.Timestamp,
		lastHash:   b.ProtocolState.LastHash,
		difficulty: b.ProtocolState.Difficulty,
		nonce:      b.ProtocolState.Nonce,
		data:       b.Data,
	}
}

func toString(h hash_input) []byte {
	return []byte(strconv.Itoa(int(h.timestamp)) + " " +
		h.lastHash + " " +
		strconv.Itoa(h.difficulty) + " " +
		strconv.Itoa(h.nonce) + " " +
		string(h.data))
}

func hash(h_input hash_input) string {
	data := toString(h_input)
	hasher := sha256.New()
	hasher.Write(data)

	return (hex.EncodeToString(hasher.Sum(nil)))
}

func bin_hash(h_input hash_input) string {
	data := hash(h_input)
	return stringToBin(data)
}

func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}
