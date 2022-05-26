package blockchain

import (
	"fmt"
	"main/block"
	"strings"
)

type Blockchain struct {
	transactionPool []string
	chain           []*block.Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *block.Block {
	b := block.NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, b := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		b.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 59))
}
