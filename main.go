package main

import (
	"fmt"
	"log"
	"main/block"
	"main/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	fmt.Println("Just do it")

	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	fmt.Println(walletA.PublicKeyStr())
	fmt.Println(walletB.PublicKeyStr())

	/// Wallet
	t := wallet.NewTransaction(
		walletA.PrivateKey(),
		walletA.PublicKey(),
		walletA.BlockchainAddress(),
		walletB.BlockchainAddress(),
		1.0,
	)

	/// Blockchain
	blockChain := block.NewBlockchain(walletM.BlockchainAddress())
	isAdded := blockChain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(),
		1.0, walletA.PublicKey(), t.GenerateSignature())

	fmt.Println("Added? ", isAdded)

}
