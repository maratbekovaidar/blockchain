package main

import (
	"fmt"
	"log"
	"main/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	fmt.Println("Just do it")

	w := wallet.NewWallet()
	//fmt.Println(w.PrivateKeyStr())
	//fmt.Println(w.PublicKeyStr())
	//
	//fmt.Println(w.PrivateKey())
	//fmt.Println(w.PublicKey())

	fmt.Println(w.BlockchainAddress())
	t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.BlockchainAddress(), "B", 1.0)
	fmt.Printf("Signature %s\n", t.GenerateSignature())
}
