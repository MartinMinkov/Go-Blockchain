package main

import (
	"strings"
	"time"
)

type ProtocolState struct {
	Timestamp  int64  `json:"timestamp"`
	Hash       string `json:"hash"`
	LastHash   string `json:"lastHash"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
}

const MINE_RATE = 200 // in ms

type Data string

type Block struct {
	ProtocolState `json:"ProtocolState"`
	Data          `json:"Data"`
}

func makeGenesis() Block {
	return Block{
		ProtocolState: ProtocolState{
			Timestamp:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC).UnixNano() / int64(time.Millisecond),
			LastHash:   "",
			Hash:       "genesis-hash",
			Difficulty: 3,
			Nonce:      0,
		},
		Data: "",
	}
}

func mineBlock(oldBlock Block, d Data) Block {
	var wantedNonce []string
	for i := 0; i < oldBlock.ProtocolState.Difficulty; i++ {
		wantedNonce = append(wantedNonce, "0")
	}

	var b_hash string
	h := hash_input{}
	h.lastHash = oldBlock.ProtocolState.Hash
	h.data = d

	mine := true
	for mine {
		h.timestamp = makeTimestamp()
		h.nonce = h.nonce + 1
		h.difficulty = determineDifficulty(
			oldBlock.ProtocolState.Timestamp,
			h.timestamp,
			oldBlock.ProtocolState.Difficulty)

		b_hash = hash(h)
		if b_hash[:oldBlock.ProtocolState.Difficulty] == strings.Join(wantedNonce, "") {
			mine = false
		}
	}

	return Block{
		ProtocolState: ProtocolState{
			Timestamp:  h.timestamp,
			LastHash:   h.lastHash,
			Difficulty: h.difficulty,
			Nonce:      h.nonce,
			Hash:       b_hash,
		},
		Data: h.data,
	}
}

func determineDifficulty(old_timestamp int64, curr_timestamp int64, difficulty int) int {
	if curr_timestamp-old_timestamp < MINE_RATE {
		return difficulty + 1
	} else {
		// Difficulty should not be below 0
		if difficulty >= 1 {
			return difficulty - 1
		} else {
			return difficulty
		}
	}
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
