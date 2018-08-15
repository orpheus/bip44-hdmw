package hdmw

import (
	"github.com/tyler-smith/go-bip39"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
)

type Wallet struct {
	Mnemonic          string
	Seed              []byte
	Entropy           []byte
	MasterExtendedKey *hdkeychain.ExtendedKey
}

func CreateWallet(password string) *Wallet {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	seed, _ := bip39.NewSeedWithErrorChecking(mnemonic, password)
	masterKey, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	wallet := Wallet{
		Mnemonic:          mnemonic,
		Seed:              seed,
		Entropy:           entropy,
		MasterExtendedKey: masterKey,
	}
	return &wallet
}