package main

import "fmt"

type Blockchain []Block

func makeBlockchain() Blockchain {
	return Blockchain{makeGenesis()}
}

func (bs *Blockchain) addBlock(d Data) {
	oldBlock := (*bs)[len(*bs)-1]
	b := mineBlock(oldBlock, d)
	*bs = append(*bs, b)
}

func isValidBlockchain(bs Blockchain) bool {
	if bs[0] != makeGenesis() {
		return false
	}

	for i := 1; i < len(bs); i++ {
		b := bs[i]

		lastDifficulty := bs[i-1].ProtocolState.Difficulty
		expectedLastHash := bs[i-1].ProtocolState.Hash
		expectedHash := hexHash(makeFromBlock(b))

		if b.ProtocolState.LastHash != expectedLastHash {
			return false
		}

		if b.ProtocolState.Hash != expectedHash {
			return false
		}

		if b.ProtocolState.Difficulty-lastDifficulty > 1 {
			return false
		}
	}
	return true
}

func (bs *Blockchain) replaceBlockchain(bsNew *Blockchain) {
	if len(*bsNew) >= len(*bs) && isValidBlockchain(*bsNew) {
		*bs = *bsNew
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
