package hdmw

import (
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"log"
)

type Wallet struct {
	Mnemonic          string
	Seed              []byte
	Entropy           []byte
	MasterExtendedKey *bip32.Key
}

func (w *Wallet) InitializeWallet(entropy []byte, password string) *Wallet {
	mnemonic, _ := bip39.NewMnemonic(entropy)
	w.Mnemonic = mnemonic
	seed := bip39.NewSeed(mnemonic, password)
	w.Seed = seed
	w.Entropy = entropy
	masterKey, _ := bip32.NewMasterKey(seed)
	w.MasterExtendedKey = masterKey

}

func CreateEntropy(bitSize int) (entropy []byte) {
	if bitSize%32 != 0 {
		log.Fatalf("bitSize must be divisible by 32")
	}
	entropy, _ = bip39.NewEntropy(bitSize)
	return entropy
}
