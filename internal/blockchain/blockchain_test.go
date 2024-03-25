package blockchain

import (
	"testing"
	"time"
)

func TestCalculateHash(t *testing.T) {
	// Create a sample block
	block := &Block{
		Index:     1,
		Timestamp: time.Now(),
		Data:      []byte("Test data"),
		PrevHash:  "prevHash",
	}

	// Calculate the hash of the block
	calculatedHash := block.CalculateHash()

	// Set the block's hash field
	block.Hash = calculatedHash

	// Define the expected hash value
	expectedHash := "8cb2237d0679ca88db6464eac60da96345513964f9f3"

	// Compare the calculated hash with the expected hash
	if block.Hash != expectedHash {
		t.Errorf("Expected hash: %s, but got: %s", expectedHash, block.Hash)
	}
}

func TestBlockchainIsValid(t *testing.T) {
	// Create a sample blockchain
	blockchain := NewBlockchain()
	blockchain.AddBlock([]byte("Test data"))

	// Validate the blockchain
	isValid := blockchain.IsValid()

	// The blockchain should be valid in this example
	if !isValid {
		t.Error("Blockchain validation failed: expected valid blockchain")
	}
}
