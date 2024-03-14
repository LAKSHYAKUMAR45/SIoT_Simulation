// block.go
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
