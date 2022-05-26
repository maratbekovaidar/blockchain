package main

import (
	"fmt"
	"log"
	"main/blockchain"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	fmt.Println("Just do it")

	// Create block
	//b := block.NewBlock(0, "init hash")\
	//b.Print()

	/// Create Blockchain
	blockChain := blockchain.NewBlockchain()
	blockChain.Print()
	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, previousHash)
	blockChain.Print()
	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, previousHash)
	blockChain.Print()

	/// Hash block
	//b := &block.Block{Nonce: 1}
	//fmt.Printf("%x\n", b.Hash())

}
