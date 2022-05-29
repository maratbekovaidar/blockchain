package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	_ "github.com/btcsuite/btcutil/base58"
	_ "golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockchainAddress string
}

func NewWallet() *Wallet {
	// 1. Creating ECDSA private key (32 bytes) public key (64 bytes)
	w := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey

	// TODO 2. Perform SHA-256 hashing on the public key (32 bytes).
	// TODO 3. Perform RIPEMD-160 hashing on the result of SHA-256 (20 bytes)
	// TODO 4. Add version byte in front of RIPEMD-160 hash (0x00 for main network).
	// TODO 5. Perform SHA-256 hash on the extended RIPEMD-160 result
	// TODO 6. Perform SHA-256 hash on the result of the previous SHA-256 hash
	// TODO 7. Take the first 4 bytes if the second SHA-256 hash for check sum
	// TODO 8. Add the 4 checksum bytes from 7 at the end of extended RIPEMD-160 hash from 4 (25 bytes)
	// TODO 9. Convert the result form a byte string into base58
	return w
}

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}
