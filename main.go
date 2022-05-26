package main

import (
	"fmt"
	"log"
	"main/block"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	fmt.Println("Just do it")
	b := block.NewBlock(0, "init hash")

	b.Print()
}
