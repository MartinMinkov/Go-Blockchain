package main

import "fmt"

type blockchain []block

func makeBlockchain() blockchain {
	return blockchain{makeGenesis()}
}

func (bs *blockchain) addBlock(d data) {
	old_b := (*bs)[len(*bs)-1]
	b := mineBlock(old_b, d)
	*bs = append(*bs, b)
}

func isValidBlockchain(bs blockchain) bool {
	if bs[0] != makeGenesis() {
		return false
	}

	for i := 1; i < len(bs); i++ {
		b := bs[i]

		hash_i := hash_input{
			timestamp:  b.protocol_state.timestamp,
			hash:       b.protocol_state.hash,
			last_hash:  b.protocol_state.last_hash,
			difficulty: b.protocol_state.difficulty,
			nonce:      b.protocol_state.nonce,
			data:       b.data,
		}

		expected_last_hash := bs[i-1].protocol_state.hash
		expected_hash := bin_hash(hash_i)

		lastDifficulty := bs[i-1].protocol_state.difficulty

		if b.protocol_state.last_hash != expected_last_hash ||
			b.protocol_state.hash != expected_hash ||
			(b.protocol_state.difficulty-lastDifficulty) > 1 {
			return false
		}
	}
	return true
}

func replaceBlockchain(bs_curr blockchain, bs_new blockchain) blockchain {
	if len(bs_new) >= len(bs_curr) && isValidBlockchain(bs_new) {
		return bs_new
	} else {
		return bs_curr
	}
}

func equal(bs1 blockchain, bs2 blockchain) bool {
	if len(bs1) != len(bs2) {
		return false
	}
	for i := range bs1 {
		b1 := bs1[i]
		b2 := bs2[i]
		if b1 != b2 {
			fmt.Println(bs1)
			fmt.Println(bs2)
			return false
		}
	}
	return true
}

func (bs blockchain) print() {
	for _, b := range bs {
		fmt.Println(b)
	}
}
