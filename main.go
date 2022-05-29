package main

import (
	"fmt"
	"log"
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
	blockChain := NewBlockchain()
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 12.0)
	previousHash := blockChain.LastBlock().Hash()
	nonce := blockChain.ProofOfWork()
	blockChain.CreateBlock(nonce, previousHash)
	blockChain.Print()

	blockChain.AddTransaction("C", "B", 2.0)
	blockChain.AddTransaction("X", "Y", 2.0)
	previousHash = blockChain.LastBlock().Hash()
	nonce = blockChain.ProofOfWork()
	blockChain.CreateBlock(nonce, previousHash)
	blockChain.Print()

	/// Hash block
	//b := &block.Block{Nonce: 1}
	//fmt.Printf("%x\n", b.Hash())

}
