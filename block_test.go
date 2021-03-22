package main

import (
	"strings"
	"testing"
)

func TestMinedBlockHasExpectedLastHash(t *testing.T) {
	b := makeGenesis()
	new_b := mineBlock(b, "foo bar")
	if new_b.last_hash != b.hash {
		t.Fatalf("Newly mined block hash (%v) does not match last block hash (%v)", new_b.last_hash, b.hash)
	}
}

func TestBlockHashHasCorrectNonceValue(t *testing.T) {
	b := makeGenesis()
	old_difficulty := b.difficulty
	new_b := mineBlock(b, "foobar")

	var expected_hash_difficulty []string
	for i := 0; i < old_difficulty; i++ {
		expected_hash_difficulty = append(expected_hash_difficulty, "0")
	}

	if new_b.hash[:old_difficulty] != strings.Join(expected_hash_difficulty, "") {
		t.Fatalf("Expected hash to have a difficulty of %v but instead got %v", b.protocol_state.difficulty, new_b.hash)
	}
}

func TestBlockDifficultyIncreases(t *testing.T) {
	b := makeGenesis()
	new_timestamp := b.protocol_state.timestamp + 100
	difficulty := determineDifficulty(b.protocol_state.timestamp, new_timestamp, b.protocol_state.difficulty)

	if b.protocol_state.difficulty+1 != difficulty {
		t.Fatalf("Expected difficulty to increase to %v after mined block but instead got %v", b.protocol_state.difficulty+1, difficulty)
	}
}
func TestBlockDifficultyDecreases(t *testing.T) {
	b := makeGenesis()
	new_timestamp := b.protocol_state.timestamp + 1100
	difficulty := determineDifficulty(b.protocol_state.timestamp, new_timestamp, b.protocol_state.difficulty)

	if b.protocol_state.difficulty-1 != difficulty {
		t.Fatalf("Expected difficulty to decrease to %v after mined block but instead got %v", b.protocol_state.difficulty-1, difficulty)
	}
}
