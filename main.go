package main

import "fmt"

func main() {
	blockchain := NewBlockChain()

	blockchain.AddBlock("Send 1 BTC")
	blockchain.AddBlock("Send 2 more BTC")

	for _, block := range blockchain.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
