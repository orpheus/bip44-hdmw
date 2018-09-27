package hdmw

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
	"log"
)

type Address struct {
	Address *hdkeychain.ExtendedKey
}

type Chain struct {
	Chain *hdkeychain.ExtendedKey
}

type Account struct {
	Account *hdkeychain.ExtendedKey
}

type Coin struct {
	Name string
	Coin *hdkeychain.ExtendedKey
}

type Wallet struct {
	Entropy     []byte
	Mnemonic    string
	Seed        []byte
	MasterNode  *hdkeychain.ExtendedKey
	PurposeNode *hdkeychain.ExtendedKey
	Coins       []*Coin
}

//ToDo: add error checking - return err
func CreateWalletWithPassword(password string) *Wallet {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	seed, _ := bip39.NewSeedWithErrorChecking(mnemonic, password)
	//@ToDo: create network params for FLO and LTC, etc
	masterKey, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	purposeNode, _ := masterKey.Child(Purpose)

	wallet := Wallet{
		Mnemonic:    mnemonic,
		Seed:        seed,
		Entropy:     entropy,
		MasterNode:  masterKey,
		PurposeNode: purposeNode,
	}

	return &wallet
}

func CreateWalletWithMnemonic(mnemonic, password string) *Wallet {
	seed, _ := bip39.NewSeedWithErrorChecking(mnemonic, password)
	//@ToDo: create network params for FLO and LTC, etc
	masterKey, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	purposeNode, _ := masterKey.Child(Purpose)

	wallet := Wallet{
		Mnemonic: mnemonic,
		Seed:     seed,
		//ToDo: generate entropy from mnemonic or seed
		MasterNode:  masterKey,
		PurposeNode: purposeNode,
	}

	return &wallet
}

func (w *Wallet) Initialize(bip44CoinConstants []uint32) (*Wallet, error) {

	for i := 0; i < len(bip44CoinConstants); i++ {
		//ToDo: make this dynamic to where it will choose the network configs based on the constant
		c, err := w.InitializeCoinNode(&chaincfg.MainNetParams, bip44CoinConstants[i])
		if err != nil {
			log.Fatal("Failed to generate coin node: terminate.")
		}

		w.Coins = append(w.Coins, c)
	}

	return w, nil
}

//pkg/errors w errors.wrap
func (w *Wallet) InitializeCoinNode(network *chaincfg.Params, bip44CoinConstant uint32) (coin *Coin, err error) {
	c, err := w.PurposeNode.Child(hdkeychain.HardenedKeyStart + bip44CoinConstant)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	c.SetNet(network)

	_coin := &Coin{
		Coin: c,
	}

	return _coin, nil
}

func (c *Coin) GenerateAccountNode(index uint32) (account *Account, err error) {
	a, err := c.Coin.Child(index)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_account := &Account{
		Account: a,
	}

	return _account, nil
}

func (a *Account) GenerateChainNode(index uint32) (chain *Chain, err error) {
	c, err := a.Account.Child(index)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_chain := &Chain{
		Chain: c,
	}

	return _chain, nil
}

func (c *Chain) GenerateAddressNode(index uint32) (address *Address, err error) {
	a, err := c.Chain.Child(index)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_address := &Address{
		Address: a,
	}

	return _address, nil
}
