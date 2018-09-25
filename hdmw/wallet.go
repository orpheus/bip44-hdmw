package hdmw

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

type Wallet struct {
	Mnemonic   string
	Seed       []byte
	Entropy    []byte
	MasterNode *hdkeychain.ExtendedKey
}

func CreateWalletWithPassword(password string) *Wallet {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	seed, _ := bip39.NewSeedWithErrorChecking(mnemonic, password)
	//@ToDo: create network params for FLO and LTC, etc
	masterKey, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)

	wallet := Wallet{
		Mnemonic:   mnemonic,
		Seed:       seed,
		Entropy:    entropy,
		MasterNode: masterKey,
	}

	return &wallet
}

func (w *Wallet) GeneratePurposeNode() *hdkeychain.ExtendedKey {
	p, err := w.MasterNode.Child(hdkeychain.HardenedKeyStart + 44)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return p
}
