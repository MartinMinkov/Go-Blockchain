package main

import "fmt"

type Blockchain []Block

func makeBlockchain() Blockchain {
	return Blockchain{makeGenesis()}
}

func (bs *Blockchain) addBlock(d Data) {
	old_b := (*bs)[len(*bs)-1]
	b := mineBlock(old_b, d)
	*bs = append(*bs, b)
}

func isValidBlockchain(bs Blockchain) bool {
	if bs[0] != makeGenesis() {
		return false
	}

	for i := 1; i < len(bs); i++ {
		b := bs[i]

		hash_i := hash_input{
			timestamp:  b.ProtocolState.Timestamp,
			hash:       b.ProtocolState.Hash,
			last_hash:  b.ProtocolState.LastHash,
			difficulty: b.ProtocolState.Difficulty,
			nonce:      b.ProtocolState.Nonce,
			data:       b.Data,
		}

		expected_last_hash := bs[i-1].ProtocolState.Hash
		expected_hash := bin_hash(hash_i)

		lastDifficulty := bs[i-1].ProtocolState.Difficulty

		if b.ProtocolState.LastHash != expected_last_hash ||
			b.ProtocolState.Hash != expected_hash ||
			(b.ProtocolState.Difficulty-lastDifficulty) > 1 {
			return false
		}
	}
	return true
}

func replaceBlockchain(bs_curr Blockchain, bs_new Blockchain) Blockchain {
	if len(bs_new) >= len(bs_curr) && isValidBlockchain(bs_new) {
		return bs_new
	} else {
		return bs_curr
	}
}

func equal(bs1 Blockchain, bs2 Blockchain) bool {
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

func (bs Blockchain) print() {
	for _, b := range bs {
		fmt.Println(b)
	}
}
