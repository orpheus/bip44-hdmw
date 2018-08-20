package hdmw

import (
	"database/sql/driver"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

type AddressNode struct {
	ExtendedKey *hdkeychain.ExtendedKey
	//@ToDo: Fill out struct
}

type ChainNode struct {
	ExtendedKey *hdkeychain.ExtendedKey
	isExternal  bool
	isInternal  bool
	Addresses   []*AddressNode
}

type AccountNode struct {
	ExtendedKey   *hdkeychain.ExtendedKey
	ExternalChain []*ChainNode
	InternalChain []*ChainNode
	isHardened    bool
}

type CoinNode struct {
	ExtendedKey *hdkeychain.ExtendedKey
	Name        string
	Ticker      string
	//@ToDo: Add Network struct
	Network    driver.Null
	Accounts   []*AccountNode
	isHardened bool
}

type PurposeNode struct {
	ExtendedKey *hdkeychain.ExtendedKey
	Index       uint32
	Coins       []*CoinNode
	isHardened  bool
}

type MasterNode struct {
	MasterExtendedKey *hdkeychain.ExtendedKey
	Child             *PurposeNode
}

type Wallet struct {
	Mnemonic   string
	Seed       []byte
	Entropy    []byte
	MasterNode *MasterNode
}

func CreateWalletWithPassword(password string) *Wallet {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	seed, _ := bip39.NewSeedWithErrorChecking(mnemonic, password)
	//@ToDo: create network params for FLO and LTC, etc
	masterKey, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)

	masterNode := &MasterNode{
		MasterExtendedKey: masterKey,
	}

	wallet := &Wallet{
		Mnemonic:   mnemonic,
		Seed:       seed,
		Entropy:    entropy,
		MasterNode: masterNode,
	}

	return wallet
}
