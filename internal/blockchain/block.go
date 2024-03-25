package blockchain

import (
    "crypto/sha256"
    "encoding/hex"
    "strconv"
    "time"
)

type Block struct {
    Index        int
    Timestamp    time.Time
    Data         []byte
    PrevHash     string
    Hash         string
}

// CalculateHash calculates the hash of the block and returns it
func (b *Block) CalculateHash() string {
    header := strconv.Itoa(b.Index) + b.Timestamp.String() + string(b.Data) + b.PrevHash
    hash := sha256.Sum256([]byte(header))
    return hex.EncodeToString(hash[:])
}
