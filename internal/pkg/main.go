// main.go
package main

import (
    "fmt"
    "github.com/yourusername/blockchain"
)

func main() {
    bc := blockchain.NewBlockchain()
    bc.AddBlock([]byte("Data 1"))
    bc.AddBlock([]byte("Data 2"))

    fmt.Println("Blockchain is valid:", bc.IsValid())
    fmt.Println("Blocks in the blockchain:")
    for _, block := range bc.Blocks {
        fmt.Printf("Index: %d, Timestamp: %s, Data: %s, Hash: %s\n", block.Index, block.Timestamp, block.Data, block.Hash)
    }
}
