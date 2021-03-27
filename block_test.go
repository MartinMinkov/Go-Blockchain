package main

import (
	"strings"
	"testing"
)

func TestMinedBlockHasExpectedLastHash(t *testing.T) {
	b := makeGenesis()
	newBlock := mineBlock(b, "foo bar")
	if newBlock.ProtocolState.LastHash != b.ProtocolState.Hash {
		t.Fatalf("Newly mined block hash (%v) does not match last block hash (%v)",
			newBlock.ProtocolState.LastHash,
			b.ProtocolState.Hash)
	}
}

func TestBlockHashHasCorrectNonceValue(t *testing.T) {
	b := makeGenesis()
	oldDifficulty := b.Difficulty
	newBlock := mineBlock(b, "foobar")

	var expectedHashDifficulty []string
	for i := 0; i < oldDifficulty; i++ {
		expectedHashDifficulty = append(expectedHashDifficulty, "0")
	}

	if newBlock.ProtocolState.Hash[:oldDifficulty] != strings.Join(expectedHashDifficulty, "") {
		t.Fatalf("Expected hash to have a difficulty of %v but instead got %v",
			b.ProtocolState.Difficulty,
			newBlock.ProtocolState.Hash)
	}
}

func TestBlockDifficultyIncreases(t *testing.T) {
	b := makeGenesis()
	newTimestamp := b.ProtocolState.Timestamp + 100
	difficulty := determineDifficulty(b.ProtocolState.Timestamp,
		newTimestamp,
		b.ProtocolState.Difficulty)

	if b.ProtocolState.Difficulty+1 != difficulty {
		t.Fatalf("Expected difficulty to increase to %v after mined block but instead got %v",
			b.ProtocolState.Difficulty+1,
			difficulty)
	}
}
func TestBlockDifficultyDecreases(t *testing.T) {
	b := makeGenesis()
	newTimestamp := b.ProtocolState.Timestamp + 1100
	difficulty := determineDifficulty(b.ProtocolState.Timestamp,
		newTimestamp,
		b.ProtocolState.Difficulty)

	if b.ProtocolState.Difficulty-1 != difficulty {
		t.Fatalf("Expected difficulty to decrease to %v after mined block but instead got %v",
			b.ProtocolState.Difficulty-1,
			difficulty)
	}
}
