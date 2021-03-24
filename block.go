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

func mineBlock(old_b Block, d Data) Block {
	var wanted_nonce []string
	for i := 0; i < old_b.ProtocolState.Difficulty; i++ {
		wanted_nonce = append(wanted_nonce, "0")
	}

	h := makeFromBlock(old_b)
	var b_hash string
	var curr_timestamp int64
	var nonce int

	mine := true
	for mine {
		curr_timestamp = makeTimestamp()
		nonce = nonce + 1

		h.timestamp = curr_timestamp
		h.nonce = nonce

		b_hash = hash(h)
		if b_hash[:old_b.ProtocolState.Difficulty] == strings.Join(wanted_nonce, "") {
			mine = false
		}
	}

	difficulty := determineDifficulty(
		old_b.ProtocolState.Timestamp,
		curr_timestamp,
		old_b.ProtocolState.Difficulty)

	return Block{
		ProtocolState: ProtocolState{
			Timestamp:  curr_timestamp,
			LastHash:   old_b.ProtocolState.Hash,
			Difficulty: difficulty,
			Nonce:      nonce,
			Hash:       b_hash,
		},
		Data: d,
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
