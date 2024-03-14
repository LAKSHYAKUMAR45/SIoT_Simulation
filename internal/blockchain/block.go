// block.go
// This file contains the definition of a block structure for the blockchain.

// Prompt:
// - Define a struct for representing a block in the blockchain.
// - Include fields such as index, timestamp, data, previous hash, and hash.
// - Implement functions for creating a new block, calculating the hash of a block, and validating a block's integrity.

package blockchain

import (
    "crypto/sha256"
    "encoding/hex"
    "time"
)

type Block struct {
    Index        int
    Timestamp    time.Time
    Data         []byte
    PrevHash     string
    Hash         string
}

// CalculateHash calculates the hash of the block
func (b *Block) CalculateHash() {
    header := string(b.Index) + b.Timestamp.String() + string(b.Data) + b.PrevHash
    hash := sha256.Sum256([]byte(header))
    b.Hash = hex.EncodeToString(hash[:])
}
