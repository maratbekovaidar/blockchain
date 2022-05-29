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

	blockChain.AddTransaction("Aidar", "Abylay", 12.0)
	blockChain.Mining()
	blockChain.Print()

	blockChain.AddTransaction("C", "B", 2.0)
	blockChain.AddTransaction("X", "Y", 2.0)
	blockChain.Mining()
	blockChain.Print()

	fmt.Printf("C %.1f\n", blockChain.CalculateTotalAmount("my_blockchain_address"))
	fmt.Printf("C %.1f\n", blockChain.CalculateTotalAmount("Aidar"))
	fmt.Printf("C %.1f\n", blockChain.CalculateTotalAmount("Abylay"))

	/// Hash block
	//b := &block.Block{Nonce: 1}
	//fmt.Printf("%x\n", b.Hash())

}
