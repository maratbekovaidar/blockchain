package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"main/transaction"
	"time"
)

type Block struct {
	Nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*transaction.Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*transaction.Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.Nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64                      `json:"timestamp"`
		Nonce        int                        `json:"nonce"`
		PreviousHash [32]byte                   `json:"previousHash"`
		Transactions []*transaction.Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.Nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

func (b *Block) Print() {
	fmt.Printf("timestamp      %d\n", b.timestamp)
	fmt.Printf("nonce          %d\n", b.Nonce)
	for _, t := range b.transactions {
		t.Print()
	}
	fmt.Printf("transactions   %x\n", b.transactions)
}
