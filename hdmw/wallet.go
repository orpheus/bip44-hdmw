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

func CreateWalletWithMnemonic(mnemonic, password string) *Wallet {
	seed, _ := bip39.NewSeedWithErrorChecking(mnemonic, password)
	//@ToDo: create network params for FLO and LTC, etc
	masterKey, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)

	wallet := Wallet{
		Mnemonic: mnemonic,
		Seed:     seed,
		//ToDo: generate entropy from mnemonic or seed
		MasterNode: masterKey,
	}

	return &wallet
}

func (w *Wallet) GeneratePurposeNode() (*hdkeychain.ExtendedKey, error) {
	p, err := w.MasterNode.Child(Purpose)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return p, nil
}

func (w *Wallet) GenerateCoinNode(bip44CoinConstant uint32) (*hdkeychain.ExtendedKey, error) {
	p, err := w.GeneratePurposeNode()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	c, err := p.Child(hdkeychain.HardenedKeyStart + bip44CoinConstant)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return c, nil
}
