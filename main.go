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

	myBlockchainAddress := "my_blockchain_address"

	blockChain := NewBlockchain(myBlockchainAddress)
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 12.0)
	blockChain.Mining()
	blockChain.Print()

	blockChain.AddTransaction("C", "B", 2.0)
	blockChain.AddTransaction("X", "Y", 2.0)
	blockChain.Mining()
	blockChain.Print()

	/// Hash block
	//b := &block.Block{Nonce: 1}
	//fmt.Printf("%x\n", b.Hash())

}
