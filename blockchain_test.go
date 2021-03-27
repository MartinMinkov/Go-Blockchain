package main

import (
	"testing"
)

func TestBlockChainStartsWithGenesis(t *testing.T) {
	bc := makeBlockchain()
	genesisHash := makeGenesis().ProtocolState.Hash

	if genesisHash != bc[0].ProtocolState.Hash {
		t.Fatalf("Expected genesis hash of (%v) but got hash (%v) instead",
			genesisHash,
			bc[0].ProtocolState.Hash)
	}
}

func TestAddingNewBlockToBlockchain(t *testing.T) {
	bc := makeBlockchain()
	bc.addBlock("It's")
	bc.addBlock("Alive!")

	if len(bc) != 3 {
		t.Fatalf("Expected blockchain length of %v but got %v instead", 3, len(bc))
	}
}

func TestValidBlockChain(t *testing.T) {
	bc := makeBlockchain()
	bc.addBlock("It's")
	bc.addBlock("Alive")
	bc.addBlock("Or")
	bc.addBlock("Is")
	bc.addBlock("It?")

	if isValidBlockchain(bc) {
		t.Fatalf("Expected valid blockchain to be true, instead got false")
	}
}

func TestInvalidBlockchain(t *testing.T) {
	bc := makeBlockchain()
	bc.addBlock("It's")
	bc.addBlock("Alive")
	bc.addBlock("Or")
	bc.addBlock("Is")
	bc.addBlock("It?")

	bc[2].Data = "invalid-data"

	if isValidBlockchain(bc) {
		t.Fatalf("Expected invalid blockchain to be false, instead got true")
	}
}

func TestReplaceValidBlockchain(t *testing.T) {
	bc1 := makeBlockchain()
	bc2 := makeBlockchain()

	bc1.addBlock("It's")
	bc1.addBlock("Alive!")

	bc2.addBlock("We")
	bc2.addBlock("Shall")
	bc2.addBlock("Surpass!")

	replaced_bc := replaceBlockchain(bc1, bc2)

	if equal(replaced_bc, bc2) {
		t.Fatalf("Expected valid blockchain to replace old blockchain, instead got old blockchain")
	}
}

func TestReplaceShortValidBlockchain(t *testing.T) {
	bc1 := makeBlockchain()
	bc2 := makeBlockchain()

	bc1.addBlock("It's")
	bc1.addBlock("Alive!")

	bc2.addBlock("Or is it?")

	replaced_bc := replaceBlockchain(bc1, bc2)

	if !equal(replaced_bc, bc1) {
		t.Fatalf("Expected short valid blockchain to not replace old blockchain, instead got short blockchain")
	}
}

func TestReplaceInvalidBlockchain(t *testing.T) {
	bc1 := makeBlockchain()
	bc2 := makeBlockchain()

	bc1.addBlock("We")
	bc1.addBlock("Shall")
	bc1.addBlock("Surpass!")

	bc2.addBlock("We")
	bc2.addBlock("Shall")
	bc2.addBlock("Surpass!")
	bc2[3].Data = "Fail!"

	replaced_bc := replaceBlockchain(bc1, bc2)
	if !equal(replaced_bc, bc1) {
		t.Fatalf("Expected invalid blockchain to not replace old blockchain, instead got invalid blockchain")
	}
}

func TestBlockWithJumpedDifficulty(t *testing.T) {
	bc := makeBlockchain()

	bc.addBlock("We")
	bc.addBlock("Shall")
	bc.addBlock("Surpass!")

	bc[2].ProtocolState.Difficulty = 0

	if isValidBlockchain(bc) {
		t.Fatalf("Expected invalid blockchain to be false, instead got true")
	}

}
