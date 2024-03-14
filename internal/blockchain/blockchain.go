// blockchain.go
package blockchain

type Blockchain struct {
    Blocks []*Block
}

// NewBlockchain creates a new blockchain with the genesis block
func NewBlockchain() *Blockchain {
    genesisBlock := &Block{Index: 0, Timestamp: time.Now(), Data: []byte("Genesis Block"), PrevHash: ""}
    genesisBlock.CalculateHash()
    return &Blockchain{Blocks: []*Block{genesisBlock}}
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data []byte) {
    prevBlock := bc.Blocks[len(bc.Blocks)-1]
    newBlock := &Block{
        Index:     prevBlock.Index + 1,
        Timestamp: time.Now(),
        Data:      data,
        PrevHash:  prevBlock.Hash,
    }
    newBlock.CalculateHash()
    bc.Blocks = append(bc.Blocks, newBlock)
}

// IsValid checks if the blockchain is valid
func (bc *Blockchain) IsValid() bool {
    for i := 1; i < len(bc.Blocks); i++ {
        currBlock := bc.Blocks[i]
        prevBlock := bc.Blocks[i-1]

        if currBlock.Hash != currBlock.CalculateHash() {
            return false
        }

        if currBlock.PrevHash != prevBlock.Hash {
            return false
        }
    }
    return true
}
