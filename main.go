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
	blockChain.CreateBlock(5, "hash 1")
	blockChain.Print()
	blockChain.CreateBlock(2, "hash 2")
	blockChain.Print()
}
