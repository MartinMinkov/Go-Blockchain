package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

type hash_input struct {
	timestamp  int64
	hash       string
	last_hash  string
	difficulty int
	nonce      int
	data       data
}

func makeFromBlock(b block) hash_input {
	return hash_input{
		timestamp:  b.protocol_state.timestamp,
		hash:       b.protocol_state.hash,
		last_hash:  b.protocol_state.last_hash,
		difficulty: b.protocol_state.difficulty,
		nonce:      b.protocol_state.nonce,
		data:       b.data,
	}
}

func toString(h hash_input) []byte {
	return []byte(strconv.Itoa(int(h.timestamp)) + " " +
		h.hash + " " +
		h.last_hash + " " +
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
