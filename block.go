package main

import (
	"strings"
	"time"
)

type protocol_state struct {
	timestamp  int64
	hash       string
	last_hash  string
	difficulty int
	nonce      int
}

const MINE_RATE = 200 // in ms

type data string

type block struct {
	protocol_state
	data
}

func makeGenesis() block {
	return block{
		protocol_state: protocol_state{
			timestamp:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC).UnixNano() / int64(time.Millisecond),
			last_hash:  "",
			hash:       "genesis-hash",
			difficulty: 3,
			nonce:      0,
		},
		data: "",
	}
}

func mineBlock(old_b block, d data) block {
	var wanted_nonce []string
	for i := 0; i < old_b.protocol_state.difficulty; i++ {
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
		if b_hash[:old_b.protocol_state.difficulty] == strings.Join(wanted_nonce, "") {
			mine = false
		}
	}

	difficulty := determineDifficulty(
		old_b.protocol_state.timestamp,
		curr_timestamp,
		old_b.difficulty)

	return block{
		protocol_state: protocol_state{
			timestamp:  curr_timestamp,
			last_hash:  old_b.protocol_state.hash,
			difficulty: difficulty,
			nonce:      nonce,
			hash:       b_hash,
		},
		data: d,
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
